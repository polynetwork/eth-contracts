const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './zionDevConfig.json'

async function main() {
    [deployer, deployer2] = await hre.ethers.getSigners();
    const netwrokId = await hre.web3.eth.getChainId();
    var config
    await readConfig(netwrokId).then((netConfig) => {
        config = netConfig
    }).catch((err) => {
        console.error(err);
        process.exit(1);
    });
    if (config.PolyChainID === undefined) {
        console.error("unknown network: invalid PolyChainID".red);
        process.exit(1);
    }

    const LockProxy = await ethers.getContractFactory("LockProxy");
    const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
    const CallerFactory = await hre.ethers.getContractFactory("CallerFactory");
    const CCM = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const CCMP = await hre.ethers.getContractFactory("EthCrossChainManager");
    const WrapperV1 = await hre.ethers.getContractFactory("PolyWrapperV1");
    const WrapperV2 = await hre.ethers.getContractFactory("PolyWrapperV2");

    console.log("\nStart , deployer:".cyan, deployer.address.blue);
    
    // deploy LockProxy
    console.log("\ndeploy LockProxy ......".cyan);
    const lockProxy = await LockProxy.deploy();
    await lockProxy.deployed();
    console.log("LockProxy deployed to:".green, lockProxy.address.blue);
  
    // deploy EthCrossChainData
    console.log("\ndeploy EthCrossChainData ......".cyan);
    const eccd = await ECCD.deploy();
    await eccd.deployed();
    console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
    
    // deploy CallerFactory
    console.log("\ndeploy CallerFactory ......".cyan);
    const cf = await CallerFactory.deploy([lockProxy.address]);
    await cf.deployed();
    console.log("CallerFactory deployed to:".green, cf.address.blue);
    
    // update Const.sol
    console.log("\nupdate Const.sol ......".cyan);
    await updateConst(config.PolyChainID, eccd.address, cf.address);
    console.log("Const.sol updated".green);
    await hre.run('compile');
    
    // deploy EthCrossChainManagerImplementation
    console.log("\ndeploy EthCrossChainManagerImplementation ......".cyan);
    const ccm = await CCM.deploy();
    await ccm.deployed();
    console.log("EthCrossChainManagerImplementation deployed to:".green, ccm.address.blue);
    
    // deploy EthCrossChainManager
    console.log("\ndeploy EthCrossChainManager ......".cyan);
    const ccmp = await CCMP.deploy(ccm.address,deployer.address,'0x');
    await ccmp.deployed();
    console.log("EthCrossChainManager deployed to:".green, ccmp.address.blue);
  
    // transfer ownership
    console.log("\ntransfer eccd's ownership to ccm ......".cyan);
    tx = await eccd.transferOwnership(ccmp.address);
    await tx.wait();
    console.log("ownership transferred".green);

    // setup contracts
    console.log("\nsetup LockProxy ......".cyan);
    tx = await lockProxy.setManagerProxy(ccmp.address);
    await tx.wait();
    console.log("LockProxy set".green);

    // deploy wrapper
    console.log("\ndeploy PolyWrapperV1 ......".cyan);
    const wrapper1 = await WrapperV1.deploy(deployer.address, config.PolyChainID);
    await wrapper1.deployed();
    console.log("PolyWrapperV1 deployed to:".green, wrapper1.address.blue);
  
    console.log("\ndeploy PolyWrapperV2 ......".cyan);
    const wrapper2 = await WrapperV2.deploy(deployer.address, config.PolyChainID);
    await wrapper2.deployed();
    console.log("PolyWrapperV2 deployed to:".green, wrapper2.address.blue);
    
    // setup wrapper
    console.log("\nsetup WrapperV1 ......".cyan);
    tx = await wrapper1.setFeeCollector(deployer.address);
    await tx.wait();
    console.log("setFeeCollector Done".green);
    tx = await wrapper1.setLockProxy(lockProxy.address);
    await tx.wait();
    console.log("setLockProxy Done".green);
  
    console.log("\nsetup WrapperV2 ......".cyan);
    tx = await wrapper2.setFeeCollector(deployer.address);
    await tx.wait();
    console.log("setFeeCollector Done".green);
    tx = await wrapper2.setLockProxy(lockProxy.address);
    await tx.wait();
    console.log("setLockProxy Done".green);

    // write config
    console.log("\nwrite config ......".cyan);
    config.Provider = hre.config.networks[config.Name].url
    config.Deployer = deployer.address
    config.EthCrossChainData = eccd.address
    config.EthCrossChainManagerImplementation = ccm.address
    config.EthCrossChainManager = ccmp.address
    config.CallerFactory = cf.address
    config.LockProxy = lockProxy.address
    config.WrapperV1 = wrapper1.address
    config.Wrapper = wrapper2.address
    console.log("constract output:\n".cyan,config);
    await writeConfig(config)
    console.log("\nwrite config done\n".green);

    console.log("\nDone.\n".magenta);
}

async function updateConst(polyChainId, eccd, callerFactory) {
  
    fs.writeFileSync('./contracts/core/cross_chain_manager/logic/Const.sol', 
    'pragma solidity ^0.5.0;\n'+
    'contract Const {\n'+
    '    bytes constant ZionCrossChainManagerAddress = hex"5747C05FF236F8d18BB21Bc02ecc389deF853cae"; \n'+
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

async function readConfig(networkId) {
    let data=fs.readFileSync(configPath,(err,data)=>{
        if (err) {
            console.error(err);
            process.exit(1);
        }else{
          previous=data.toString();
        }  
    });
    var json=JSON.parse(data.toString())
    for (let i=0; i<json.Networks.length; i++) {
        if (json.Networks[i].NetworkId == networkId) {
            return json.Networks[i]
        }
    }
    console.error("network do not exisit in config".red);
    process.exit(1);
}

async function writeConfig(networkConfig) {
    if (networkConfig.NetworkId === undefined) {
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
    var writeIndex = json.Networks.length + 1
    for (let i=0; i<json.Networks.length; i++) {
        if (json.Networks[i].NetworkId == networkConfig.NetworkId) {
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

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });