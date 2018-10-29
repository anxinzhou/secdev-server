pragma solidity ^0.4.21;

library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
    // benefit is lost if 'b' is also tested.
    // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
    if (a == 0) {
      return 0;
    }

    uint256 c = a * b;
    require(c / a == b);

    return c;
  }

  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b > 0); // Solidity only automatically asserts when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold

    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b <= a);
    uint256 c = a - b;

    return c;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    require(c >= a);

    return c;
  }

  function mod(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b != 0);
    return a % b;
  }
}

library Helpers {
    // returns whether `array` contains `value`.
    function addressArrayContains(address[] array, address value) internal pure returns (bool) {
        for (uint256 i = 0; i < array.length; i++) {
            if (array[i] == value) {
                return true;
            }
        }
        return false;
    }

    // returns the digits of `inputValue` as a string.
    // example: `uintToString(12345678)` returns `"12345678"`
    function uintToString(uint256 inputValue) internal pure returns (string) {
        // figure out the length of the resulting string
        uint256 length = 0;
        uint256 currentValue = inputValue;
        do {
            length++;
            currentValue /= 10;
        } while (currentValue != 0);
        // allocate enough memory
        bytes memory result = new bytes(length);
        // construct the string backwards
        uint256 i = length - 1;
        currentValue = inputValue;
        do {
            result[i--] = byte(48 + currentValue % 10);
            currentValue /= 10;
        } while (currentValue != 0);
        return string(result);
    }
}


library Message {
    // layout of message :: bytes:
    // offset 32: 0-32:: bytes32 - transaction hash
    // offset 52: 32-52:: address recipient address
    // offset 84: 52-84:: uint256 (big endian) - value

    function getTransactionHash(bytes message) internal pure returns (bytes32) {
        bytes32 hash;
        assembly {
            hash := mload(add(message, 32))
        }
        return hash;
    }

    function getRecipients(bytes message) internal pure returns (address) {
        address recipient;
        assembly {
            recipient := mload(add(message, 52))
        }
        return recipient;
    }

    function getValues(bytes message) internal pure returns (uint256) {
        uint256 value;
        assembly {
            value := mload(add(message, 84))
        }
        return value;
    }
}


library MessageSigning {
    function recoverAddressFromSignedMessage(bytes signature, bytes message) internal pure returns (address) {
        require(signature.length == 65);
        bytes32 r;
        bytes32 s;
        bytes1 v;
        assembly {
            r := mload(add(signature, 0x20))
            s := mload(add(signature, 0x40))
            v := mload(add(signature, 0x60))
        }
        return ecrecover(hashMessage(message), uint8(v), r, s);
    }

    function hashMessage(bytes message) internal pure returns (bytes32) {
        bytes memory prefix = "\x19Ethereum Signed Message:\n";
        return keccak256(abi.encodePacked(prefix, Helpers.uintToString(message.length), message));
    }
}

