pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract ERC721Proxy {
    address public implementation;
    mapping(address => mapping(address => bool)) public authorized;
    
    constructor(address _implementation) {
        implementation = _implementation;
    }
    
    function createNewERC721(string memory _name, string memory _symbol, address _creator) public returns (address) {
        ERC721 newContract = new BiosurfERC721(_name, _symbol, _creator);
        newContract.setTransferability(false);
        authorized[_creator][address(newContract)] = true;
        return address(newContract);
    }
    
    function upgradeImplementation(address _newImplementation) public {
        implementation = _newImplementation;
    }
    
    function transferOwnership(address _contractAddress) public {
        require(authorized[msg.sender][_contractAddress], "Not authorized");
        BiosurfERC721 erc721Contract = BiosurfERC721(_contractAddress);
        require(erc721Contract.creationTime() + 4 years <= block.timestamp, "Not eligible for ownership transfer yet");
        erc721Contract.transferOwnership(msg.sender);
        authorized[msg.sender][_contractAddress] = false;
        erc721Contract.setTransferability(true);
    }
}

contract BiosurfERC721 is ERC721 {
    bool public transferable;
    address public creator;
    uint256 public creationTime;
    
    constructor(string memory _name, string memory _symbol, address _creator) ERC721(_name, _symbol) {
        transferable = false;
        creator = _creator;
        creationTime = block.timestamp;
    }
    
    function setTransferability(bool _transferable) public {
        transferable = _transferable;
    }
    
    function transferFrom(address from, address to, uint256 tokenId) public override {
        require(transferable, "Token transfer not allowed");
        super.transferFrom(from, to, tokenId);
    }
    
    function transferOwnership(address newOwner) public override onlyOwner {
        require(block.timestamp >= creationTime + 4 years, "Ownership transfer not available yet");
        super.transferOwnership(newOwner);
    }
    
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public override {
        if (Address.isContract(to)) {
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

interface IOwnable {
    function transferOwnership(address newOwner) external;
}