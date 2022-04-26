const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './contractConfig.json'
var whiteListConfig = './whiteListConfig.json'

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

    const EthCrossChainData = await hre.ethers.getContractFactory("EthCrossChainData");
    const CallerFactory = await hre.ethers.getContractFactory("CallerFactoryWithAdmin");
    const EthCrossChainManager = await hre.ethers.getContractFactory("EthCrossChainManager");
    let polyId = config.PolyChainID
    let eccd
    let ccmi
    let ccm
    let cf

    console.log("\nDeploy contracts on chain with Poly_Chain_Id:".cyan, polyId);
    
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
        var whiteList = []
        await readWhiteList(hre.network.name).then((res) => {
            if (res !== undefined) {
                whiteList = res
            }
        }).catch();
        cf = await CallerFactory.deploy(whiteList);
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
        EthCrossChainManagerImplementation = await hre.ethers.getContractFactory("OntEvmCrossChainManagerImplementation");
        ccmi = await EthCrossChainManagerImplementation.deploy();
        await ccmi.deployed();
        console.log("EthCrossChainManagerImplementation deployed to:".green, ccmi.address.blue);
        config.EthCrossChainManagerImplementation = ccmi.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManagerImplementation already deployed at".green, config.EthCrossChainManagerImplementation.blue)
        EthCrossChainManagerImplementation = await hre.ethers.getContractFactory("OntEvmCrossChainManagerImplementation");
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
        console.log("eccd ownership already transferred".green);
    } else {
        console.log("\ntransfer eccd's ownership to ccm ......".cyan);
        tx = await eccd.transferOwnership(ccm.address);
        await tx.wait();
        console.log("ownership transferred".green);
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
    '    bytes constant ZionCrossChainManagerAddress = hex"5747C05FF236F8d18BB21Bc02ecc389deF853cae"; \n'+
    '    \n'+
    '    address constant EthCrossChainDataAddress = '+eccd+'; \n'+
    '    address constant EthCrossChainCallerFactoryAddress = '+callerFactory+'; \n'+
    '    uint constant chainId = '+polyChainId+'; \n'+
    '    address constant EVENT_WITNESS = 0x2b1143484bf5097A29678FD9592f75FE4639CA08; // mainnet \n}', 
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

async function readWhiteList(networkName) {
    let jsonData
    try {
        jsonData = fs.readFileSync(whiteListConfig)
    } catch(err) {
        return
    }
    if (jsonData === undefined) {
        return
    }
    var json=JSON.parse(jsonData.toString())
    if (json.WhiteLists === undefined) {
        return
    }
    for (let i=0; i<json.WhiteLists.length; i++) {
        if (json.WhiteLists[i].Name == networkName) {
            return json.WhiteLists[i].WhiteList
        }
    }
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });
