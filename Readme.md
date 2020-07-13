* EthCrossChain.sol是ETH上管理合约，可处理btc<->eth, ont<->ont的跨链交易，需要布署。

* MockOntToken.sol是ERC20的币，用来模仿ETH上的OEP-4代币，需要布署。

* MockBtcToken.sol是ERC20的币，用来模仿ETH上的BTC，需要布署。

* Utils.sol包含一些常用的转换函数，不需要单独布署。

* ZeroCopySink.sol, ZeroCopySource.sol 不需要单独布署。

* IERC20, ERC20Token是用来被继承的合约，不需要单独布署。







```
./abigen --sol ./contracts/core/v2.0/CrossChainManager/EthCrossChainData.sol --pkg eccd_abi > ./go_abi/eccd_abi.go

./abigen --sol ./contracts/core/v2.0/CrossChainManager/EthCrossChainManager.sol --pkg eccm_abi > ./go_abi/eccm_abi.go

./abigen --sol ./contracts/core/v2.0/CrossChainManager/EthCrossChainManagerProxy.sol --pkg eccmp_abi > ./go_abi/eccmp_abi.go

./abigen --sol ./contracts/core/v2.0/lockproxy/LockProxy.sol --pkg lock_proxy_abi > ./go_abi/lock_proxy.go

./abigen --sol ./contracts/core/v2.0/assets/btc/BTCX.sol --pkg btcx_abi > ./go_abi/btcx_abi.go


./abigen --sol ./contracts/core/v2.0/assets/ont/ONTX.sol --pkg ontx_abi > ./go_abi/ontx_abi.go

./abigen --sol ./contracts/core/v2.0/assets/ong/ONGX.sol --pkg ongx_abi > ./go_abi/ongx_abi.go

./abigen --sol ./contracts/core/v2.0/assets/oep4template/OEP4Template.sol --pkg oep4_abi > ./go_abi/oep4_abi.go

./abigen --sol ./contracts/core/v2.0/assets/erc20template/ERC20Template.sol --pkg erc20_abi > ./go_abi/erc20_abi.go

```


redeemScript: 552102dec9a415b6384ec0a9331d0cdf02020f0f1e5731c327b86e2b5a92455a289748210365b1066bcfa21987c3e207b92e309b95ca6bee5f1133cf04d6ed4ed265eafdbc21031104e387cd1a103c27fdc8a52d5c68dec25ddfb2f574fbdca405edfd8c5187de21031fdb4b44a9f20883aff505009ebc18702774c105cb04b1eecebcb294d404b1cb210387cda955196cc2b2fc0adbbbac1776f8de77b563c6d2a06a77d96457dc3d0d1f2102dd7767b6a7cc83693343ba721e0f5f4c7b4b8d85eeb7aec20d227625ec0f59d321034ad129efdab75061e8d4def08f5911495af2dae6d3e9a4b6e7aeb5186fa432fc57ae



## 测试网生产环境：


#### 管理合约

- Eth链Hash: <code>0x48a77f43c0d7a6d6f588c4758dba22bf6c5d95a0</code>

- ONT链Hash(反序后的): <code>TODO</code>

#### 代理合约：

- Eth链Hash: <code>TODO</code>

- ONT链Hash(反序后的): <code>TODO</code>

#### 资产合约：

- BTC资产：

    BTC链hash: <code>c330431496364497d7257839737b5e4596f5ac06</code>

    ETH链Hash: <code>0x80Dd1fC1E405d3BC8d9b713Aa8a5dF25B0a55663</code>

    ONT链Hash(反序后的): <code>TODO</code>


- ETH资产：

    ETH链Hash: <code>0x0000000000000000000000000000000000000000</code>

    ONT链Hash(反序后的): <code>TODO</code>


- ONT资产：

    ETH链Hash: <code>TODO</code>

    ONT链Hash(反序后的): <code>0000000000000000000000000000000000000001</code>


- ONG资产：

    ETH链Hash: <code>TODO</code>

    ONT链Hash(反序后的): <code>0000000000000000000000000000000000000002</code>


- OEP4资产：

    ETH链Hash: <code>TODO</code>

    ONT链Hash(反序后的): <code>TODO</code>


- ERC20资产：

    ETH链Hash: <code>TODO</code>

    ONT链Hash(反序后的): <code>TODO</code>




## 开发环境：
ETHURL: http://172.168.3.77:8545
#### ETH管理合约的代理合约：eccmp

<code>0x87fBd8539bd91C980d4DD2dcc10A40167cDdebd3</code>

#### ETH管理合约的数据合约：eccd

