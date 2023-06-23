const { ethers } = require("hardhat");
require("dotenv").config();

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log(`Deploying contracts with the account: ${deployer.address}`);

  const ERC721Proxy = await ethers.getContractFactory("ERC721Proxy");
  const smsuances = await ERC721Proxy.deploy(process.env.URI);

  console.log(`7msuances contract deployed at address: ${smsuances.address} `);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });