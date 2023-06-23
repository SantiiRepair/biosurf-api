// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Pausable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract SMSUANCESERC721 is ERC721, ERC721Burnable, ERC721Pausable, ERC721URIStorage, AccessControl {
    using Counters for Counters.Counter;

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    Counters.Counter private _tokenIdCounter;

    mapping(uint256 => bool) private _tokenIsTransferable;

    constructor(string memory name, string memory symbol) ERC721(name, symbol) {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(PAUSER_ROLE, msg.sender);
        _setupRole(MINTER_ROLE, msg.sender);
    }

    function pause() public {
        require(hasRole(PAUSER_ROLE, msg.sender));
        _pause();
    }

    function unpause() public {
        require(hasRole(PAUSER_ROLE, msg.sender));
        _unpause();
    }

    function mint(address to, string memory tokenURI) public {
        require(hasRole(MINTER_ROLE, msg.sender));
        _tokenIdCounter.increment();
        uint256 newTokenId = _tokenIdCounter.current();
        _mint(to, newTokenId);
        _setTokenURI(newTokenId, tokenURI);
    }

    function ownerOf(uint256 tokenId) public view override returns (address) {
        return super.ownerOf(tokenId);
    }

    function tokenIsTransferable(uint256 tokenId) public view returns (bool) {
        return _tokenIsTransferable[tokenId];
    }

    function contractName() public pure returns (string memory) {
        return "BiosurfERC721";
    }

    function contractSymbol() public pure returns (string memory) {
        return "BSF";
    }

    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    mapping(address => bool) public authorized;

    function authorize(address user, address contractAddress) public onlyOwner {
        authorized[contractAddress] = true;
    }

    function revokeAuthorization(address contractAddress) public onlyOwner {
        authorized[contractAddress] = false;
    }

    function createNewToken(address to, bool transferable) public onlyOwner returns (uint256) {
        _tokenIdCounter.increment();
        uint256 newTokenId = _tokenIdCounter.current();
        _mint(to, newTokenId);
        _tokenIsTransferable[newTokenId] = transferable;
        return newTokenId;
    }

    function burn(uint256 tokenId) public onlyOwner {
        _burn(tokenId);
    }

    function setTokenTransferability(uint256 tokenId, bool transferable) public onlyOwner {
        _tokenIsTransferable[tokenId] = transferable;
    }

    function isApprovedForAll(address owner, address operator) public view override returns (bool) {
        if (authorized[operator]) {
            return true;
        }
        return super.isApprovedForAll(owner, operator);
    }

    function transferFrom(address from, address to, uint256 tokenId) public override {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        require(_tokenIsTransferable[tokenId], "ERC721: token is not transferable");

        _transfer(from, to, tokenId);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public override {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        require(_tokenIsTransferable[tokenId], "ERC721: token is not transferable");

        _safeTransfer(from, to, tokenId, _data);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId) public override {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfercaller is not owner nor approved");
        require(_tokenIsTransferable[tokenId], "ERC721: token is not transferable");

        _safeTransfer(from, to, tokenId, "");
    }

    function _approve(address to, uint256 tokenId) internal override {
        super._approve(to, tokenId);
        emit ApprovalForAll(ownerOf(tokenId), to, true);
    }
}