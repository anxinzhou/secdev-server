pragma solidity ^0.4.25;

import './basicToken.sol';

contract GameToken is BasicToken {
    // creator of this contract
    address internal _owner;

    // only authorized game machine can consume and reward user token
    mapping (address => bool) public _authorizdedMachines;


    // erc721
    struct Avatar {
      uint256 gene;
      uint256 avatarLevel;
      bool weaponed;
    }

    uint constant internal MAXLEVEL= 2;

    mapping (uint256 => address) internal _avatarOwner;

    mapping (uint256 => Avatar) public avatar;

    mapping (address => uint256) internal _ownedAvatars;



    event Reward(address indexed machine, address indexed player, uint256 value);
    event Consume(address indexed machine, address indexed player, uint256 value);

    constructor (
        uint256 totalSupply,
        string tokenName,
        string tokenSymbol,
        uint8 decimalUnits
    ) BasicToken(totalSupply,tokenName,tokenSymbol,decimalUnits) public {
        _owner = msg.sender;
    }

    modifier onlyOwner()  {
        require(msg.sender == _owner);
        _;
    }

    modifier onlyAuthorizedMachine()  {
        require(_authorizdedMachines[msg.sender]);
        _;
    }

    function ownerOf(uint256 tokenId) public view returns (address) {
        address owner = _avatarOwner[tokenId];
        require(owner != address(0));
        return owner;
    }

    function ownedAvatars(address owner) public view returns (uint256){
        return _ownedAvatars[owner];
    }

    function mint(address to, uint256 tokenId) public {
        require(to !=address(0));
        _avatarOwner[tokenId] = to ;
        _ownedAvatars[to]= tokenId;
        avatar[tokenId].gene= now %2;
    }

    function upgrade(uint256 tokenId) public {
        require(avatar[tokenId].avatarLevel < MAXLEVEL);
        avatar[tokenId].avatarLevel +=1;
        avatar[tokenId].weaponed = false;
    }

    function equipWeapon(uint256 tokenId) public {
        require(!avatar[tokenId].weaponed);
        avatar[tokenId].weaponed = true;
    }

    function addGameMachine(address machine) public onlyOwner() {
        _authorizdedMachines[machine] = true;
    }

    function removeGameMachine(address machine) public onlyOwner() {
        _authorizdedMachines[machine] = false;
    }

    function reward(address to, uint256 value) public onlyAuthorizedMachine() returns (bool) {
        _transfer(msg.sender, to, value);
        emit Reward(msg.sender, to, value);
        return true;
    }

    function consume(address by, uint256 value) public onlyAuthorizedMachine() returns (bool){
        _transfer(by, msg.sender, value);
        emit Consume(msg.sender, by, value);
        return true;
    }

}