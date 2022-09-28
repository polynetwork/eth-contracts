const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const Web3 = require("web3");
const colors = require('colors');
const { expect } = require("chai");
hre.web3 = new Web3(hre.network.provider);

async function main() {

  [deployer] = await hre.ethers.getSigners();

  let ccmp_addr;
  let old_ccm_addr;
  let ccd_addr;
  let poly_id

  try {
    ccmp_addr = getCCMP();
  } catch(error) {
    throw error;
  }

  const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
  const CCM = await hre.ethers.getContractFactory("EthCrossChainManager");
  const SUSPEND_CCM = await hre.ethers.getContractFactory("EthCrossChainManagerForUpgrade");
  const CCMP = await hre.ethers.getContractFactory("EthCrossChainManagerProxy");
  
  let ccmp = await CCMP.attach(ccmp_addr);

  old_ccm_addr = await ccmp.getEthCrossChainManager();
  let old_ccm = await CCM.attach(old_ccm_addr);

  poly_id = await old_ccm.chainId();
  ccd_addr = await old_ccm.EthCrossChainDataAddress();
  let eccd = await ECCD.attach(ccd_addr);

  console.log("Upgrade ".cyan, hre.network.name.blue);

  // deploy EthCrossChainManager
  console.log("\ndeploy new EthCrossChainManager ......".cyan);
  let ccm = await SUSPEND_CCM.deploy(eccd.address, poly_id);
  await ccm.deployed();
  console.log("New EthCrossChainManager deployed to:".green, ccm.address.blue);

  // transfer ownership
  console.log("\ntransfer new ccm's ownership to ccmp ......".cyan);
  tx = await ccm.transferOwnership(ccmp.address);
  await tx.wait();
  console.log("ownership transferred".green);

  // paused
  console.log("\nCheck if paused and pause if not".cyan);
  let paused_1 = await ccmp.paused();
  let paused_2 = await old_ccm.paused();
  if (!paused_1 && paused_2) {
    tx = await ccmp.pause();
    await tx.wait();
  } else if (!paused_1 && !paused_2) {
    tx = await ccmp.pauseEthCrossChainManager();
    await tx.wait();
  } else if (paused_1 && !paused_2) {
    tx = await ccmp.unpause();
    await tx.wait();
    tx = await ccmp.pauseEthCrossChainManager();
    await tx.wait();
  }
  expect(await ccmp.paused()).to.equal(true);
  expect(await old_ccm.paused()).to.equal(true);
  console.log("all paused".green);

  // update to new
  console.log("\nUpdate to new ccm".cyan);
  tx = await ccmp.upgradeEthCrossChainManager(ccm.address);
  await tx.wait();
  console.log("updated".green);

  // unpaused
  console.log("\nUnpaused all".cyan);
  tx = await ccmp.unpauseEthCrossChainManager();
  await tx.wait();
  expect(await ccmp.paused()).to.equal(false);
  expect(await ccm.paused()).to.equal(false);
  console.log("unpaused".green);
 
  console.log("\nDone.\n".magenta);

}

async function getCCMP() {
    const chainId = await hre.web3.eth.getChainId();
  switch (chainId) {
    // testnet
    case 97: // bsc-test
      return "0x7d93fb660776751702d049aDB93968e309315C51";
    // case 256: // heco-test
    // case 80001: // polygon-test

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