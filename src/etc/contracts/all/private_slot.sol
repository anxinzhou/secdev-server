pragma solidity ^0.4.21;

contract Slot{

    string internal  name="slot";
    string internal symbol="slt";
    uint internal  decimals=0;
    uint internal  totalSupply;
    uint internal transferRate=100;
    mapping (address => uint) internal balances;
    mapping (address => mapping (address => uint)) internal allowed;

    event Transfer(address indexed _from,address indexed _to,uint _value);
    event Approval(address indexed _owner,address indexed _spender,uint _value);

    constructor (uint initialSupply ) 
    public {
        totalSupply = initialSupply;
        balances[msg.sender] = totalSupply;
    }

    function balanceOf(address _owner)external view returns (uint _balance){
        _balance = balances[_owner];
    }

    function _transfer(address _from, address _to, uint _value) private returns (bool) {
        require(_value <= balances[_from]);
        require(balances[_to] + _value >= balances[_to]);
        balances[_from] -= _value;
        balances[_to] += _value;
        emit Transfer(_from, _to, _value);
	return true;
    }

    function transfer(address _to,uint _value) external returns (bool){
	return _transfer(msg.sender, _to, _value);
    }

    function transferFrom(address _from, address _to, uint _value) external returns (bool){
        require(_value <= allowed[_from][msg.sender]);
	return _transfer(_from, _to, _value);
    }

    function approve(address _spender,uint _value)external returns (bool _success){
        allowed[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        _success = true;
    }

    function allowance(address _owner,address _spender) external view returns (uint _remaining){
        _remaining = allowed[_owner][_spender];
    }

}

library Message {
    // layout of message :: bytes:
    // offset 0: 20 bytes :: address - recipient address
    // offset 20: 32 bytes :: uint256 (big endian) - value
    // offset 52: 32 bytes :: bytes32 - transaction hash


    // mload always reads 32 bytes.
    // if mload reads an address it only interprets the last 20 bytes as the address.
    // so we can and have to start reading recipient at offset 20 instead of 32.
    // if we were to read at 32 the address would contain part of value and be corrupted.
    // when reading from offset 20 mload will ignore 12 bytes followed
    // by the 20 recipient address bytes and correctly convert it into an address.
    // this saves some storage/gas over the alternative solution
    // which is padding address to 32 bytes and reading recipient at offset 32.
    // for more details see discussion in:
    // https://github.com/paritytech/parity-bridge/issues/61

    function getRecipients(bytes message, uint num) internal pure returns (address [] memory) {
        address [] memory recipients = new address[](num);
        // solium-disable-next-line security/no-inline-assembly
        address recipient;
        uint offset = 52;
        for (uint i=0;i<num;++i){
            assembly {
                recipient := mload(add(message, offset))
            }
            recipients[i] = recipient;
            offset += 52;
        }
        return recipients;
    }

    function getValues(bytes message, uint num) internal pure returns (uint256 [] memory) {
        uint256 [] memory values = new uint256[](num);
        // solium-disable-next-line security/no-inline-assembly
        uint256 value;
        uint offset = 84;
        for(uint i=0; i< num; ++i){
            assembly {
                value := mload(add(message, offset))
            }
            values[i]=value;
            offset+=52;
        }
        return values;
    }

    function getTransactionHash(bytes message) internal pure returns (bytes32) {
        bytes32 hash;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            hash := mload(add(message, 32))
        }
        return hash;
    }
}

