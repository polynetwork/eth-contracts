const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './polyConfig.json'

async function main() {
    [admin] = await hre.ethers.getSigners();
    var config = {}
    await readConfig(hre.network.name).then((netConfig) => {
        if (netConfig !== undefined) {
            config = netConfig
        } else {
            console.error("unknown network: network config does not exisit".red);
            process.exit(1)
        }
    }).catch((err) => {
        console.error(err);
        process.exit(1);
    });
    if (config.EthCrossChainManager === undefined) {
        console.error("invalid network config: EthCrossChainManager does not exisit".red);  
    }

    const CCM = await hre.ethers.getContractFactory("EthCrossChainManager");
    let ccm = await CCM.attach(config.EthCrossChainManager)
    
    console.log("\nset white list on chain with Poly_Chain_Id:".cyan, config.PolyChainID);

    let fromContractWls = []
    let contractMethodWls = []
    
    if (config.LockProxy !== undefined) {
        let name = "LockProxy"
        let contract = config.LockProxy
        let method = ["0x756e6c6f636b"] // unlock

        console.log("\ncheck ".cyan+name.cyan+" WhiteList ......".cyan);
        let flag1 = await ccm.whiteListFromContract(contract)
        let flag2 = await ccm.whiteListContractMethodMap(contract, method[0])
        if (flag1 == true) {
            console.log(name.green + " already in fromContractWhiteList".green);
        } else {
            console.log(name.blue + " will be add to fromContractWhiteList".blue);
            fromContractWls.push(contract)
        }
        if (flag2 == true) {
            console.log(name.green + " already in contractMethodWhiteList".green);
        } else {
            console.log(name.blue + " will be add to contractMethodWhiteList".blue);
            abiData = hre.web3.eth.abi.encodeParameters(['address','bytes[]'], [contract, method])
            contractMethodWls.push(abiData)
        }
    }
    
    if (config.LockProxyPip4 !== undefined) {
        let name = "LockProxyPip4"
        let contract = config.LockProxyPip4
        let method = ["0x756e6c6f636b"] // unlock

        console.log("\ncheck ".cyan+name.cyan+" WhiteList ......".cyan);
        let flag1 = await ccm.whiteListFromContract(contract)
        let flag2 = await ccm.whiteListContractMethodMap(contract, method[0])
        if (flag1 == true) {
            console.log(name.green + " already in fromContractWhiteList".green);
        } else {
            console.log(name.blue + " will be add to fromContractWhiteList".blue);
            fromContractWls.push(contract)
        }
        if (flag2 == true) {
            console.log(name.green + " already in contractMethodWhiteList".green);
        } else {
            console.log(name.blue + " will be add to contractMethodWhiteList".blue);
            abiData = hre.web3.eth.abi.encodeParameters(['address','bytes[]'], [contract, method])
            contractMethodWls.push(abiData)
        }
    }
    
    if (config.NftProxy !== undefined) {
        let name = "NftProxy"
        let contract = config.NftProxy
        let method = ["0x756e6c6f636b"] // unlock
        
        console.log("\ncheck ".cyan+name.cyan+" WhiteList ......".cyan);
        let flag1 = await ccm.whiteListFromContract(contract)
        let flag2 = await ccm.whiteListContractMethodMap(contract, method[0])
        if (flag1 == true) {
            console.log(name.green + " already in fromContractWhiteList".green);
        } else {
            console.log(name.blue + " will be add to fromContractWhiteList".blue);
            fromContractWls.push(contract)
        }
        if (flag2 == true) {
            console.log(name.green + " already in contractMethodWhiteList".green);
        } else {
            console.log(name.blue + " will be add to contractMethodWhiteList".blue);
            abiData = hre.web3.eth.abi.encodeParameters(['address','bytes[]'], [contract, method])
            contractMethodWls.push(abiData)
        }
    }
    
    if (config.Swapper !== undefined) {
        let name = "Swapper"
        let contract = config.Swapper
        let method = ["0x6164645f6c6971756964697479", "0x72656d6f76655f6c6971756964697479", "0x73776170"]
        // ["add_liquidity", "remove_liquidity", "swap"]
        
        console.log("\ncheck ".cyan+name.cyan+" WhiteList ......".cyan);
        let flag1 = await ccm.whiteListFromContract(contract)
        let flag2 = await ccm.whiteListContractMethodMap(contract, method[0])
        if (flag1 == true) {
            console.log(name.green + " already in fromContractWhiteList".green);
        } else {
            console.log(name.blue + " will be add to fromContractWhiteList".blue);
            fromContractWls.push(contract)
        }
        if (flag2 == true) {
            console.log(name.green + " already in contractMethodWhiteList".green);
        } else {
            console.log(name.blue + " will be add to contractMethodWhiteList".blue);
            abiData = hre.web3.eth.abi.encodeParameters(['address','bytes[]'], [contract, method])
            contractMethodWls.push(abiData)
        }
    }

    if (fromContractWls.length != 0) {
        console.log("\nsetFromContractWhiteList ......".cyan);
        tx = await ccm.setFromContractWhiteList(fromContractWls);
        await tx.wait();
        console.log("setFromContractWhiteList done".green);
    }

    if (contractMethodWls.length != 0) {
        console.log("\nsetContractMethodWhiteList ......".cyan);
        tx = await ccm.setContractMethodWhiteList(contractMethodWls);
        await tx.wait();
        console.log("setContractMethodWhiteList done".green);
    }

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

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });