pragma solidity ^0.5.2;

contract IERC20 {
    function transferFrom(address _from, address _to, uint _value) external returns (bool success);
    function transfer(address to, uint value) public returns (bool success);
    function transfer(address to, uint value, bytes memory data) public returns (bool success);
    function allowance(address _owner, address _spender) external view returns (uint256 remaining);
    function mint(address owner, uint value) external;
    function burn(address owner, uint value) external;
}
