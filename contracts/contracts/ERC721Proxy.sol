// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC721.sol";

contract ProxyProtocol {
    address private _admin;
    mapping(address => bool) private _userCreatedContract;
    address[] private _createdContracts;

    constructor() {
        _admin = msg.sender;
    }
-
    function implementation() public returns (address) {
        require(
            !_userCreatedContract[msg.sender],
            "User has already created a contract"
        );
        address INEW = address(
            new SMSUANCES("SMSUANCES", "SMS", _admin, msg.sender)
        );
        _createdContracts.push(INEW);
        _userCreatedContract[msg.sender] = true;
        return INEW;
    }

    function getContractCounts() public view returns (uint256) {
        return _createdContracts.length;
    }

    function getContractAddresses() public view returns (address[] memory) {
        return _createdContracts;
    }

    function getSmsuancesBalance(
        address userAddress,
        address nftContractAddress
    ) public view returns (uint256) {
        ERC721 nftContract = ERC721(nftContractAddress);
        return nftContract.balanceOf(userAddress);
    }
}
