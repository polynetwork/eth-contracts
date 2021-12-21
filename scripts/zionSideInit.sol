const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

async function main() {
    [admin, creator] = await hre.ethers.getSigners();
  
    // check if given networkId is registered
    await getPolyChainId().then((polyId) => {
      console.log("\nDeploy contracts on chain with Poly_Chain_Id:".cyan, polyId);
    }).catch((error) => {
      throw error;
    });
    
    const nativeLockProxy = '0x7d79D936DA7833c7fe056eB450064f34A327DcA8';
    const polyId = await getPolyChainId();

    const Empty = await ethers.getContractFactory("empty");
    const CCMP = await hre.ethers.getContractFactory("EthCrossChainManager");
    const LockProxy = await ethers.getContractFactory("LockProxy");
    const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
    const CallerFactory = await hre.ethers.getContractFactory("CallerFactory");
    const CCM = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
    const WrapperV1 = await hre.ethers.getContractFactory("PolyWrapperV1");
    const WrapperV2 = await hre.ethers.getContractFactory("PolyWrapperV2");

    console.log("\nStart , admin:".cyan, admin.address.blue);

    // deploy empty
    console.log("\ndeploy Empty ......".cyan);
    const empty = await Empty.connect(creator).deploy();
    await empty.deployed();
    console.log("Empty deployed to:".green, empty.address.blue);
    
    // deploy EthCrossChainManager
    console.log("\ndeploy EthCrossChainManager ......".cyan);
    const ccmp = await CCMP.connect(creator).deploy(empty.address,admin.address,'0x');
    await ccmp.deployed();
    console.log("EthCrossChainManager deployed to:".green, ccmp.address.blue);
    
    // deploy LockProxy
    console.log("\ndeploy LockProxy ......".cyan);
    const lockProxy = await LockProxy.connect(admin).deploy();
    await lockProxy.deployed();
    console.log("LockProxy deployed to:".green, lockProxy.address.blue);
  
    // deploy EthCrossChainData
    console.log("\ndeploy EthCrossChainData ......".cyan);
    const eccd = await ECCD.connect(admin).deploy({gasLimit: 8000000});
    await eccd.deployed();
    console.log("EthCrossChainData deployed to:".green, eccd.address.blue);
    
    // deploy CallerFactory
    console.log("\ndeploy CallerFactory ......".cyan);
    const cf = await CallerFactory.connect(admin).deploy([lockProxy.address, nativeLockProxy]);
    await cf.deployed();
    console.log("CallerFactory deployed to:".green, cf.address.blue);
    
    // update Const.sol
    console.log("\nupdate Const.sol ......".cyan);
    await updateConst(eccd.address, cf.address);
    console.log("Const.sol updated".green);
    await hre.run('compile');
    
    // deploy EthCrossChainManagerImplementation
    console.log("\ndeploy EthCrossChainManagerImplementation ......".cyan);
    const ccm = await CCM.connect(admin).deploy();
    await ccm.deployed();
    console.log("EthCrossChainManagerImplementation deployed to:".green, ccm.address.blue);
    
    // setup EthCrossChainManager
    console.log("\nsetup EthCrossChainManager ......".cyan);
    tx = await ccmp.connect(admin).upgradeTo(ccm.address);
    await tx.wait();
    console.log("set".green);
  
    // transfer ownership
    console.log("\ntransfer eccd's ownership to ccm ......".cyan);
    tx = await eccd.connect(admin).transferOwnership(ccmp.address);
    await tx.wait();
    console.log("ownership transferred".green);

    // setup contracts
    console.log("\nsetup LockProxy ......".cyan);
    tx = await lockProxy.connect(admin).setManagerProxy(ccmp.address);
    await tx.wait();
    console.log("LockProxy set".green);

    // deploy wrapper
    console.log("\ndeploy PolyWrapperV1 ......".cyan);
    const wrapper1 = await WrapperV1.deploy(admin.address, polyId);
    await wrapper1.deployed();
    console.log("PolyWrapperV1 deployed to:".green, wrapper1.address.blue);
  
    console.log("\ndeploy PolyWrapperV2 ......".cyan);
    const wrapper2 = await WrapperV2.deploy(admin.address, polyId);
    await wrapper2.deployed();
    console.log("PolyWrapperV2 deployed to:".green, wrapper2.address.blue);
    
    // setup wrapper
    console.log("\nsetup WrapperV1 ......".cyan);
    tx = await wrapper1.setFeeCollector(admin.address);
    await tx.wait();
    console.log("setFeeCollector Done".green);
    tx = await wrapper1.setLockProxy(lockProxy.address);
    await tx.wait();
    console.log("setLockProxy Done".green);
  
    console.log("\nsetup WrapperV2 ......".cyan);
    tx = await wrapper2.setFeeCollector(admin.address);
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
    
    // // mainnet
    // case 1: // eth-main
    //   return 2;
    // case 56: // bsc-main
    //   return 6;
    // case 128: // heco-main
    //   return 7;
    // case 137: // polygon-main
    //   return 17;
    // case 66: // ok-main
    //   return 12;
    // case 1718: // plt-main
    //   return 8;

    // // testnet
    // case 3: // eth-test
    //   return 2;
    // case 97: // bsc-test
    //   return 79;
    // case 256: // heco-test
    //   return 7;
    // case 80001: // polygon-test
    //   return 216;
    // case 65: // ok-test
    //   return 200;
    // case 421611: // arbitrum-test
    //   return 205;
    // case 77: // xdai-test
    //   return 206;
    // case 69: // op-test
    //   return 207;
    // case 101: // plt-test
    //   return 208;
    // case 31091: // curve-test
    //   return 210;

    case 108:
      return 77;

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
