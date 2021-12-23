const hre = require("hardhat");
const Web3 = require("web3");
const fs = require("fs");
hre.web3 = new Web3(hre.network.provider);
require("colors");

async function main() {
    [deployer, deployer2] = await hre.ethers.getSigners();
    var polyID;
    var ChainName;

    // check if given networkId is registered
    await getPolyChainId().then((polyIDandName) => {
      console.log("\nDeploy contracts on chain with Poly_Chain_Id:".cyan, polyIDandName);
      polyID= polyIDandName[0]
      ChainName=polyIDandName[1]
    }).catch((error) => {
      throw error;
    });
    NetProvider=hre.config.networks[ChainName].url
    console.log(NetProvider)

    //const polyId = await getPolyChainId();
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
    const wrapper1 = await WrapperV1.deploy(deployer.address, polyID);
    await wrapper1.deployed();
    console.log("PolyWrapperV1 deployed to:".green, wrapper1.address.blue);
  
    console.log("\ndeploy PolyWrapperV2 ......".cyan);
    const wrapper2 = await WrapperV2.deploy(deployer.address, polyID);
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

   //output json 
    var config = {
      Name: ChainName,
      PolyChainID : polyID,
      Provider: NetProvider, 
      Deployer:deployer.address,  
      EthCrossChainData : eccd.address,      
      EthCrossChainManagerImplemetation : ccm.address,
      EthCrossChainManager: ccmp.address,
      CallerFactory:cf.address,
      LockProxy: lockProxy.address,
      Swapper:"0x0000000000000000000000000000000000000000",
      Wrapper:wrapper1.address,
      WrapperV2:wrapper2.address,
    };
    console.log("\nconstract output\n",config);
  //read previous config
    let data=fs.readFileSync("./zionDevConfig.json",(err,data)=>{
        if (err) {
          throw err;
        }else{
          previous=data.toString();
        }  
    });
   //add new config
   //var buffer=JSON.stringify(data)
    var json=JSON.parse(data.toString())
    json.Network[json.Network.length]=config
    var jsonConfig =JSON.stringify(json,null,"\t")
    var outputPath = './zionDevConfig.json';
    //console.log("\njson output\n",jsonConfig);
    try {
      fs.writeFileSync(outputPath, jsonConfig);
    } catch (err) {
      console.error(err);
    }
}

async function updateConst(eccd, callerFactory) {
    const polyChainId = await getPolyChainId();
  
    fs.writeFileSync('./contracts/core/cross_chain_manager/logic/Const.sol', 
    'pragma solidity ^0.5.0;\n'+
    'contract Const {\n'+
    '    bytes constant ZionCrossChainManagerAddress = hex"5747C05FF236F8d18BB21Bc02ecc389deF853cae"; \n'+
    '    bytes constant ZionValidaterManagerAddress = hex"A4Bf827047a08510722B2d62e668a72FCCFa232C"; \n'+
    '    \n'+
    '    address constant EthCrossChainDataAddress = '+eccd+'; \n'+
    '    address constant EthCrossChainCallerFactoryAddress = '+callerFactory+'; \n'+
    '    uint constant chainId = '+polyChainId[0]+'; \n}', 
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
      return [2,"mainnet"];
    case 56: // bsc-main
      return [6,"bsc"];
    case 128: // heco-main
      return [7,"heco"];
    case 137: // polygon-main
      return [17,"polygon"];
    case 66: // ok-main
      return [12,"ok"];
    case 1718: // plt-main
      return [8,"plt"];

    // testnet
    case 3: // ropsten
      return [202,"ropsten"];
    case 4: // rinkeby
      return [402,"rinkeby"];
    case 5: // goerli
      return [502,"goerli"];
    case 42: // kovan
      return [302,"kovan"];
    case 97: // bsc-test
      return [1000006,"bsc_testnet"];
    case 256: // heco-test
      return [7,"heco_testnet"];
    case 80001: // polygon-test
      return [216,"polygon_testnet"];
    case 65: // ok-test
      return [200,"ok_testnet"];
    case 421611: // arbitrum-test
      return [300,"arbitrum_testnet"];
    case 4002: // ftm-test
      return [400,"fantom_testnet"];
    case 43113: // avax-test
      return [500,"avax_testnet"];
    case 77: // xdai-test
      return [600,"xdai_testnet"];
    case 69: // op-test
      return [200,"op_testnet"];
    case 101: // plt-test
      return [208,"plt_testnet"];
    case 31091: // curve-test
      return [210,"curve_testnet"];
    case 10897: // zion_side_dev
      return [77,"zion_side_dev"];

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
