const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const crypto = require("crypto");
const Web3 = require("web3");
const { expect } = require("chai");
const bytes = require("@ethersproject/bytes");
hre.web3 = new Web3(hre.network.provider);

describe("CCMtest", async function () {

  let empty32 = '0x0000000000000000000000000000000000000000000000000000000000000000';
  let empty20 = '0x0000000000000000000000000000000000000000';
  let deployer;
  let addr1;
  let addr2;
  let eccd;
  let factory;
  let ccmi;
  let ccmp;
  let ccm;
  let caller;
  let calleri;
  let callerp;

  describe("initialize", function () {
    it("Should do deployment ",async function () {
      [deployer,addr1,addr2] = await hre.ethers.getSigners();

      await hre.run('compile');
      
      // deploy EthCrossChainData
      const ECCD = await hre.ethers.getContractFactory("EthCrossChainData");
      eccd = await ECCD.deploy();
      await eccd.deployed();
      
      // deploy CallerFactory
      const CallerFactory = await hre.ethers.getContractFactory("CallerFactory");
      factory = await CallerFactory.deploy([]);
      await factory.deployed();
      
      // update Const.sol
      await updateConst(eccd.address, factory.address);
      await hre.run('compile');
      
      // deploy EthCrossChainManagerImplementation
      const CCM = await hre.ethers.getContractFactory("EthCrossChainManagerImplementation");
      ccmi = await CCM.deploy();
      await ccmi.deployed();
        
      // deploy EthCrossChainManager
      const CCMP = await hre.ethers.getContractFactory("EthCrossChainManager");
      ccmp = await CCMP.deploy(ccmi.address,deployer.address,'0x');
      await ccmp.deployed();

      ccm = await CCM.attach(ccmp.address);
    });

    it("Should transfer eccd's owner", async function () {
      expect(await eccd.owner()).to.equal(deployer.address);
      await eccd.transferOwnership(ccm.address);
      expect(await eccd.owner()).to.equal(ccm.address);
    });
  });

  describe("ProxyFactory", function () {

    it("Should deploy caller implemetaion & caller", async function () {
      // deploy caller implementation
      const Caller = await hre.ethers.getContractFactory("EthCrossChainCaller");
      const CallerMock = await hre.ethers.getContractFactory("CallerImplementationMock");
      calleri = await CallerMock.deploy();
      await calleri.deployed();
      
      let salt = 77777;
      let preCalcAddress = await factory.getDeploymentAddress(salt, addr1.address);
      expect(await factory.isChild(preCalcAddress)).to.equal(false);
      // deploy caller ,  initialize(address).selector = 0xc4d66de8
      await factory.connect(addr1).deploy(salt, calleri.address, addr1.address, '0xc4d66de8000000000000000000000000'+ccm.address.slice(2));

      expect(await factory.isChild(preCalcAddress)).to.equal(true);
      expect(await factory.isChild(calleri.address)).to.equal(false);

      caller = await CallerMock.attach(preCalcAddress);
      callerp = await Caller.attach(preCalcAddress);

      expect(await caller.connect(addr2).whoAmI()).to.equal(1);
    });

    it("Should deploy caller via signature", async function () {
      const CallerSigMsgGen = await hre.ethers.getContractFactory("CallerSigMsgGen");
      let smg = await CallerSigMsgGen.deploy();
      let salt = 66666;
      let logic = calleri.address;
      let admin = addr2.address;
      let data = '0xc4d66de8000000000000000000000000'+ccm.address.slice(2);
      let signer = addr2;
      let msg = await smg.getSigMsg(salt, logic, admin, data, factory.address);
      let sig = await signer.signMessage(bytes.arrayify(msg));
      let preCalcAddress = await factory.getDeploymentAddress(salt, signer.address);
      expect(await factory.isChild(preCalcAddress)).to.equal(false);
      await factory.connect(addr1).deploySigned(salt, logic, admin, data, sig);
      expect(await factory.isChild(preCalcAddress)).to.equal(true);
    });

    it("Should not get correct caller while deploy via fake signature", async function () {
      const CallerSigMsgGen = await hre.ethers.getContractFactory("CallerSigMsgGen");
      let smg = await CallerSigMsgGen.deploy();
      let salt_fake = 1234;
      let salt = 1111;
      let logic = calleri.address;
      let admin = addr2.address;
      let data = '0xc4d66de8000000000000000000000000'+ccm.address.slice(2);
      let signer = addr2;
      let msg = await smg.getSigMsg(salt_fake, logic, admin, data, factory.address);
      let sig = await signer.signMessage(msg);
      let preCalcAddress = await factory.getDeploymentAddress(salt, signer.address);
      expect(await factory.isChild(preCalcAddress)).to.equal(false);
      await factory.connect(addr1).deploySigned(salt, logic, admin, data, sig);
      expect(await factory.isChild(preCalcAddress)).to.equal(false);
    });

    it("Should fail if not admin try to set implementation", async function () {
      await expect(callerp.connect(addr2).upgradeTo(ccmi.address)).to.be.reverted;
    });

    it("Should success if admin try to set implementation", async function () {
      const CallerMock2 = await hre.ethers.getContractFactory("CallerImplementationMock_2");
      calleri2 = await CallerMock2.deploy();
      await calleri2.deployed();
      await callerp.connect(addr1).upgradeTo(calleri2.address);
      expect(await caller.connect(addr2).whoAmI()).to.equal(2);
      await callerp.connect(addr1).upgradeTo(calleri.address);
      expect(await caller.connect(addr2).whoAmI()).to.equal(1);
    });

    it("Should fail if not admin try to change admin", async function () {
      await expect(callerp.connect(addr2).changeAdmin(addr2.address)).to.be.reverted;
    });

    it("Should success if admin try to change admin", async function () {
      await callerp.connect(addr1).changeAdmin(addr2.address);
      await expect(callerp.connect(addr1).changeAdmin(addr1.address)).to.be.reverted;
      await callerp.connect(addr2).changeAdmin(addr1.address);
    });
  });

  describe("EthCrossChainManager", function () {
    
    // upgradeTo
    it("Should fail if not admin try to set implementation", async function () {
      const CCMMock = await hre.ethers.getContractFactory("CCMMock");
      let mock = await CCMMock.deploy();
      await expect(ccmp.connect(addr2).upgradeTo(mock.address)).to.be.reverted;
    });

    it("Should success if admin try to set implementation", async function () {
      const CCMMock = await hre.ethers.getContractFactory("CCMMock");
      let mock = await CCMMock.deploy();
      let mockccm = await CCMMock.attach(ccm.address);
      await ccmp.connect(deployer).upgradeTo(mock.address);
      expect(await mockccm.connect(addr2).iAmMock()).to.equal(true);
      await ccmp.connect(deployer).upgradeTo(ccmi.address);
      await expect( mockccm.connect(addr2).iAmMock()).to.be.reverted;
    });
    
    // changeAdmin
    it("Should fail if not admin try to change admin", async function () {
      await expect(ccmp.connect(addr2).changeAdmin(addr2.address)).to.be.reverted;
    });

    it("Should success if admin try to change admin", async function () {
      await ccmp.connect(deployer).changeAdmin(addr2.address);
      await expect(ccmp.connect(deployer).changeAdmin(deployer.address)).to.be.reverted;
      await ccmp.connect(addr2).changeAdmin(deployer.address);
    });

    // crossChain
    it("Should success when try to call crossChain from valid caller (child of factory) ", async function () {
      let args = "0x123456"
      let index = await eccd.getEthTxHashIndex();
      expect(await eccd.getEthTxHash(index)).to.equal(empty32);
      await expect(caller.lock(args)).to.not.reverted;
      expect(await eccd.getEthTxHash(index)).to.not.equal(empty32);
      expect(await eccd.getEthTxHashIndex()).to.equal(index+1);
    });
    
    it("Should fail when try to call crossChain from invalid caller (not child of factory) ", async function () {
      let args = "0x123456"
      await calleri.initialize(ccm.address);
      await expect(calleri.lock(args)).to.be.reverted;
    });
    
    // TODO
    // initGenesisBlock
    it("Should success when try to call initGenesisBlock at uninitialized CCM ", async function () {
      // header at height 999
      let rawHeader = "0xf90288a07d12f7d0baf7a2703dbd67b54b68ce7a095dfe6c2db0fd9f332bcef02b4373e4a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d4934794c095448424a5ecd5ca7ccdadfaad127a9d7e88eca05eeb9c115c1235fb4f19dbf988f2922b312acbc8483019a1943b3337594e82d1a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018203e784607570868084617777c7b8900000000000000000000000000000000000000000000000000000000000000000f86ef86994258af48e28e4a6846e931ddff8e1cdf8579821e5946a708455c8777630aac9d1e7702d13f7a865b27c948c09d936a1b408d6e0afaa537ba4e06c4504a0ae94ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c94c095448424a5ecd5ca7ccdadfaad127a9d7e88ec80c080a063746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365880000000000000000"
    
      expect(await eccd.getCurEpochStartHeight()).to.equal(0);
      await ccm.connect(addr1).initGenesisBlock(rawHeader)
      expect(await eccd.getCurEpochStartHeight()).to.equal(1000);
      let vals = await eccd.getCurEpochValidatorPkBytes();

    });
    
    // TODO
    it("Should fail when try to call initGenesisBlock at already initialized CCM ", async function () {
      let rawHeader = "0xf90288a07d12f7d0baf7a2703dbd67b54b68ce7a095dfe6c2db0fd9f332bcef02b4373e4a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d4934794c095448424a5ecd5ca7ccdadfaad127a9d7e88eca05eeb9c115c1235fb4f19dbf988f2922b312acbc8483019a1943b3337594e82d1a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018203e784607570868084617777c7b8900000000000000000000000000000000000000000000000000000000000000000f86ef86994258af48e28e4a6846e931ddff8e1cdf8579821e5946a708455c8777630aac9d1e7702d13f7a865b27c948c09d936a1b408d6e0afaa537ba4e06c4504a0ae94ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c94c095448424a5ecd5ca7ccdadfaad127a9d7e88ec80c080a063746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365880000000000000000"
      
      expect(await eccd.getCurEpochStartHeight()).to.equal(1000);
      await expect(ccm.connect(addr1).initGenesisBlock(rawHeader)).to.be.reverted;
    });

    // TODO
    // changeEpoch
    it("Should fail when try to call changeEpoch with incorrect header ", async function () {
      // invalid header 
      let _invalidHeader = "0x"
      let rawSeals = "0xf9010cb841b0e9aaec13e2fa8e7f3a24194f1efd533478d1a16673d64409340878bc882bcb5d89111cf0669085d0938ffe06ae67a45a2c1e5a6d7d3b93ea5551433812641a01b8419dc0b38ab186feec929cd1bcf8215353edaa73294569b3f718d1ba849e4b47252f69d48bfe1eefb0690cad8289747d4556bba3e952614836d5f8084527290e7600b8419a9f172e46b1234465b96bdfb0a2cf971662ad3f94911ba9082013f78262e6fc04c08c8d3b99f42bdc5d5e0fb7c8b67e5807bf168e64bcbb240b8f1584a74fd601b841867de707c4cc0dbd733a96259652480279c016019070a53887fe004b9c2b0d923f1d5ecbcb556a03d7a153cfac95448955af80b27321d7869f9ad18fd2b7747e00"

      expect(await eccd.getCurEpochStartHeight()).to.equal(1000);
      await expect(ccm.connect(addr1).changeEpoch(_invalidHeader, rawSeals)).to.be.reverted;
    });
    
    // TODO
    it("Should fail when try to call changeEpoch with incorrect seals ", async function () {
      let rawHeader = "0xf902b2a08adbb7aa118074c58ce20966b19734a4a0cfa2898f5bfcd01b086b068351ff5aa01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347948c09d936a1b408d6e0afaa537ba4e06c4504a0aea08a5364fbb3e7d3c5076c9c887d84e9569d277975ab9b428046a311a60543639ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000182044b84577ae92780846177784fb8ba0000000000000000000000000000000000000000000000000000000000000000f898f89394258af48e28e4a6846e931ddff8e1cdf8579821e5946a708455c8777630aac9d1e7702d13f7a865b27c948c09d936a1b408d6e0afaa537ba4e06c4504a0ae94ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c94c095448424a5ecd5ca7ccdadfaad127a9d7e88ec94d47a4e56e9262543db39d9203cf1a2e53735f83494bfb558f0dceb07fbb09e1c283048b551a431092180c080a063746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365880000000000000000"
      let _invalidSeals = "0xf9014fb841017109bf1d901853c28308d7565aeeef143fae070d36700257af30cd4730528c1e4233554f6c35d008069b21a98a3eb58a2c375c2c0f03c58e2e4878d270726801b8418378937fb262559217c22a9b58b292eb913f017c29a4ff2201e69ba5c4d53ccc40bba4705f427521146e51dbc2d60cdc7aba129c15b79e94b9d11fd41874413100b84188bf7375d0b884e65de33c6b1d6e87fbd1b5010df5dbd93a47a3d1fa435e41df10aac4fbf091b798509252c9058011852a8ef78f0bed2879b9bd4ba87b4c509800b841a0811f10073b27a0683532ba65e4bcbdc60be479de9b542951ca36f3615dc3cd05dc60413ad27dcf2eee1ad83932a82b2b0a97a971c9c15e31455fb402fc597101b841d6d838080755e680a4e35e0a389f7d23703604ed528a572b297c0d04d169562b00f828f0a0589822e2a2e3cbf2c329931da4b14fce03ac4007087f221e37689400"

      expect(await eccd.getCurEpochStartHeight()).to.equal(1000);
      await expect(ccm.connect(addr1).changeEpoch(rawHeader, _invalidSeals)).to.be.reverted;
    });
    
    // TODO
    it("Should fail when try to call changeEpoch with block which is not the last block of current epoch ", async function () {
      // header of 1098 block (should be 1099)
      let rawHeader = "0xf9021ca03829bd8b8ab1b8a1c9a3c70b0f949fa5e9492f764eea365e42d727be7459bc58a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347946a708455c8777630aac9d1e7702d13f7a865b27ca08a5364fbb3e7d3c5076c9c887d84e9569d277975ab9b428046a311a60543639ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000182044a845790cd5980846177784ea50000000000000000000000000000000000000000000000000000000000000000c4c080c080a063746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365880000000000000000"
      let rawSeals = "0xf9010cb841f2bd0fb5fa83ca6a5dce8dffd67a63dcdff7266b1022652a996e4642eb5a6f2d0fd2ff55a4eb979d5b081751209af8d52c345cb01137382f2118792f650876c300b841d009bf5d3724e52e162a74e6c332a42c4c96ef63f056d219acc48102d3b4d3fc57a54167988f7f4c738b7004c427ee20cc32b5d89a8643bdced0c43fdb358daa01b841e038e8b07e73ba38d801c0158e46c72ae0997ca51d9739834f6b2c0f039e25bd7c859b17c0e46d828da2e62a87213620c07b1f977505477cc57aaa2672c8b4dc01b84132fe2a78b1e5a41f844a0c57819203ed105d5558a9ddbefa80c161a651c0607c685235096689379522a67fdb3edea5a818b33e490d1300b51493dc68dbad66c201"

      expect(await eccd.getCurEpochStartHeight()).to.equal(1000);
      await expect(ccm.connect(addr1).changeEpoch(rawHeader, rawSeals)).to.be.reverted;
    });

    // TODO
    it("Should success when try to call changeEpoch with correct header & seals ", async function () {
      let rawHeader = "0xf902b2a08adbb7aa118074c58ce20966b19734a4a0cfa2898f5bfcd01b086b068351ff5aa01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347948c09d936a1b408d6e0afaa537ba4e06c4504a0aea08a5364fbb3e7d3c5076c9c887d84e9569d277975ab9b428046a311a60543639ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000182044b84577ae92780846177784fb8ba0000000000000000000000000000000000000000000000000000000000000000f898f89394258af48e28e4a6846e931ddff8e1cdf8579821e5946a708455c8777630aac9d1e7702d13f7a865b27c948c09d936a1b408d6e0afaa537ba4e06c4504a0ae94ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c94c095448424a5ecd5ca7ccdadfaad127a9d7e88ec94d47a4e56e9262543db39d9203cf1a2e53735f83494bfb558f0dceb07fbb09e1c283048b551a431092180c080a063746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365880000000000000000"
      let rawSeals = "0xf9010cb841e065730b6ab96d2172eb9c34c7b89dcec57a4933ffdc0e941ac0beb91a69861604007302525252a94919e0f302830cd444ab43b52f593a8a287afc82aa5ef08f00b841e3696e39bb78b65fb9a9af5960e95147a8a2b54fdb78876a1ac5adeb2d09497c33bbf05bd0cdcc921e16ab66a3b409268b551ee9f805db587e19557d3e5bad7701b8417ed6a92254435c47ce41c8119621a1b9a2e1eda12241912c8127b5fd3ee76d9d5b9326dc50ae6e4f9d142086a7eb4b99f5bdd270d86d22adbe4591c33c6f5c4700b8418aece0db7a7534f6d8f3b2b49c5bd5aae27c20fe8637f57670c24c8255a7119d1d70eb279d1dceb397514f6c7182981aad0b20febd18d00653056fb837d205cc01"

      expect(await eccd.getCurEpochStartHeight()).to.equal(1000);
      await ccm.connect(addr1).changeEpoch(rawHeader, rawSeals);
      expect(await eccd.getCurEpochStartHeight()).to.equal(1100);
    });
    
    

    // TODO
    // verifyHeaderAndExecuteTx
    it("Should success when try to call verifyHeaderAndExecuteTx with correct header & seals & proof & crossTx ", async function () {
    });
    
    // TODO
    it("Should fail when try to call verifyHeaderAndExecuteTx with incorrect header ", async function () {
    });
    
    // TODO
    it("Should fail when try to call verifyHeaderAndExecuteTx with incorrect seals ", async function () {
    });
    
    // TODO
    it("Should fail when try to call verifyHeaderAndExecuteTx with incorrect proof ", async function () {
    });
    
    // TODO
    it("Should fail when try to call verifyHeaderAndExecuteTx with incorrect crossTx ", async function () {
    });


  });
});


async function updateConst(eccd, callerFactory) {
  // const polyChainId = await getPolyChainId();
  polyChainId = 2;

  await fs.writeFile('./contracts/core/cross_chain_manager/logic/Const.sol', 
  'pragma solidity ^0.5.0;\n'+
  'contract Const {\n'+
  '    bytes constant ZionCrossChainManagerAddress = hex"0000000000000000000000000000000000001003"; \n'+
  '    address constant EthCrossChainDataAddress = '+eccd+'; \n'+
  '    address constant EthCrossChainCallerFactoryAddress = '+callerFactory+'; \n'+
  '    uint constant chainId = '+polyChainId+'; \n}', 
  function(err) {
    if (err) {
        return console.error(err);
    }
  }); 
}

