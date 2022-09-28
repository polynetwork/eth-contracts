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

    [fromContractWls, contractMethodWls] = getWL()

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

function getWL() {
    let fromContractWLs = ["0x2b07784826146A1Fd2e86D3f4697000B6b97a153","0xcdd05Ebaa1FA06a5cE9eB67663aE9Ec78B37bd5B"]
    // something like [address, address, ...]
    // e.g. ["0x2b07784826146A1Fd2e86D3f4697000B6b97a153","0xcdd05Ebaa1FA06a5cE9eB67663aE9Ec78B37bd5B"]
    
    let contractMethodWLs = [["0x2b07784826146A1Fd2e86D3f4697000B6b97a153",["0x756e6c6f636b"]], ["0xcdd05Ebaa1FA06a5cE9eB67663aE9Ec78B37bd5B",["0x6164645f6c6971756964697479", "0x72656d6f76655f6c6971756964697479", "0x73776170"]]]
    // something like [[address,[bytes]], [[address,[bytes], ...]
    // e.g. [["0x2b07784826146A1Fd2e86D3f4697000B6b97a153",["0x756e6c6f636b"]], ["0xcdd05Ebaa1FA06a5cE9eB67663aE9Ec78B37bd5B",["0x6164645f6c6971756964697479", "0x72656d6f76655f6c6971756964697479", "0x73776170"]]]
    
    let contractMethodWLData = []
    for (let i=0;i<contractMethodWLs.length;i++) {
        contractMethodWLData.push(hre.web3.eth.abi.encodeParameters(['address','bytes[]'], contractMethodWLs[i]))
    }
    return [fromContractWLs, contractMethodWLData]
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