contract BasicToken {
    using SafeMath for uint256;

    uint256 internal _totalSupply;

    string public name;

    uint8 public decimals;

    string public symbol;

    mapping (address => uint256) internal _balances;

    mapping (address => mapping (address => uint256)) internal _allowed;


    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    constructor (
        uint256 totalSupply,
        string tokenName,
        string tokenSymbol,
        uint8 decimalUnits
    ) public {
        _totalSupply = totalSupply;
        _balances[msg.sender] = totalSupply;
        name = tokenName;
        decimals = decimalUnits;
        symbol = tokenSymbol;
    }

    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address owner) public view returns (uint256) {
        return _balances[owner];
    }

    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowed[owner][spender];
    }

    function _transfer(address from, address to, uint256 value) internal {
        require(value <= _balances[msg.sender]);
        require(to != address(0));

        _balances[from] = _balances[from].sub(value);
        _balances[to] = _balances[to].add(value);
    }

    function transfer(address to, uint256 value) public returns (bool) {
        _transfer(msg.sender, to, value);
        emit Transfer(msg.sender, to, value);
        return true;
    }

    function approve(address spender, uint256 value) public returns (bool) {
        require(spender != address(0));

        _allowed[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool) {
        require(value <= _allowed[from][msg.sender]);
        _allowed[from][msg.sender] = _allowed[from][msg.sender].sub(value);

        _transfer(from, to, value);
        emit Transfer(from, to, value);
        return true;
    }
}

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
      bool armored;
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
    }

    function equipWeapon(uint256 tokenId, address user) public {
        require(_avatarOwner[tokenId]==user);
        require(!avatar[tokenId].weaponed);
        avatar[tokenId].weaponed = true;
    }

    function equipArmor(uint256 tokenId, address user) public {
        require(_avatarOwner[tokenId]==user);
        require(!avatar[tokenId].armored);
        avatar[tokenId].armored = true;
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

contract BridgeToken is GameToken{

    mapping (bytes32 => bool) public payed;
    uint256 public requiredSignatures;

    event Exchange(address user, uint amount);
    event Pay(address user, uint amount ,bytes32 transactionHash);
    event ExchangeNFT(uint256 tokenID, address owner, uint256 gene, uint256 avatarLevel, bool weaponed, bool armored);

    constructor (uint256 totalSupply,
                string tokenName,
                string tokenSymbol,
                uint8 decimalUnits,
                uint _requiredSignatures)
    public GameToken(totalSupply,tokenName,tokenSymbol,decimalUnits) {
        requiredSignatures = _requiredSignatures;
        _owner = msg.sender;

    }

    // returns whether signatures (whose components are in `vs`, `rs`, `ss`)
    // contain distinct correct signatures with a number of `requiredSignatures`
    // where signer is in `_authorizdedMachines`
    // that signed `message`
    function hasEnoughValidSignatures(bytes message, uint8[] vs, bytes32[] rs, bytes32[] ss) internal view returns (bool) {
        // not enough signatures
        if (vs.length < requiredSignatures) {
            return false;
        }

        bytes32 hash = MessageSigning.hashMessage(message);
        address [] memory encountered_addresses = new address[](vs.length);

        for (uint256 i = 0; i < requiredSignatures; i++) {
            address recovered_address = ecrecover(hash, vs[i], rs[i], ss[i]);
            // only signatures by addresses in `addresses` are allowed
            if (!_authorizdedMachines[recovered_address]) {
                return false;
            }
            // duplicate signatures are not allowed
            if (Helpers.addressArrayContains(encountered_addresses, recovered_address)) {
                return false;
            }
            encountered_addresses[i] = recovered_address;
        }
        return true;
    }

    function setRequiredSignatures(uint newRequiredSignatures) public onlyOwner(){
        requiredSignatures = newRequiredSignatures;
    }

    function exchange(address user,uint amount) public{
        _transfer(user, _owner, amount);
        emit Exchange(user, amount);
    }

    function exchangeNFT (uint256 tokenID) public {
        address avatarOwner = _avatarOwner[tokenID];
        require(msg.sender == avatarOwner);
        _ownedAvatars[avatarOwner]=0;
        _avatarOwner[tokenID]=0;
        uint256 gene = avatar[tokenID].gene ;
        uint256 avatarLevel = avatar[tokenID].avatarLevel;
        bool weaponed = avatar[tokenID].weaponed;
        bool armored = avatar[tokenID].armored;
        avatar[tokenID].gene=0;
        avatar[tokenID].avatarLevel = 0;
        avatar[tokenID].weaponed = false;
        avatar[tokenID].armored = false;
        emit ExchangeNFT(tokenID,avatarOwner, gene, avatarLevel, weaponed, armored);
    }

    function payNFT (uint256 tokenID, address avatarOwner, uint256 gene, uint256 avatarLevel, bool weaponed, bool armored) public {
        _ownedAvatars[avatarOwner]=tokenID;
        _avatarOwner[tokenID]=avatarOwner;
        avatar[tokenID].gene = gene;
        avatar[tokenID].avatarLevel = avatarLevel;
        avatar[tokenID].weaponed = weaponed;
        avatar[tokenID].armored = armored;
    }

    // vs, rs, ss is used for ecrecover(a built-in function of solidity)
    // to recover signer address
    //
    // message include
    // 32bytes -- bytes32 transaction hash
    // 20bytes -- address recipient address
    // 32bytes -- uint payment value
    function pay(uint8 []vs, bytes32 []rs, bytes32 []ss, bytes message) public{
        require(message.length == 84);

        // check that at least `requiredSignatures` `authorities` have signed `message`
        //require(hasEnoughValidSignatures(message, vs, rs, ss));

        address recipient = Message.getRecipients(message);
        uint256 value = Message.getValues(message);
        bytes32 hash = Message.getTransactionHash(message);

        require(!payed[hash]);

        payed[hash] = true;
        _transfer(_owner, recipient, value);

        emit Pay(recipient, value, hash);
    }
}