const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

async function main() {
    [deployer, deployer2] = await hre.ethers.getSigners();
  
    // check if given networkId is registered
    await getPolyChainId().then((polyId) => {
      console.log("\nDeploy contracts on chain with Poly_Chain_Id:".cyan, polyId);
    }).catch((error) => {
      throw error;
    });

    const polyId = await getPolyChainId();
    const LockProxy = await ethers.getContractFactory("LockProxy");
    const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
    const CallerFactory = await hre.ethers.getContractFactory("CallerFactory");
    const CCM = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const CCMP = await hre.ethers.getContractFactory("EthCrossChainManager");
    const WrapperV1 = await hre.ethers.getContractFactory("PolyWrapperV1");
    const WrapperV2 = await hre.ethers.getContractFactory("PolyWrapperV2");

    console.log("\nStart , deployer:".cyan, deployer.address.blue);
    
    // deploy LockProxy
    console.log("\ndeploy LockProxy ......".cyan);
    const lockProxy = await LockProxy.deploy();
    await lockProxy.deployed();
    console.log("LockProxy deployed to:".green, lockProxy.address.blue);
  
    // deploy EthCrossChainData
    console.log("\ndeploy EthCrossChainData ......".cyan);
    const eccd = await ECCD.deploy();
    await eccd.deployed();
    console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
    
    // deploy CallerFactory
    console.log("\ndeploy CallerFactory ......".cyan);
    const cf = await CallerFactory.deploy([lockProxy.address]);
    await cf.deployed();
    console.log("CallerFactory deployed to:".green, cf.address.blue);
    
    // update Const.sol
    console.log("\nupdate Const.sol ......".cyan);
    await updateConst(eccd.address, cf.address);
    console.log("Const.sol updated".green);
    await hre.run('compile');
    
    // deploy EthCrossChainManagerImplementation
    console.log("\ndeploy EthCrossChainManagerImplementation ......".cyan);
    const ccm = await CCM.deploy();
    await ccm.deployed();
    console.log("EthCrossChainManagerImplementation deployed to:".green, ccm.address.blue);
    
    // deploy EthCrossChainManager
    console.log("\ndeploy EthCrossChainManager ......".cyan);
    const ccmp = await CCMP.deploy(ccm.address,deployer.address,'0x');
    await ccmp.deployed();
    console.log("EthCrossChainManager deployed to:".green, ccmp.address.blue);
  
    // transfer ownership
    console.log("\ntransfer eccd's ownership to ccm ......".cyan);
    tx = await eccd.transferOwnership(ccmp.address);
    await tx.wait();
    console.log("ownership transferred".green);

    // setup contracts
    console.log("\nsetup LockProxy ......".cyan);
    tx = await lockProxy.setManagerProxy(ccmp.address);
    await tx.wait();
    console.log("LockProxy set".green);

    // deploy wrapper
    console.log("\ndeploy PolyWrapperV1 ......".cyan);
    const wrapper1 = await WrapperV1.deploy(deployer.address, polyId);
    await wrapper1.deployed();
    console.log("PolyWrapperV1 deployed to:".green, wrapper1.address.blue);
  
    console.log("\ndeploy PolyWrapperV2 ......".cyan);
    const wrapper2 = await WrapperV2.deploy(deployer.address, polyId);
    await wrapper2.deployed();
    console.log("PolyWrapperV2 deployed to:".green, wrapper2.address.blue);
    
    // setup wrapper
    console.log("\nsetup WrapperV1 ......".cyan);
    tx = await wrapper1.setFeeCollector(deployer.address);
    await tx.wait();
    console.log("setFeeCollector Done".green);
    tx = await wrapper1.setLockProxy(lockProxy.address);
    await tx.wait();
    console.log("setLockProxy Done".green);
  
    console.log("\nsetup WrapperV2 ......".cyan);
    tx = await wrapper2.setFeeCollector(deployer.address);
    await tx.wait();
    console.log("setFeeCollector Done".green);
    tx = await wrapper2.setLockProxy(lockProxy.address);
    await tx.wait();
    console.log("setLockProxy Done".green);

    console.log("\nDone.\n".magenta);
}

async function updateConst(eccd, callerFactory) {
    const polyChainId = await getPolyChainId();
  
    await fs.writeFile('./contracts/core/cross_chain_manager/logic/Const.sol', 
    'pragma solidity ^0.5.0;\n'+
    'contract Const {\n'+
    '    bytes constant ZionCrossChainManagerAddress = hex"5747C05FF236F8d18BB21Bc02ecc389deF853cae"; \n'+
    '    bytes constant ZionValidaterManagerAddress = hex"A4Bf827047a08510722B2d62e668a72FCCFa232C"; \n'+
    '    \n'+
    '    address constant EthCrossChainDataAddress = '+eccd+'; \n'+
    '    address constant EthCrossChainCallerFactoryAddress = '+callerFactory+'; \n'+
    '    uint constant chainId = '+polyChainId+'; \n}', 
    function(err) {
      if (err) {
          console.error(err);
          process.exit(1);
      }
    }); 
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

    // testnet
    case 3: // ropsten
      return 202;
    case 4: // rinkeby
      return 402;
    case 5: // goerli
      return 502;
    case 42: // kovan
      return 302;
    case 97: // bsc-test
      return 1000006;
    case 256: // heco-test
      return 7;
    case 80001: // polygon-test
      return 216;
    case 65: // ok-test
      return 200;
    case 421611: // arbitrum-test
      return 300;
    case 4002: // ftm-test
      return 400;
    case 43113: // avax-test
      return 500;
    case 77: // xdai-test
      return 600;
    case 69: // op-test
      return 200;
    case 101: // plt-test
      return 208;
    case 31091: // curve-test
      return 210;

    // hardhat devnet
    case 31337:
      return 77777;

    // unknown chainid
    default: 
      throw new Error("fail to get Poly_Chain_Id, unknown Network_Id: "+chainId);
  }
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
      console.error(err)
      process.exit(1)
  });
