const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

var configPath = './zionDevConfig.json'

async function main() {
    console.log(hre.config.networks[hre.network.name].polyId);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });