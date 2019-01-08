pragma solidity ^0.4.25;

interface ERC721Metadata /* is ERC721 */ {
    /// @notice A descriptive name for a collection of NFTs in this contract
    function name() external view returns (string _name);

    /// @notice An abbreviated name for NFTs in this contract
    function symbol() external view returns (string _symbol);

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    /// @dev Throws if `_tokenId` is not a valid NFT. URIs are defined in RFC
    ///  3986. The URI may point to a JSON file that conforms to the "ERC721
    ///  Metadata JSON Schema".
    function tokenURI(uint256 _tokenId) external view returns (string);
}

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
    function append(string a, string b) internal pure returns (string) {
        return string(abi.encodePacked(a, b));
    }

    function uint2str(uint i) internal pure returns (string){
        if (i == 0) return "0";
        uint j = i;
        uint length;
        while (j != 0){
            length++;
            j /= 10;
        }
        bytes memory bstr = new bytes(length);
        uint k = length - 1;
        while (i != 0){
            bstr[k--] = byte(48 + i % 10);
            i /= 10;
        }
        return string(bstr);
    }


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
        string tokenSymbol
    ) public {
        decimals = 18;
        _totalSupply = totalSupply* 10**uint(decimals);
        _balances[msg.sender] = _totalSupply;
        name = tokenName;
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
        require(value <= _balances[from]);
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
    uint public exchangeRate = 500;
    uint public exchangeBase = 1;

    event Mint(address user, uint amount);
    event Burn(address user, uint amount);

    constructor (
        uint256 totalSupply,
        string tokenName,
        string tokenSymbol
    ) BasicToken(totalSupply,tokenName,tokenSymbol) payable public {
        _owner = msg.sender;
    }

    function mint(address user, uint amount) public returns(bool){
        _balances[user]+=amount;
    }

    function burn(address user, uint amount) public returns(bool) {
        require(_balances[user]>=amount);
        _balances[user]-=amount;
    }

    function exchangeForToken(address user) public payable returns (bool) {
        uint tokenAmount = msg.value * exchangeRate/exchangeBase;
        _balances[user] +=tokenAmount;
    }

    function exchangeForEther(address user, uint amount) public returns (bool) {
        uint etherAmount = amount*exchangeBase/exchangeRate;
        require(_balances[user]>=amount);
        _balances[user]-=amount;
        user.transfer(etherAmount);
    }

    function reward(address to, uint256 value) public returns (bool) {
        _balances[to] +=value;
        return true;
    }

    function consume(address by, uint256 value) public returns (bool){
        require(_balances[by]>=value);
        _balances[by] -=value;
        return true;
    }

    mapping(uint256=>string) _tokenURI;
    mapping(address=>uint256) ownedToken;
    mapping(uint256=>address) ownerOfToken;

    string private _name = "avatar";
    string private _symbol = "AVT";
    string private _baseURI = "https://secdev.site/";

    function NFTName() external view returns (string) {
        return _name;
    }

    /// @notice An abbreviated name for NFTs in this contract
    function NFTSymbol() external view returns (string) {
        return _symbol;
    }

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    /// @dev Throws if `_tokenId` is not a valid NFT. URIs are defined in RFC
    ///  3986. The URI may point to a JSON file that conforms to the "ERC721
    ///  Metadata JSON Schema".
    function tokenURI(uint256 _tokenId) external view returns (string) {
        return _tokenURI[_tokenId];
    }

    function create(address _user, uint256 _tokenID) external returns (bool) {
        ownerOfToken[_tokenID] = _user;
        ownedToken[_owner] = _tokenID;
        _tokenURI[_tokenID] = Helpers.append(_baseURI,Helpers.uint2str(_tokenID));
    }
}