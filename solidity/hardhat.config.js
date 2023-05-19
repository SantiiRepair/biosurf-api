/**
* @type import('hardhat/config').HardhatUserConfig
*/

require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
require("dotenv").config();

module.exports = {
 solidity: {
   version: "0.8.18",
   settings: {
     optimizer: {
       enabled: true,
       runs: 200,
     },
   },
 },
 networks: {
   hardhat: {},
   ethereum: {
     url: process.env.MAINNET_RPC,
     accounts: [process.env.MNEMONIC],
   },
   sepolia: {
     url: process.env.SEPOLIA_RPC,
     accounts: [process.env.MNEMONIC],
   },
 },
};