const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const crypto = require("crypto");
const Web3 = require("web3");
const { expect } = require("chai");
const bytes = require("@ethersproject/bytes");
hre.web3 = new Web3(hre.network.provider);

describe("LockProxyTest", async function () {

  let empty32 = '0x0000000000000000000000000000000000000000000000000000000000000000';
  let empty20 = '0x0000000000000000000000000000000000000000';
  let chainId = 7777;
  let deployer;
  let addr1;
  let addr2;
  let addrs;
  let lockProxy;
  let ccm;
  let ccmpMock;
  let tunnelCaller;

  describe("initialize", function () {

    it("Should do deployment ",async function () {
        [deployer, addr1, addr2, ...addrs] = await ethers.getSigners();

        const CallerFactory = await hre.ethers.getContractFactory("CallerFactory");
        let factory = await CallerFactory.deploy();
        await factory.deployed();

        CCMMock = await ethers.getContractFactory("CCMMock");
        let ccmi = await CCMMock.deploy();
        await ccmi.deployed();

        const CCMP = await hre.ethers.getContractFactory("EthCrossChainManager");
        let ccmp = await CCMP.deploy(ccmi.address,deployer.address,'0x');
        await ccmp.deployed();
  
        ccm = await CCMMock.attach(ccmp.address);
        
        const LockProxy = await ethers.getContractFactory("LockProxy");
        lockProxy = await LockProxy.deploy();
        await lockProxy.deployed();

        const TunnelCCMCaller = await ethers.getContractFactory("TunnelCCMCaller");
        tunnelCalleri = await TunnelCCMCaller.deploy();
        await tunnelCalleri.deployed();
        
        let salt = 777;
        let tunnelCallerAddress = await factory.getDeploymentAddress(salt, addr1.address);
        await factory.connect(addr1).deploy(salt, tunnelCalleri.address, addr1.address, '0xc4d66de8000000000000000000000000'+deployer.address.slice(2));
  
        tunnelCaller = await TunnelCCMCaller.attach(tunnelCallerAddress);

    });

    it("Should setup contracts", async function () {
        await tunnelCaller.setRealCCM(ccm.address);
        await tunnelCaller.bindCallerHash(chainId, tunnelCaller.address);

        await lockProxy.setManagerProxy(tunnelCaller.address);
        await lockProxy.bindProxyHash(chainId, lockProxy.address);
        await lockProxy.bindAssetHash(empty20, chainId, empty20);
    });
  });

  describe("crossChain", function () {
      it("Should lock & unlock correctly", async function () {
          let addr1_balance_before = await getBalance(addr1.address);
          let addr2_balance_before = await getBalance(addr2.address);
          let lockProxy_balance_before = await getBalance(lockProxy.address);
          expect(await ccm.connect(addr1)._toChainId()).to.equal(0);
          expect(await ccm.connect(addr1)._toContract()).to.equal('0x');
          expect(await ccm.connect(addr1)._method()).to.equal('0x');
          expect(await ccm.connect(addr1)._txData()).to.equal('0x');
          expect(await ccm.connect(addr1)._fromContract()).to.equal('0x');
          expect(await ccm.connect(addr1)._fromChainId()).to.equal(0);

          let amount = 77777777;
          // crossChain
          let res = await lockProxy.connect(addr1).lock(empty20, chainId, addr2.address, amount, {value: amount});
          let recepit = await hre.web3.eth.getTransactionReceipt(res.hash);
          let gas = recepit.gasUsed * 10**9;
          expect(await getBalance(addr1.address)).to.equal(addr1_balance_before - amount - gas);
          expect(await getBalance(addr2.address)).to.equal(addr2_balance_before);
          expect(await getBalance(lockProxy.address)).to.equal(lockProxy_balance_before + amount);
          
          
          let toChainId = await ccm.connect(addr1)._toChainId();
          let toContract = await ccm.connect(addr1)._toContract();
          let method = await ccm.connect(addr1)._method();
          let args = await ccm.connect(addr1)._txData();
          let fromContract = await ccm.connect(addr1)._fromContract();
          let fromChainId = await ccm.connect(addr1)._fromChainId();
          expect(toChainId).to.equal(chainId);
          expect(toContract).to.equal(tunnelCaller.address.toLowerCase());
          expect(method).to.equal('0x756e77726170'); // unwrap : 756e77726170
          expect(fromContract).to.equal(tunnelCaller.address.toLowerCase());
          expect(fromChainId).to.equal(chainId);
          
          // verifyHeaderAndExecuteTx
          await ccm.connect(addr1).verifyHeaderAndExecuteTx(toContract, method, args, fromContract, fromChainId);
          expect(await getBalance(addr2.address)).to.equal(addr2_balance_before + amount);
          expect(await getBalance(lockProxy.address)).to.equal(lockProxy_balance_before);

      });
  });


});

async function getBalance(addr) {
    let b = await hre.web3.eth.getBalance(addr);
    return b-1+1;
}