library MessageSigning {
    function recoverAddressFromSignedMessage(bytes signature, bytes message) internal pure returns (address) {
        require(signature.length == 65);
        bytes32 r;
        bytes32 s;
        bytes1 v;
        // solium-disable-next-line security/no-inline-assembly
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

library Helpers {
    /// returns whether `array` contains `value`.
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

    /// returns whether signatures (whose components are in `vs`, `rs`, `ss`)
    /// contain `requiredSignatures` distinct correct signatures
    /// where signer is in `allowed_signers`
    /// that signed `message`
    function hasEnoughValidSignatures(bytes message, uint8[] vs, bytes32[] rs, bytes32[] ss, address[] allowed_signers, uint256 requiredSignatures) internal pure returns (bool) {
        // not enough signatures
        if (vs.length < requiredSignatures) {
            return false;
        }

        bytes32 hash = MessageSigning.hashMessage(message);
        address []memory encountered_addresses = new address[](allowed_signers.length);

        for (uint256 i = 0; i < requiredSignatures; i++) {
            address recovered_address = ecrecover(hash, vs[i], rs[i], ss[i]);
            // only signatures by addresses in `addresses` are allowed
            if (!addressArrayContains(allowed_signers, recovered_address)) {
                return false;
            }
            // duplicate signatures are not allowed
            if (addressArrayContains(encountered_addresses, recovered_address)) {
                return false;
            }
            encountered_addresses[i] = recovered_address;
        }
        return true;
    }

}

contract PrivateSlot is Slot{
    
    struct SignaturesCollection {
        /// Signed message.
        bytes message;
        /// Authorities who signed the message.
        address[] authorities;
        /// Signatures
        bytes[] signatures;
    }

    uint256 public requiredSignatures;
    address[] public authorities;
    mapping(bytes32 => address[]) deposits;
    mapping (bytes32 => SignaturesCollection) signatures;
    mapping(bytes32 => address[]) syncs;
    
    

    event WithdrawSignatureSubmitted(bytes32 messageHash);
    /// Collected signatures which should be relayed to main chain.
    event CollectedSignatures(address authorityResponsibleForRelay, bytes32 messageHash);
    
    event DepositConfirmation(address recipient, uint256 value, bytes32 transactoinHash);
    event Login(address indexed privateAdd, address indexed publicAdd, uint time);
    event Sync(address []  _from, address [] _to, uint [] _amount, uint [] time);
    event SyncConfirmation(address indexed _from, address indexed _to, uint time);
    event Pay(address indexed _user, uint _amount , bytes32 transactoinHash);
    event Award(address indexed _user, uint _amount );
    event Consume(address indexed _user, uint _amount);
    event ExchangeToken(address user, uint amount);

    constructor (uint initialSupply, uint _requiredSignatures, address[] _authorities) Slot(initialSupply)
    public {
        requiredSignatures = _requiredSignatures;
        authorities = _authorities;
    }
    
    modifier onlyAuthority(){
        require (Helpers.addressArrayContains(authorities, msg.sender));
        _;
    }

    function addAuthority(address newAuthority) public onlyAuthority(){
        authorities.push(newAuthority);
    }
    
    function setRequiredSignatures(uint newRequiredSignatures) public onlyAuthority(){
        requiredSignatures = newRequiredSignatures;
    }

    function getRequiredSignatures() view public returns(uint _requiredSignatures){
        _requiredSignatures = requiredSignatures;
    }

    function consume(address _user, uint _amount) public returns(bool _success){
        require(balances[_user]>=_amount);
        balances[_user] -= _amount;
        emit Consume(_user, _amount);
        _success = true;
    }
    
    function sync(address [] _user, uint [] _amount, uint [] _time) public {
        // var hash = keccak256(_user, time);
        // //require(!Helpers.addressArrayContains(syncs[hash],msg.sender));
        // syncs[hash].push(msg.sender);
        // if (syncs[hash].length != requiredSignatures){
        //     emit SyncConfirmation(_user, publicAddress[_user], time);
        //     return;
        // }else{
        //     emit Sync(_user, publicAddress[_user], balances[_user],time);
        //     balances[_user] = 0;
        // }
        uint num = _user.length;
        for (uint i = 0; i < num ; ++i) {
            require(balances[_user[i]]>= _amount[i]);
            balances[_user[i]] -= _amount[i];
        }
        emit Sync(_user, _user, _amount, _time);
    }

    function submitSignature(bytes message, bytes signature) public onlyAuthority(){

        require(msg.sender == MessageSigning.recoverAddressFromSignedMessage(signature, message));

        bytes32 hash = keccak256(message);

        // each authority can only provide one signature per message
        require(!Helpers.addressArrayContains(signatures[hash].authorities, msg.sender));
        signatures[hash].message = message;
        signatures[hash].authorities.push(msg.sender);
        signatures[hash].signatures.push(signature);

        // TODO: this may cause troubles if requiredSignatures len is changed
        if (signatures[hash].authorities.length == requiredSignatures) {
            emit CollectedSignatures(msg.sender, hash);
        } else {
            emit WithdrawSignatureSubmitted(hash);
        }
    } 

    function pay(address _user, uint _amount, bytes32 _transactionHash) public onlyAuthority() {
        bytes32 hash = keccak256(abi.encodePacked(_user, _amount, _transactionHash));
        deposits[hash].push(msg.sender);
        if(deposits[hash].length == requiredSignatures) {
            balances[_user] += _amount;
            emit Pay(_user, _amount, _transactionHash);
        } 

    }

    function award(address _user, uint _amount) public {
        //require(msg.sender==owner);
        require(balances[_user] + _amount>= balances[_user]);
        balances[_user] += _amount;
        //emit Award(_user, _amount);
    }

    /// Get signature
    function signature(bytes32 messageHash, uint256 index) public view returns (bytes) {
        return signatures[messageHash].signatures[index];
    }

    /// Get message
    function message(bytes32 message_hash) public view returns (bytes) {
        return signatures[message_hash].message;
    }

}

