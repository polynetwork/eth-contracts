pragma solidity >=0.5.0;

import "../libs/token/ERC721/ERC721.sol";
import "../libs/utils/Address.sol";

contract CrossChainNFTMapping is ERC721 {
    using Address for address; 

    address private lpAddr;

    constructor (address _lpAddr, string memory name, string memory symbol) public ERC721(name, symbol) {
        require(_lpAddr.isContract(), "lockproxy address must be contract.");
        lpAddr = _lpAddr;
    }

    modifier onlyProxy() {
        require(msg.sender == lpAddr, "");
        _;
    }

    function mintWithURI(address to, uint256 tokenId, string memory uri) public onlyProxy {
        require(!_exists(tokenId), "token id already exist");
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }
}