<code> 0xA2aA6B82504B4db1fA12f90681016acEb0baB7F0</code>

#### 管理合约  eccm

- Eth链Hash:  old <code>0x32400183a587117DaAB27869C7a9ebedEcb0Df3f</code>, new <code>0xf2536d132592EBC27D51E9E42CD18DaB8f9fde80</code>

- ONT链Hash(反序后的): <code>0000000000000000000000000000000000000009</code>

#### 代理合约：

- Eth链Hash: <code></code>

- ONT链Hash(反序后的): <code></code>

#### 资产合约：

- BTC资产:

    BTC链Hash: <code>c330431496364497d7257839737b5e4596f5ac06</code>

    ETH链Hash: <code>0x8539A4d5CBa55466788D2eaD4B5308b40d47572D</code>

    ONT链Hash(反序后的): <code></code>

- ETH资产：

ETH链Hash: <code>0x0000000000000000000000000000000000000000</code>

ONT链Hash(反序后的): <code></code>


- ONT资产：

    ETH链Hash: <code></code>

    ONT链Hash(反序后的): <code>0000000000000000000000000000000000000001</code>


- ONG资产：

    ETH链Hash: <code></code>

    ONT链Hash(反序后的): <code>0000000000000000000000000000000000000002</code>


- OEP4资产：

    ETH链Hash: <code></code>

    ONT链Hash(反序后的): <code>371c5a021351742d3877c9b77d856a8e1f86fb48</code>


- ERC20资产：

    ETH链Hash: <code></code>

    ONT链Hash(反序后的): <code>40adbc8be26f87e94401d165ffe190cd45108a7a</code>

## truffle test instruction

#### npm install 
```
npm install --save-dev @openzeppelin/test-helpers

```



[truffle 文档](https://learnblockchain.cn/docs/truffle/getting-started/interacting-with-your-contracts.html)

[web3.js api reference](https://web3.tryblockchain.org/Web3.js-api-refrence.html)


<code> truffle compile</code>

<code> truffle migrate</code>

<code> truffle test .\test\mytoken1.js</code>

The difference between <code>mytoken1.js</code> and <code>mytoken2.js</code> is almost nothing.

<code>mytoken1.js</code> is the official recommended style.

<code>mytoken2.js</code> is a better way to adjust if different test cases are independent.


## In truffle console / truffle development mode

#### Get contract instance in truffle console:
1. <code>let instance = await MetaCoin.deployed()</code>
2. Or suppose contract address ： <code>0x133739AB03b9be2b885cC11e3B9292FDFf45440E</code>, 
    we can use <code>let instance = await MyToken.at("0x133739AB03b9be2b885cC11e3B9292FDFf45440E")</code> to obtain the instance.

#### Call contract function (pre-invoke)
```
instance.name() / await instance.name()

isntance.symobl()

(await instance.totalSupply()).toNumber()

```

#### send Transaction: (invoke)
<code>let result = await instance.transfer(accounts[1], 100)</code>

the returned structure is the following:
```
result = {
    "tx": .., 
    "receipt": ..,
    "logs": []
}

```

<code> await instance.approve(accounts[2], 100)</code>

<code> (await instance.allowance(accounts[0], accounts[2])).toNumber()</code>

<code> let result1 = await instance.transferFrom(accounts[0], accounts[2], 1, {from:accounts[2]})
</code> ({from: ***} to indicate the msg.sender of this transaction)

#### about ether

Note: in truffle console, <code>accounts</code> by default means <code>eth.web3.accounts</code>

query balance : <code>web3.eth.getBalance(accounts[0])</code>

send ether: <code>web3.eth.sendTransaction({from: accounts[0], to: accounts[1], value: web3.utils.toWei('1', "ether")})</code>






BTC Info

Redeem：
552102dec9a415b6384ec0a9331d0cdf02020f0f1e5731c327b86e2b5a92455a289748210365b1066bcfa21987c3e207b92e309b95ca6bee5f1133cf04d6ed4ed265eafdbc21031104e387cd1a103c27fdc8a52d5c68dec25ddfb2f574fbdca405edfd8c5187de21031fdb4b44a9f20883aff505009ebc18702774c105cb04b1eecebcb294d404b1cb210387cda955196cc2b2fc0adbbbac1776f8de77b563c6d2a06a77d96457dc3d0d1f2102dd7767b6a7cc83693343ba721e0f5f4c7b4b8d85eeb7aec20d227625ec0f59d321034ad129efdab75061e8d4def08f5911495af2dae6d3e9a4b6e7aeb5186fa432fc57ae

redeem-hash160(key):
c330431496364497d7257839737b5e4596f5ac06
