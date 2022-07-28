const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './contractConfig.json'

async function main() {
    [admin] = await hre.ethers.getSigners();
    const netwrokId = await hre.web3.eth.getChainId();
    var config
    await readConfig(hre.network.name).then((netConfig) => {
        config = netConfig
    }).catch((err) => {
        console.error(err);
        process.exit(1);
    });
    if (config.PolyChainID === undefined) {
        console.error("invalid network: invalid PolyChainID".red);
        process.exit(1);
    }
    if (config.EthCrossChainData === undefined) {
        console.error("invalid network: invalid EthCrossChainData".red);
        process.exit(1);
    }
    if (config.CallerFactory === undefined) {
        console.error("invalid network: invalid CallerFactory".red);
        process.exit(1);
    }
    if (config.EthCrossChainManager === undefined) {
        console.error("invalid network: invalid EthCrossChainManager".red);
        process.exit(1);
    }

    var CCM = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const CCMP = await hre.ethers.getContractFactory("EthCrossChainManager");
    const ccmp = await CCMP.attach(config.EthCrossChainManager)

    console.log("\nUpdate ccm on chain with Poly_Chain_Id:".cyan, config.PolyChainID.toString());
    console.log("\nStart update, admin:".cyan, admin.address.blue);
    
    // update Const.sol
    console.log("\nupdate Const.sol ......".cyan);
    await updateConst(config.PolyChainID, config.EthCrossChainData, config.CallerFactory);
    console.log("Const.sol updated".green);
    await hre.run('compile');
    
    // deploy new EthCrossChainManagerImplementation
    console.log("\ndeploy new EthCrossChainManagerImplementation ......".cyan);
    CCM = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const ccmi = await CCM.deploy();
    await ccmi.deployed();
    console.log("New EthCrossChainManagerImplementation deployed to:".green, ccmi.address.blue);
    
    // update EthCrossChainManager
    console.log("\nupdate EthCrossChainManager ......".cyan);
    tx = await ccmp.upgradeTo(ccmi.address);
    await tx.wait();
    console.log("EthCrossChainManager updated");

    // write config
    console.log("\nwrite config ......".cyan);
    config.EthCrossChainManagerImplementation = ccmi.address
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

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });