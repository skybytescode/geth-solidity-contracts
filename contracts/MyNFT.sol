// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @title MyNFT
/// @notice An ERC-721 collection built on OpenZeppelin. Only the owner can
/// mint; transfers need the token owner or an approved address. The earlier
/// hand-written version let anyone mint and anyone transfer any token.
contract MyNFT is ERC721, Ownable {
    uint256 private _nextTokenId = 1;
    string private _baseTokenURI;

    constructor(string memory baseURI_, address initialOwner)
        ERC721("MyNFT", "MNFT")
        Ownable(initialOwner)
    {
        _baseTokenURI = baseURI_;
    }

    /// @notice Mints the next token to `to` and returns its ID.
    function mint(address to) external onlyOwner returns (uint256 tokenId) {
        tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
    }

    /// @notice Number of tokens minted so far.
    function totalSupply() external view returns (uint256) {
        return _nextTokenId - 1;
    }

    function _baseURI() internal view override returns (string memory) {
        return _baseTokenURI;
    }
}
