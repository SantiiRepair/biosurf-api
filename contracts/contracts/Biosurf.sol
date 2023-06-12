// SPDX-License-Identifier: MIT
 
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

interface IOwnable {
    function transferOwnership(address newOwner) external;
}

contract ERC721Proxy {
    address public implementation;
    address public owner;
    mapping(address => mapping(address => bool)) public authorized;
    
    constructor(address _implementation) {
        implementation = _implementation;
        owner = msg.sender;
    }
    
    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }
    
    function createNewERC721(string memory _name, string memory _symbol, address _creator) public onlyOwner returns (address) {
        BiosurfERC721 newContract = new BiosurfERC721(_name, _symbol, _creator);
        newContract.setTransferability(false);
        authorized[_creator][address(newContract)] = true;
        return address(newContract);
    }
    
    function upgradeImplementation(address _newImplementation) public onlyOwner {
        implementation = _newImplementation;
    }
    
    function transferOwnership(address _contractAddress) public onlyOwner {
        require(authorized[msg.sender][_contractAddress], "Not authorized");
        BiosurfERC721 erc721Contract = BiosurfERC721(_contractAddress);
        erc721Contract.transferOwnership(msg.sender);
        authorized[msg.sender][_contractAddress] = false;
        erc721Contract.setTransferability(true);
    }
}

contract BiosurfERC721 is ERC721, IOwnable {
    bool public transferable;
    address public creator;
    uint256 public creationTime;
    address public owner;
    
    constructor(string memory _name, string memory _symbol, address _creator) ERC721(_name, _symbol) {
        transferable = false;
        creator = _creator;
        creationTime = block.timestamp;
        owner = msg.sender;
    }
    
    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }
    
    function setTransferability(bool _transferable) public onlyOwner {
        transferable = _transferable;
    }

    function transferFrom(address from, address to, uint256 tokenId) public override {
        require(transferable, "Token transfer not allowed");
        super.transferFrom(from, to, tokenId);
    }
    
    function transferOwnership(address to) public onlyOwner {
        require(block.timestamp >= creationTime + 4 * 365 days, "Ownership transfer not available yet");
        owner = to;
        IOwnable(to).transferOwnership(to);
    }
    
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public override {
        if (Address.isContract(to)){
            require(transferable, "Token transfer not allowed");
            require(_isApprovedOrOwner(_msgSender(), tokenId), "Transfer caller is not owner nor approved");
            _safeTransfer(from, to, tokenId, data);
            IOwnable(to).transferOwnership(_msgSender()); 
        } else {
            super.safeTransferFrom(from, to, tokenId, data);
        }
    }
    
    function safeTransferFrom(address from, address to, uint256 tokenId) public override {
        if (Address.isContract(to)) {
            require(transferable, "Token transfer not allowed");
            require(_isApprovedOrOwner(_msgSender(), tokenId), "Transfer caller is not owner nor approved");
            _safeTransfer(from, to, tokenId, "");
            IOwnable(to).transferOwnership(_msgSender());
        } else {
            super.safeTransferFrom(from, to, tokenId);
        }
    }
}