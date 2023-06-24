/**
* @type import('hardhat/config').HardhatUserConfig
*/

require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
require("dotenv").config();

module.exports = {
  solidity: {
    version: "0.8.1",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    hardhat: {},
    polygon: {
      url: process.env.POLYGON_RPC,
      accounts: [process.env.MNEMONIC],
    },
    mumbai: {
      url: process.env.MUMBAI_RPC,
      accounts: [process.env.MNEMONIC],
    },
  },
};