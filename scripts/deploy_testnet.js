const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const Web3 = require("web3");
const colors = require('colors');
const { deployContract } = require("ethereum-waffle");
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

  const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
  const CCM = await hre.ethers.getContractFactory("EthCrossChainManagerNoWhiteList");
  const CCMP = await hre.ethers.getContractFactory("EthCrossChainManagerProxy");
  const LockProxy = await hre.ethers.getContractFactory("LockProxy");
  const WrapperV2 = await hre.ethers.getContractFactory("PolyWrapper");
  
  // deploy EthCrossChainData
  console.log("\ndeploy EthCrossChainData ......".cyan);
  const eccd = await ECCD.deploy();
  await eccd.deployed();
  console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
  
  // deploy EthCrossChainManager
  console.log("\ndeploy EthCrossChainManager ......".cyan);
  const ccm = await CCM.deploy(eccd.address, polyId);
  await ccm.deployed();
  console.log("EthCrossChainManager deployed to:".green, ccm.address.blue);
  
  // deploy EthCrossChainManagerProxy
  console.log("\ndeploy EthCrossChainManagerProxy ......".cyan);
  const ccmp = await CCMP.deploy(ccm.address);
  await ccmp.deployed();
  console.log("EthCrossChainManagerProxy deployed to:".green, ccmp.address.blue);

  // transfer ownership
  console.log("\ntransfer eccd's ownership to ccm ......".cyan);
  tx = await eccd.transferOwnership(ccm.address);
  await tx.wait();
  console.log("ownership transferred".green);

  console.log("\ntransfer ccm's ownership to ccmp ......".cyan);
  tx = await ccm.transferOwnership(ccmp.address);
  await tx.wait();
  console.log("ownership transferred".green);

  // deploy LockProxy
  console.log("\ndeploy LockProxy ......".cyan);
  const lockproxy = await LockProxy.deploy();
  await lockproxy.deployed();
  console.log("LockProxy deployed to:".green, lockproxy.address.blue);

  // setup LockProxy
  console.log("\nsetup LockProxy ......".cyan);
  tx = await lockproxy.setManagerProxy(ccmp.address);
  await tx.wait();
  console.log("setManagerProxy Done".green);

  // deploy wrapper
  console.log("\ndeploy PolyWrapperV2 ......".cyan);
  const wrapper = await WrapperV2.deploy(deployer.address, polyId);
  await wrapper.deployed();
  console.log("PolyWrapperV2 deployed to:".green, wrapper.address.blue);
  
  // setup wrapper
  console.log("\nsetup WrapperV2 ......".cyan);
  tx = await wrapper.setFeeCollector(deployer.address);
  await tx.wait();
  console.log("setFeeCollector Done".green);
  tx = await wrapper.setLockProxy(lockproxy.address);
  await tx.wait();
  console.log("setLockProxy Done".green);

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
    // case 10: // optimistic-main
    //   return ;

    // testnet
    case 3: // ropsten-test
      return 2;
    case 4: // rinkeby
      return 402;
    case 5: // goerli
      return 502;
    case 42: // kovan
      return 302;
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
    case 69: // optimistic-test
      return 210;
    case 4002: // ftm-test
        return 208;
    case 43113: // avax-test
        return 209;

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