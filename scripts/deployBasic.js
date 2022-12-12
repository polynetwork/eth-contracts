const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './polyConfig.json'

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

    const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
    const CCM = await hre.ethers.getContractFactory("EthCrossChainManager");
    const CCMP = await hre.ethers.getContractFactory("EthCrossChainManagerProxy");
    let polyId = config.PolyChainID
    let eccd
    let ccm
    let ccmp
    
    console.log("\nDeploy basic contracts on chain with Poly_Chain_Id:".cyan, polyId);
    
    if (config.EthCrossChainData === undefined) {
        // deploy EthCrossChainData
        console.log("\ndeploy EthCrossChainData ......".cyan);
        eccd = await ECCD.deploy();
        await eccd.deployed();
        console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
        config.EthCrossChainData = eccd.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainData already deployed at".green, config.EthCrossChainData.blue)
        eccd = await ECCD.attach(config.EthCrossChainData) 
    }

    if (config.EthCrossChainManager === undefined) {
        // deploy EthCrossChainManager
        console.log("\ndeploy EthCrossChainManager ......".cyan);
        ccm = await CCM.deploy(eccd.address, polyId, [], []);
        await ccm.deployed();
        console.log("EthCrossChainManager deployed to:".green, ccm.address.blue);
        config.EthCrossChainManager = ccm.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManager already deployed at".green, config.EthCrossChainManager.blue)
        ccm = await CCM.attach(config.EthCrossChainManager) 
    }

    if (config.EthCrossChainManagerProxy === undefined) {
        // deploy EthCrossChainManagerProxy
        console.log("\ndeploy EthCrossChainManagerProxy ......".cyan);
        ccmp = await CCMP.deploy(ccm.address);
        await ccmp.deployed();
        console.log("EthCrossChainManagerProxy deployed to:".green, ccmp.address.blue);
        config.EthCrossChainManagerProxy = ccmp.address
        writeConfig(config)
    } else {
        console.log("\nEthCrossChainManagerProxy already deployed at".green, config.EthCrossChainManagerProxy.blue)
        ccmp = await CCMP.attach(config.EthCrossChainManagerProxy) 
    }

    let eccdOwner = await eccd.owner()
    let ccmOwner = await ccm.owner()
    // transfer ownership
    if (eccdOwner == ccm.address) {
        console.log("eccd ownership already transferred".green);
    } else {
        console.log("\ntransfer eccd's ownership to ccm ......".cyan);
        tx = await eccd.transferOwnership(ccm.address);
        await tx.wait();
        console.log("ownership transferred".green);
    }
    if (ccmOwner == ccmp.address) {
        console.log("ccm ownership already transferred".green);
    } else {
        console.log("\ntransfer ccm's ownership to ccmp ......".cyan);
        tx = await ccm.transferOwnership(ccmp.address);
        await tx.wait();
        console.log("ownership transferred".green);
    }

    // config
    // config.Name = hre.network.name
    // config.PolyChainID = hre.config.networks[config.Name].polyId
    // config.Provider = hre.config.networks[config.Name].url
    // config.Deployer = deployer.address
    // config.EthCrossChainData = eccd.address
    // config.EthCrossChainManager = ccm.address
    // config.EthCrossChainManagerProxy = ccmp.address
    // config.LockProxy = lockProxy.address
    // config.LockProxyPip4 = lockProxyPip4.address
    // config.Swapper = swapper.address
    // config.Wrapper = wrapper2.address
    // config.WrapperV1 = wrapper1.address
    // config.WrapperV2 = wrapper2.address
    // config.WrapperV3 = wrapper3.address
    // config.PrivateKeyNo = 1
    console.log("constract output:\n".cyan,config);
    await writeConfig(config)
    console.log("\nwrite config done\n".green);

    console.log("\nDone.\n".magenta);
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