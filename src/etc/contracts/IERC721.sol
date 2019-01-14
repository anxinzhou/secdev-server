pragma solidity ^0.5.2;

contract IERC721 {
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    function balanceOf(address owner) public view returns (uint256 );
    function ownerOf(uint256 tokenId) public view returns (address );
    function approve(address to, uint256 tokenId) public;
    function getApproved(uint256 tokenId) public view returns (address );

    function setApprovalForAll(address operator, bool _approved) public;
    function isApprovedForAll(address owner, address operator) public view returns (bool);

    function transferFrom(address from, address to, uint256 tokenId) public;
    function safeTransferFrom(address from, address to, uint256 tokenId) public;

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public;
    function onERC721Received(address operator, address from, uint256 tokenId, bytes memory data) public returns (bytes4);
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function tokenURI(uint256 tokenId) external view returns (string memory);

    //for bridge
    function mint(address to, uint256 tokenId, string memory uri) public;

    //for bridge
    function burn(uint256 tokenId) external;
}