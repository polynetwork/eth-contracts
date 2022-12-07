const hre = require("hardhat");
const Web3 = require("web3");
const abi = require("@ethersproject/abi");
const fs = require("fs");
const { ethers } = require("hardhat");
const { keccak256 } = require("ethereumjs-util");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './contractConfig.json'

async function main() {
    [deployer] = await hre.ethers.getSigners();

    var config = {}
    await readConfig(hre.network.name).then((netConfig) => {
        if (netConfig !== undefined) {
            config = netConfig
        }
    }).catch((err) => {
        console.error(err);
        process.exit(1);
    });
    if (config.Name === undefined) {
        config.Name = hre.network.name    
    }
    if (config.PolyChainID === undefined) {
        if (hre.config.networks[hre.network.name].polyId === undefined) {
            console.error("unknown network: invalid PolyChainID".red);
            process.exit(1);
        }
        config.PolyChainID = hre.config.networks[hre.network.name].polyId
    }
    if (config.Provider === undefined) {
        config.Provider = hre.config.networks[hre.network.name].url
        writeConfig(config)
    }
    if (config.Deployer === undefined) {
        config.Deployer = deployer.address
    }

    const LockProxy = await ethers.getContractFactory("LockProxy");
    const ProxyAdmin = await ethers.getContractFactory("ProxyAdmin");
    const RippleLockProxy = await ethers.getContractFactory("RippleLockProxy");
    const EthCrossChainData = await hre.ethers.getContractFactory("EthCrossChainData");
    const CallerFactory = await hre.ethers.getContractFactory("CallerFactoryWithAdmin");
    let EthCrossChainManagerImplementation = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const EthCrossChainManager = await hre.ethers.getContractFactory("EthCrossChainManager");
    const WrapperV1 = await hre.ethers.getContractFactory("PolyWrapperV1");
    const WrapperV2 = await hre.ethers.getContractFactory("PolyWrapperV2");
    const WrapperV3 = await hre.ethers.getContractFactory("PolyWrapperV3");
    let polyId = config.PolyChainID
    let eccd
    let ccmi
    let ccm
    let cf
    let lockProxy
    let wrapper1
    let wrapper2
    let wrapper3
    let proxyAdmin
    let rippleLockProxy
    let rippleLockProxyImpl
    let abiCoder = new abi.AbiCoder
    let rippleName = "Poly Peg Ripple Token"
    let rippleSymbol = "pXRP"
    let rippleDeciamls = 6
    let rippleInitData = "0xde7ea79d" + abiCoder.encode(['string','string','uint8','address'],[rippleName,rippleSymbol,rippleDeciamls,deployer.address]).slice(2); 
    let rippleSalt = 2;

    console.log("\nDeploy contracts on chain with Poly_Chain_Id:".cyan, polyId);
    
    if (config.LockProxy === undefined) {
        // deploy LockProxy
        console.log("\ndeploy LockProxy ......".cyan);
        lockProxy = await LockProxy.deploy();
        await lockProxy.deployed();
        console.log("LockProxy deployed to:".green, lockProxy.address.blue);
        config.LockProxy = lockProxy.address
        writeConfig(config)
    } else {
        console.log("\nLockProxy already deployed at".green, config.LockProxy.blue)
        lockProxy = await LockProxy.attach(config.LockProxy) 
    }
    
    if (config.EthCrossChainData === undefined) {
        // deploy EthCrossChainData
        console.log("\ndeploy EthCrossChainData ......".cyan);
        eccd = await EthCrossChainData.deploy();
        await eccd.deployed();
        console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
        config.EthCrossChainData = eccd.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainData already deployed at".green, config.EthCrossChainData.blue)
        eccd = await EthCrossChainData.attach(config.EthCrossChainData) 
    }
    
    if (config.CallerFactory === undefined) {
        // deploy CallerFactory
        console.log("\ndeploy CallerFactory ......".cyan);
        cf = await CallerFactory.deploy([lockProxy.address]);
        await cf.deployed();
        console.log("CallerFactory deployed to:".green, cf.address.blue);
        config.CallerFactory = cf.address
        writeConfig(config)
    } else {
        console.log("\nCallerFactory already deployed at".green, config.CallerFactory.blue)
        cf = await CallerFactory.attach(config.CallerFactory) 
    }
    
    if (config.EthCrossChainManagerImplementation === undefined) {
        // update Const.sol
        console.log("\nupdate Const.sol ......".cyan);
        await updateConst(config.PolyChainID, eccd.address, cf.address);
        console.log("Const.sol updated".green);
        await hre.run('compile');

        // deploy EthCrossChainManagerImplementation
        console.log("\ndeploy EthCrossChainManagerImplementation ......".cyan);
        EthCrossChainManagerImplementation = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
        ccmi = await EthCrossChainManagerImplementation.deploy();
        await ccmi.deployed();
        console.log("EthCrossChainManagerImplementation deployed to:".green, ccmi.address.blue);
        config.EthCrossChainManagerImplementation = ccmi.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManagerImplementation already deployed at".green, config.EthCrossChainManagerImplementation.blue)
        ccmi = await EthCrossChainManagerImplementation.attach(config.EthCrossChainManagerImplementation) 
    }
    
    if (config.EthCrossChainManager === undefined) {
        // deploy EthCrossChainManager
        console.log("\ndeploy EthCrossChainManager ......".cyan);
        ccm = await EthCrossChainManager.deploy(ccmi.address,deployer.address,'0x');
        await ccm.deployed();
        console.log("EthCrossChainManager deployed to:".green, ccm.address.blue);
        config.EthCrossChainManager = ccm.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManager already deployed at".green, config.EthCrossChainManager.blue)
        ccm = await EthCrossChainManager.attach(config.EthCrossChainManager) 
    }

    let eccdOwner = await eccd.owner()
    // transfer ownership
    if (eccdOwner == ccm.address) {
        console.log("\neccd ownership already transferred".green);
    } else {
        console.log("\ntransfer eccd's ownership to ccm ......".cyan);
        tx = await eccd.transferOwnership(ccm.address);
        await tx.wait();
        console.log("ownership transferred".green);
    }

    let alreadySetCCMP = await lockProxy.managerProxyContract();
    if (alreadySetCCMP == ccm.address) {
        console.log("\nmanagerProxyContract already set".green);
    } else {
        // setup LockProxy
        console.log("\nsetup LockProxy ......".cyan);
        tx = await lockProxy.setManagerProxy(ccm.address);
        await tx.wait();
        console.log("setManagerProxy Done".green);
    }
    
    if (config.WrapperV1 === undefined) {
        // deploy WrapperV1
        console.log("\ndeploy WrapperV1 ......".cyan);
        wrapper1 = await WrapperV1.deploy(deployer.address, polyId);
        await wrapper1.deployed();
        console.log("WrapperV1 deployed to:".green, wrapper1.address.blue);
        config.WrapperV1 = wrapper1.address
        writeConfig(config)
    } else {
        console.log("\nWrapperV1 already deployed at".green, config.WrapperV1.blue)
        wrapper1 = await WrapperV1.attach(config.WrapperV1) 
    }

    let alreadySetLockProxy1 = await wrapper1.lockProxy();
    let alreadySetFeeCollector1 = await wrapper1.feeCollector();
    console.log("\nsetup WrapperV1 ......".cyan);
    if (alreadySetLockProxy1 == lockProxy.address) {
        console.log("wrapper1 lockProxy already set".green);
    } else {
        // setLockProxy
        console.log("setLockProxy ......".cyan);
        tx = await wrapper1.setLockProxy(lockProxy.address);
        await tx.wait();
        console.log("setLockProxy Done".green);
    }
    if (alreadySetFeeCollector1 != "0x0000000000000000000000000000000000000000") {
        console.log("wrapper1 feeCollector already set".green);
    } else {
        // setFeeCollector
        console.log("setFeeCollector ......".cyan);
        tx = await wrapper1.setFeeCollector(deployer.address);
        await tx.wait();
        console.log("setFeeCollector Done".green);
    }
    
    if (config.WrapperV2 === undefined) {
        // deploy WrapperV2
        console.log("\ndeploy WrapperV2 ......".cyan);
        wrapper2 = await WrapperV2.deploy(deployer.address, polyId);
        await wrapper2.deployed();
        console.log("WrapperV2 deployed to:".green, wrapper2.address.blue);
        config.WrapperV2 = wrapper2.address
        config.Wrapper = wrapper2.address
        writeConfig(config)
    } else {
        console.log("\nWrapperV2 already deployed at".green, config.WrapperV2.blue)
        wrapper2 = await WrapperV2.attach(config.WrapperV2) 
        if (config.wrapper != wrapper2.address) {
            config.Wrapper = wrapper2.address
            writeConfig(config)
        }
    }

    let alreadySetLockProxy2 = await wrapper2.lockProxy();
    let alreadySetFeeCollector2 = await wrapper2.feeCollector();
    console.log("\nsetup WrapperV2 ......".cyan);
    if (alreadySetLockProxy2 == lockProxy.address) {
        console.log("wrapper2 lockProxy already set".green);
    } else {
        // setLockProxy
        console.log("setLockProxy ......".cyan);
        tx = await wrapper2.setLockProxy(lockProxy.address);
        await tx.wait();
        console.log("nsetLockProxy Done".green);
    }
    if (alreadySetFeeCollector2 != "0x0000000000000000000000000000000000000000") {
        console.log("wrapper2 feeCollector already set".green);
    } else {
        // setFeeCollector
        console.log("setFeeCollector ......".cyan);
        tx = await wrapper2.setFeeCollector(deployer.address);
        await tx.wait();
        console.log("setFeeCollector Done".green);
    }
    
    if (config.WrapperV3 === undefined) {
        // deploy WrapperV3
        console.log("\ndeploy WrapperV3 ......".cyan);
        wrapper3 = await WrapperV3.deploy(deployer.address, polyId);
        await wrapper3.deployed();
        console.log("WrapperV3 deployed to:".green, wrapper3.address.blue);
        config.WrapperV3 = wrapper3.address
        config.Wrapper = wrapper3.address
        writeConfig(config)
    } else {
        console.log("\nWrapperV3 already deployed at".green, config.WrapperV3.blue)
        wrapper3 = await WrapperV3.attach(config.WrapperV3) 
        if (config.wrapper != wrapper3.address) {
            config.Wrapper = wrapper3.address
            writeConfig(config)
        }
    }

    if (config.RippleLockProxyImplementation === undefined) {
        // deploy rippleLockProxy Implementation
        console.log("\ndeploy RippleLockProxyImplementation ......".cyan);
        rippleLockProxyImpl = await RippleLockProxy.deploy(rippleName, rippleSymbol, rippleDeciamls);
        await rippleLockProxyImpl.deployed();
        console.log("RippleLockProxyImplementation deployed to:".green, rippleLockProxyImpl.address.blue);
        config.RippleLockProxyImplementation = rippleLockProxyImpl.address;
        writeConfig(config)
    } else {
        console.log("\nRippleLockProxyImplementation already deployed at".green, config.RippleLockProxyImplementation.blue)
        rippleLockProxyImpl = await RippleLockProxy.attach(config.RippleLockProxyImplementation) 
    }

    if (config.ProxyAdmin === undefined) {
        // deploy ProxyAdmin 
        console.log("\ndeploy ProxyAdmin ......".cyan);
        proxyAdmin = await ProxyAdmin.deploy();
        await proxyAdmin.deployed();
        console.log("ProxyAdmin deployed to:".green, proxyAdmin.address.blue);
        config.ProxyAdmin = proxyAdmin.address;
        writeConfig(config)
    } else {
        console.log("\nProxyAdmin already deployed at".green, config.ProxyAdmin.blue)
        proxyAdmin = await ProxyAdmin.attach(config.ProxyAdmin) 
    }

    if (config.RippleLockProxy === undefined) {
        // deploy rippleLockProxy
        console.log("\ndeploy RippleLockProxy......".cyan);
        tx = await cf.deploy(rippleSalt, rippleLockProxyImpl.address, proxyAdmin.address, rippleInitData);
        await tx.wait();
        let rippleLockProxyAddress = await cf.getDeploymentAddress(rippleSalt, deployer.address);
        rippleLockProxy = await RippleLockProxy.attach(rippleLockProxyAddress);
        console.log("RippleLockProxy deployed to:".green, rippleLockProxy.address.blue);
        config.RippleLockProxy = rippleLockProxy.address;
        writeConfig(config)
    } else {
        console.log("\nRippleLockProxy already deployed at".green, config.RippleLockProxy.blue)
        rippleLockProxy = await RippleLockProxy.attach(config.RippleLockProxy) 
    }
    let xrp = await rippleLockProxy.token();
    console.log("\nXRP-Token deployed at".green, xrp.blue);

    let alreadySetCCMPRipple = await rippleLockProxy.managerProxyContract();
    if (alreadySetCCMPRipple == ccm.address) {
        console.log("\nmanagerContract of ripple LockProxy already set".green);
    } else {
        // setup RippleLockProxy
        console.log("\nsetup RippleLockProxy ......".cyan);
        tx = await rippleLockProxy.setManager(ccm.address);
        await tx.wait();
        console.log("setManager Done".green);
    }

    let alreadyAddLockProxy3 = await wrapper3.isValidLockProxy(lockProxy.address);
    let alreadySetFeeCollector3 = await wrapper3.feeCollector();
    let alreadyAddRippleLockProxy3 = await wrapper3.isValidLockProxy(rippleLockProxy.address);
    console.log("\nsetup WrapperV3 ......".cyan);
    if (alreadyAddLockProxy3) {
        console.log("wrapper3 lockProxy already add".green);
    } else {
        // addLockProxy
        console.log("addLockProxy ......".cyan);
        tx = await wrapper3.addLockProxy(lockProxy.address);
        await tx.wait();
        console.log("addLockProxy Done".green);
    }
    if (alreadyAddRippleLockProxy3) {
        console.log("wrapper3 rippleLockProxy already add".green);
    } else {
        // addLockProxy
        console.log("add RippleLockProxy ......".cyan);
        tx = await wrapper3.addLockProxy(rippleLockProxy.address);
        await tx.wait();
        console.log("add RippleLockProxy Done".green);
    }
    if (alreadySetFeeCollector3 != "0x0000000000000000000000000000000000000000") {
        console.log("wrapper3 feeCollector already set".green);
    } else {
        // setFeeCollector
        console.log("setFeeCollector ......".cyan);
        tx = await wrapper3.setFeeCollector(deployer.address);
        await tx.wait();
        console.log("setFeeCollector Done".green);
    }

    // write config
    console.log("constract output:\n".cyan,config);
    await writeConfig(config)
    console.log("\nwrite config done\n".green);

    console.log("\nDone.\n".magenta);
}

async function updateConst(polyChainId, eccd, callerFactory) {
  
    fs.writeFileSync('./contracts/core/cross_chain_manager/logic/Const.sol', 
    'pragma solidity ^0.5.0;\n'+
    'contract Const {\n'+
    '    bytes constant ZionCrossChainManagerAddress = hex"0000000000000000000000000000000000001003"; \n'+
    // '    bytes constant ZionCrossChainManagerAddress = hex"5747C05FF236F8d18BB21Bc02ecc389deF853cae"; \n'+
    '    \n'+
    '    address constant EthCrossChainDataAddress = '+eccd+'; \n'+
    '    address constant EthCrossChainCallerFactoryAddress = '+callerFactory+'; \n'+
    '    uint constant chainId = '+polyChainId+'; \n}', 
    function(err) {
        if (err) {
            console.error(err);
            process.exit(1);
        }
    }); 
}

async function readConfig(networkName) {
    let jsonData
    try {
        jsonData = fs.readFileSync(configPath)
    } catch(err) {
        if (err.code == 'ENOENT') {
            createEmptyConfig()
            return
        }else{
            console.error(err);
            process.exit(1);
        }
    }
    if (jsonData === undefined) {
        return
    }
    var json=JSON.parse(jsonData.toString())
    if (json.Networks === undefined) {
        return
    }
    for (let i=0; i<json.Networks.length; i++) {
        if (json.Networks[i].Name == networkName) {
            return json.Networks[i]
        }
    }
    // console.error("network do not exisit in config".red);
    // process.exit(1);
}

async function writeConfig(networkConfig) {
    if (networkConfig.Name === undefined) {
        console.error("invalid network config".red);
        process.exit(1);
    }
    let data=fs.readFileSync(configPath,(err,data)=>{
        if (err) {
            console.error(err);
            process.exit(1);
        }else{
          previous=data.toString();
        }  
    });
    var json = JSON.parse(data.toString())  
    var writeIndex = json.Networks.length 
    for (let i=0; i<json.Networks.length; i++) {
        if (json.Networks[i].Name == networkConfig.Name) {
            writeIndex = i
            break
        }
    }
    json.Networks[writeIndex] = networkConfig
    var jsonConfig = JSON.stringify(json,null,"\t")
    try {
        fs.writeFileSync(configPath, jsonConfig);
    } catch (err) {
        console.error(err);
        process.exit(1);
    }
}

function createEmptyConfig() {
    var json = {Networks: []}
    var jsonConfig = JSON.stringify(json,null,"\t")
    try {
        fs.writeFileSync(configPath, jsonConfig);
    } catch (err) {
        console.error(err);
        process.exit(1);
    }
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });
