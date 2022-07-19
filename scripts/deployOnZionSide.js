const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './devConfig.json'

async function main() {
    [admin, creator] = await hre.ethers.getSigners();
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
        config.Deployer = admin.address
    }

    const nativeLockProxy = '0x7d79D936DA7833c7fe056eB450064f34A327DcA8';
    const Empty = await ethers.getContractFactory("empty");
    const LockProxy = await ethers.getContractFactory("LockProxy");
    const EthCrossChainData = await hre.ethers.getContractFactory("EthCrossChainData");
    const CallerFactory = await hre.ethers.getContractFactory("CallerFactoryWithAdmin");
    let EthCrossChainManagerImplementation = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const EthCrossChainManager = await hre.ethers.getContractFactory("EthCrossChainManager");
    const WrapperV1 = await hre.ethers.getContractFactory("PolyWrapperV1");
    const WrapperV2 = await hre.ethers.getContractFactory("PolyWrapperV2");
    let polyId = config.PolyChainID
    let empty
    let eccd
    let ccmi
    let ccm
    let cf
    let lockProxy
    let wrapper1
    let wrapper2

    console.log("\nDeploy contracts on chain with Poly_Chain_Id:".cyan, polyId);
    
    if (config.Empty === undefined) {
        // deploy Empty
        console.log("\ndeploy Empty ......".cyan);
        empty = await Empty.connect(creator).deploy();
        await empty.deployed();
        console.log("Empty deployed to:".green, empty.address.blue);
        config.Empty = empty.address
        writeConfig(config)
    } else {
        console.log("\nEmpty already deployed at".green, config.Empty.blue)
        empty = await Empty.attach(config.Empty) 
    }
    
    if (config.EthCrossChainManager === undefined) {
        // deploy EthCrossChainManager
        console.log("\ndeploy EthCrossChainManager ......".cyan);
        ccm = await EthCrossChainManager.connect(creator).deploy(empty.address,admin.address,'0x');
        await ccm.deployed();
        console.log("EthCrossChainManager deployed to:".green, ccm.address.blue);
        config.EthCrossChainManager = ccm.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManager already deployed at".green, config.EthCrossChainManager.blue)
        ccm = await EthCrossChainManager.attach(config.EthCrossChainManager) 
    }
    
    if (config.LockProxy === undefined) {
        // deploy LockProxy
        console.log("\ndeploy LockProxy ......".cyan);
        lockProxy = await LockProxy.connect(admin).deploy();
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
        eccd = await EthCrossChainData.connect(admin).deploy();
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
        cf = await CallerFactory.connect(admin).deploy([lockProxy.address]);
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
        ccmi = await EthCrossChainManagerImplementation.connect(admin).deploy();
        await ccmi.deployed();
        console.log("EthCrossChainManagerImplementation deployed to:".green, ccmi.address.blue);
        config.EthCrossChainManagerImplementation = ccmi.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManagerImplementation already deployed at".green, config.EthCrossChainManagerImplementation.blue)
        ccmi = await EthCrossChainManagerImplementation.attach(config.EthCrossChainManagerImplementation) 
    }
    
    let currentImplementation = await getProxyImplementation(ccm.address)
    if (currentImplementation != ccmi.address) {
        // setup EthCrossChainManager
        console.log("\nsetup EthCrossChainManager ......".cyan);
        tx = await ccm.connect(admin).upgradeTo(ccmi.address);
        await tx.wait();
        console.log("set".green);
    } else {
        console.log("\nEthCrossChainManager implementation already set".green);
    }

    let eccdOwner = await eccd.owner()
    // transfer ownership
    if (eccdOwner == ccm.address) {
        console.log("eccd ownership already transferred".green);
    } else {
        console.log("\ntransfer eccd's ownership to ccm ......".cyan);
        tx = await eccd.transferOwnership(ccm.address);
        await tx.wait();
        console.log("ownership transferred".green);
    }

    let alreadySetCCMP = await lockProxy.managerProxyContract();
    if (alreadySetCCMP == ccm.address) {
        console.log("managerProxyContract already set".green);
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
        wrapper1 = await WrapperV1.connect(admin).deploy(admin.address, polyId);
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
        tx = await wrapper1.setFeeCollector(admin.address);
        await tx.wait();
        console.log("setFeeCollector Done".green);
    }
    
    if (config.WrapperV2 === undefined) {
        // deploy WrapperV2
        console.log("\ndeploy WrapperV2 ......".cyan);
        wrapper2 = await WrapperV2.connect(admin).deploy(admin.address, polyId);
        await wrapper2.deployed();
        console.log("WrapperV2 deployed to:".green, wrapper2.address.blue);
        config.WrapperV2 = wrapper2.address
        config.Wrapper = wrapper2.address
        writeConfig(config)
    } else {
        console.log("\nWrapperV2 already deployed at".green, config.WrapperV2.blue)
        wrapper2 = await WrapperV2.attach(config.WrapperV2) 
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
        console.log("wrapper1 feeCollector already set".green);
    } else {
        // setFeeCollector
        console.log("setFeeCollector ......".cyan);
        tx = await wrapper2.setFeeCollector(admin.address);
        await tx.wait();
        console.log("setFeeCollector Done".green);
    }

    // write config
    console.log("constract output:\n".cyan,config);
    await writeConfig(config)
    console.log("\nwrite config done\n".green);

    console.log("\nDone.\n".magenta);
}

async function getProxyImplementation(contractAddr) {
    const IMPLEMENTATION_SLOT = '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc';
    let implementation = await hre.web3.eth.getStorageAt(contractAddr, IMPLEMENTATION_SLOT);
    return implementation
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
