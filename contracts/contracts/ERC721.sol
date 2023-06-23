// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Pausable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract SMSUANCES is
    ERC721,
    ERC721Burnable,
    ERC721Pausable,
    ERC721URIStorage,
    AccessControl
{
    address private _admin;
    address private _next;
    uint256 public constant TRANSFER_LOCK_TIME = 5 * 365 days;

    mapping(uint256 => uint256) private _tokenLock;

    constructor(
        string memory name,
        string memory symbol,
        address admin,
        address next
    ) ERC721(name, symbol) {
        _admin = admin;
        _next = next;
    }

    modifier onlyAdmin() {
        require(msg.sender == _admin, "SMSUANCES: caller is not the admin");
        _;
    }

    function ownerOf(uint256 tokenId) public view override returns (address) {
        return super.ownerOf(tokenId);
    }

    function tokenIsTransferable(uint256 tokenId) public view returns (bool) {
        return
            this.tokenIsTransferable(tokenId) &&
            (block.timestamp >= _tokenLock[tokenId]);
    }

    function contractName() public pure returns (string memory) {
        return "SMSUANCES";
    }

    function contractSymbol() public pure returns (string memory) {
        return "SMS";
    }

    mapping(address => bool) public authorized;

    function authorize(address user, address contractAddress) public onlyAdmin {
        authorized[contractAddress] = true;
    }

    function revokeAuthorization(address contractAddress) public onlyAdmin {
        authorized[contractAddress] = false;
    }

    function createNewToken(address to) public onlyAdmin returns (uint256) {
        uint256 newTokenId = createNewToken(_next);
        _tokenLock[newTokenId] = block.timestamp + TRANSFER_LOCK_TIME;
        return newTokenId;
    }

    function burn(uint256 tokenId) public override onlyAdmin {
        super.burn(tokenId);
        _tokenLock[tokenId] = 0;
    }

    function setTokenTransferability(
        uint256 tokenId,
        bool transferable
    ) public onlyAdmin {
        require(
            block.timestamp >= _tokenLock[tokenId],
            "SMSUANCES: token is locked for transfer"
        );

        this.setTokenTransferability(tokenId, transferable);
    }

    function isApprovedForAll(
        address owner,
        address operator
    ) public view override returns (bool) {
        if (authorized[operator]) {
            return true;
        }
        return super.isApprovedForAll(owner, operator);
    }

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override {
        require(
            _isApprovedOrOwner(_msgSender(), tokenId),
            "ERC721: transfer caller is not owner nor approved"
        );
        require(
            this.tokenIsTransferable(tokenId),
            "ERC721: token is not transferable"
        );

        super.transferFrom(from, to, tokenId);
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public virtual override {
        require(
            _isApprovedOrOwner(_msgSender(), tokenId),
            "ERC721: transfer caller is not owner nor approved"
        );
        require(
            this.tokenIsTransferable(tokenId),
            "ERC721: token is not transferable"
        );

        super.safeTransferFrom(from, to, tokenId);
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory _data
    ) public virtual override {
        require(
            _isApprovedOrOwner(_msgSender(), tokenId),
            "ERC721: transfer caller is not owner nor approved"
        );
        require(
            this.tokenIsTransferable(tokenId),
            "ERC721: token is not transferable"
        );

        super.safeTransferFrom(from, to, tokenId, _data);
    }

    function _approve(address to, uint256 tokenId) internal override {
        super._approve(to, tokenId);
        emit ApprovalForAll(ownerOf(tokenId), to, true);
    }
}
