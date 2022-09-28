const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const Web3 = require("web3");
const colors = require('colors');
hre.web3 = new Web3(hre.network.provider);

async function main() {

  [deployer] = await hre.ethers.getSigners();
  let polyId;
  
  // check if given networkId is registered
  await getPolyChainId().then((_polyId) => {
    polyId = _polyId;
    console.log("\nDeploy EthCrossChainManager on chain with Poly_Chain_Id:".cyan, _polyId);
  }).catch((error) => {
    throw error;
  });;

  console.log("Start , deployer:".cyan, deployer.address.blue);
  
  // deploy EthCrossChainData
  console.log("\ndeploy EthCrossChainData ......".cyan);
  const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
  const eccd = await ECCD.deploy();
  await eccd.deployed();
  console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
  
  // deploy EthCrossChainManager
  console.log("\ndeploy EthCrossChainManager ......".cyan);
  const CCM = await hre.ethers.getContractFactory("EthCrossChainManager");
  const ccm = await CCM.deploy(eccd.address, polyId, [], []);
  await ccm.deployed();
  console.log("EthCrossChainManager deployed to:".green, ccm.address.blue);
  
  // deploy EthCrossChainManagerProxy
  console.log("\ndeploy EthCrossChainManagerProxy ......".cyan);
  const CCMP = await hre.ethers.getContractFactory("EthCrossChainManagerProxy");
  const ccmp = await CCMP.deploy(ccm.address);
  await ccmp.deployed();
  console.log("EthCrossChainManagerProxy deployed to:".green, ccmp.address.blue);

  // transfer ownership
  console.log("\ntransfer eccd's ownership to ccm ......".cyan);
  await eccd.transferOwnership(ccm.address);
  console.log("ownership transferred".green);

  console.log("\ntransfer ccm's ownership to ccmp ......".cyan);
  await ccm.transferOwnership(ccmp.address);
  console.log("ownership transferred".green);

  console.log("\nDone.\n".magenta);

}

async function getPolyChainId() {
  const chainId = await hre.web3.eth.getChainId();
  switch (chainId) {
    
    // mainnet
    case 1: // eth-main
      return 2;
    case 56: // bsc-main
      return 6;
    case 128: // heco-main
      return 7;
    case 137: // polygon-main
      return 17;
    case 66: // ok-main
      return 12;
    case 1718: // plt-main
      return 8;
    case 42161: // arbitrum-main
      return 19;

    // testnet
    case 3: // eth-test
      return 2;
    case 97: // bsc-test
      return 79;
    case 256: // heco-test
      return 7;
    case 80001: // polygon-test
      return 202;
    case 65: // ok-test
      return 200;
    case 101: // plt-test
      return 107;
    case 421611: // arbitrum-test
      return 205;

    // hardhat devnet
    case 31337:
      return 77777;

    // unknown chainid
    default: 
      throw new Error("fail to get Poly_Chain_Id, unknown Network_Id: "+chainId);
  }
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });