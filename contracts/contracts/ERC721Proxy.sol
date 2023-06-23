// SPDX-License-Identifier: UNLICENSED

pragma solidity^0.8.0;

import "./ERC721.sol";

contract SMSUANCESERC721Proxy is SMSUANCESERC721 {
    address private _admin;

    constructor(string memory name, string memory symbol, address admin) SMSUANCESERC721(name, symbol) {
        _admin = admin;
    }

    function updateAdmin(address newAdmin) external onlyAdmin {
        require(newAdmin != address(0), "BiosurfERC721Proxy: new admin is the zero address");
        _admin = newAdmin;
    }

    modifier onlyAdmin() {
        require(msg.sender == _admin, "BiosurfERC721Proxy: caller is not the admin");
        _;
    }

    function ownerOf(uint256 tokenId) public view override returns (address) {
        return super.ownerOf(tokenId);
    }

    function tokenIsTransferable(uint256 tokenId) public view returns (bool) {
        return super.tokenIsTransferable(tokenId);
    }

    function contractName() public pure returns (string memory) {
        return "SMSUANCESERC721";
    }

    function contractSymbol() public pure returns (string memory) {
        return "SMS";
    }

    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    mapping(address => bool) public authorized;

    function authorize(address user, address contractAddress) public onlyAdmin {
        authorized[contractAddress] = true;
    }

    function revokeAuthorization(address contractAddress) public onlyAdmin {
        authorized[contractAddress] = false;
    }

    function createNewToken(address to, bool transferable) public onlyAdmin returns (uint256) {
        return super.createNewToken(to, transferable);
    }

    function burn(uint256 tokenId) public onlyAdmin {
        super.burn(tokenId);
    }

    function setTokenTransferability(uint256 tokenId, bool transferable) public onlyAdmin {
        super.setTokenTransferability(tokenId, transferable);
    }

    function isApprovedForAll(address owner, address operator) public view override returns (bool) {
        if (authorized[operator]) {
            return true;
        }
        return super.isApprovedForAll(owner, operator);
    }

    function transferFrom(address from, address to, uint256 tokenId) public override {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        require(tokenIsTransferable(tokenId), "ERC721: token is not transferable");

        super.transferFrom(from, to, tokenId);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public override {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        require(tokenIsTransferable(tokenId), "ERC721: token is not transferable");

        super.safeTransferFrom(from, to, tokenId, _data);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId) public override {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        require(tokenIsTransferable(tokenId), "ERC721: token is not transferable");

        super.safeTransferFrom(from, to, tokenId);
    }

    function _approve(address to, uint256 tokenId) internal override {
        super._approve(to, tokenId);
        emit ApprovalForAll(ownerOf(tokenId), to, true);
    }
}