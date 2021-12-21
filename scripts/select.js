const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const Web3 = require("web3");
const colors = require('colors');
hre.web3 = new Web3(hre.network.provider);

async function main() {

    const A = await hre.ethers.getContractFactory("ProxyAdmin");
    const B = await hre.ethers.getContractFactory("UpgradeabilityProxy");
    
    console.log("deploy a...".yellow)
    let a = await A.deploy();
    console.log("ProxyAdmin at: ".cyan,a.address.blue);

    console.log("deploy b...".yellow)
    let b = await B.deploy(a.address, '0x');
    console.log("UpgradeabilityProxy at: ".cyan,b.address.blue);
    
    console.log("verify a...".yellow);
    await verifyAtEtherscan(a.address, []);
    
    console.log("verify b...".yellow);
    await verifyAtEtherscan(b.address, [a.address,'0x']);


    // let proxyaddr = '0x8a19286584940000072Df9957F73b13845906068';
    
    // await getProxyAdmin(proxyaddr)
    // .then((res) => {
    //     console.log('Admin: ',res)
    // })
    // .catch((err) => {
    //     throw err;
    // });

    // await getProxyImplementation(proxyaddr)
    // .then((res) => { 
    //     console.log('Implementation: ',res)
    // })
    // .catch((err) => {
    //     throw err;
    // });
}

function verifyAtEtherscan(address, constructorArguments) {
    return hre.run("verify:verify", {
      address: address,
      constructorArguments: constructorArguments,
    });
}

function getProxyAdmin(contractAddr) {
    const ADMIN_SLOT = '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103';
    return hre.web3.eth.getStorageAt(contractAddr, ADMIN_SLOT);
}

function getProxyImplementation(contractAddr) {
    const IMPLEMENTATION_SLOT = '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc';
    return hre.web3.eth.getStorageAt(contractAddr, IMPLEMENTATION_SLOT);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });



