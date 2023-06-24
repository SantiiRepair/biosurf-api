// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract SMSUANCES is ERC721, ERC721Burnable, ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;
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

    function waterdog() public returns (uint256) {
        uint256 newTokenId = _tokenIds.current();
        _mint(_next, newTokenId);
        _tokenLock[newTokenId] = block.timestamp + TRANSFER_LOCK_TIME;
        _tokenIds.increment();
        return newTokenId;
    }

    function custom(address to) public returns (uint256) {
        uint256 newTokenId = _tokenIds.current();
        _mint(to, newTokenId);
        _tokenIds.increment();
        return newTokenId;
    }

    function _burn(
        uint256 tokenId
    ) internal virtual override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
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
        emit Approval(_next, to, tokenId);
    }

    function tokenURI(
        uint256 tokenId
    )
        public
        view
        virtual
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}
