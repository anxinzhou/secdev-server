pragma solidity ^0.4.25;

import './utils.sol';
import './basic_token.sol';

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
