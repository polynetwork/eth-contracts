// require("dotenv").config();

require("@nomiclabs/hardhat-etherscan");
require("@nomiclabs/hardhat-waffle");
//require("hardhat-gas-reporter");
//require("solidity-coverage");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

let PRIVATE_KEY = '';
let PRIVATE_KEY_1 = '';
let ETHERSCAN_API_KEY = '';

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: {
    version: "0.5.17",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  networks: {
    hardhat: {
      initialBaseFeePerGas: 0, // workaround from https://github.com/sc-forks/solidity-coverage/issues/652#issuecomment-896330136 . Remove when that issue is closed.
    },
    mainnet: {
      url: `https://mainnet.infura.io/v3/f7fd127c230e49b292bafd6c338995ea`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    arbitrum: {
      url: `https://arb1.arbitrum.io/rpc`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    bsc: {
      url: `http://172.168.3.40:8545`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    heco: {
      url: `https://http-mainnet.hecochain.com`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    ok: {
      url: `https://exchainrpc.okex.org/`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    polygon: {
      url: `https://rpc-mainnet.matic.network`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },


    ropsten: {
      url: `https://ropsten.infura.io/v3/9bca539684b6408d9dbcbb179e593eab`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    arbitrum_testnet: {
      url: `https://rinkeby.arbitrum.io/rpc`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    bsc_testnet: {
      url: `https://data-seed-prebsc-1-s1.binance.org:8545`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    heco_testnet: {
      url: `https://http-testnet.hecochain.com`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    ok_testnet: {
      url: `https://exchaintestrpc.okex.org/`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    polygon_testnet: {
      url: `https://rpc-mumbai.maticvigil.com`,
      accounts: [`0x${PRIVATE_KEY}`,`0x${PRIVATE_KEY_1}`]
    },
    // ropsten: {
    //   url: process.env.ROPSTEN_URL || "",
    //   accounts:
    //     process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [],
    // },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: ETHERSCAN_API_KEY,
  },
};