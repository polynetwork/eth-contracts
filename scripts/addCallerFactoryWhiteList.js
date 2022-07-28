const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './contractConfig.json'
var whiteListConfig = './whiteListConfig.json'

async function main() {
    console.log("\nSet whiteList on :".cyan, hre.network.name);

    const CallerFactory = await hre.ethers.getContractFactory("CallerFactoryWithAdmin");
    var CallerFactoryAddress = await readCallerProxyAddress(hre.network.name)
    var whiteList = await readWhiteList(hre.network.name)
    let factory = await CallerFactory.attach(CallerFactoryAddress)

    let tx = await factory.setChildren(whiteList, true)
    await tx.wait()

    console.log("\nDone.\n".magenta);
}

async function readWhiteList(networkName) {
    let jsonData
    try {
        jsonData = fs.readFileSync(whiteListConfig)
    } catch(err) {
        console.error("fail to read whiteList config file. ",err);
        process.exit(1);
    }
    if (jsonData === undefined) {
        console.error("empty whiteList config");
        process.exit(1);
    }
    var json=JSON.parse(jsonData.toString())
    if (json.WhiteLists === undefined) {
        console.error("invalid whiteList config");
        process.exit(1);
    }
    for (let i=0; i<json.WhiteLists.length; i++) {
        if (json.WhiteLists[i].Name == networkName) {
            return json.WhiteLists[i].WhiteList
        }
    }
    console.error("network do not exisit in config".red);
    process.exit(1);
}

async function readCallerProxyAddress(networkName) {
    let jsonData
    try {
        jsonData = fs.readFileSync(configPath)
    } catch(err) {
        console.error("fail to read config file. ",err);
        process.exit(1);
    }
    if (jsonData === undefined) {
        console.error("empty config");
        process.exit(1);
    }
    var json=JSON.parse(jsonData.toString())
    if (json.Networks === undefined) {
        console.error("invalid config");
        process.exit(1);
    }
    for (let i=0; i<json.Networks.length; i++) {
        if (json.Networks[i].Name == networkName) {
            return json.Networks[i].CallerFactory
        }
    }
    console.error("network do not exisit in config".red);
    process.exit(1);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });
