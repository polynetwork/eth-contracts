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

  await hre.run('compile');
  
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

  // deploy lockProxy
  console.log("\ndeploy LockProxy ......".cyan);
  const LockProxy = await hre.ethers.getContractFactory("LockProxy");
  const lockproxy = await LockProxy.deploy();
  await lockproxy.deployed();
  await lockproxy.setManagerProxy(ccmp.address);
  console.log("LockProxy deployed to:".green, lockproxy.address.blue);

  // deploy swapper
  console.log("\ndeploy Swapper ......".cyan);
  const Swapper = await hre.ethers.getContractFactory("Swapper");
  // TODO config here for different chain
  const swapper = await Swapper.deploy(deployer.address, polyId, 10, lockproxy.address, ccmp.address, '0x82af49447d8a07e3bd95bd0d56f35241523fbab1', '0x34d4a23A1FC0C694f0D74DDAf9D8d564cfE2D430');
  await swapper.deployed();
  console.log("Swapper deployed to:".green, swapper.address.blue);

  // set whiteList
  console.log("\nset WhiteList ......".cyan);
  await ccm.setFromContractWhiteList([swapper.address, lockproxy.address]);
  await ccm.setContractMethodWhiteList(['0x000000000000000000000000'+lockproxy.address.slice(2)+'0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000006756e6c6f636b0000000000000000000000000000000000000000000000000000']);
  console.log("whiteList set".green);


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

