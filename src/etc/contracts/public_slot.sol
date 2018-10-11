pragma solidity ^0.4.25;

import './utils.sol';
import './basic_token.sol';

contract PublicSlot is Slot{


    mapping (bytes32 => bool) public pays;
    uint256 public requiredSignatures;
    address[] public authorities;
    mapping(bytes32 => address[]) deposits;

    event Exchange(address indexed _from, address indexed _to, uint _amount);
    event Pay(bytes32 transactionHash);
    event DepositConfirmation(address recipient, uint256 value, bytes32 transactionHash);

    constructor (uint initialSupply, uint _requiredSignatures, address[] _authorities) Slot(initialSupply)
    public {
        requiredSignatures = _requiredSignatures;
        authorities = _authorities;
    }

    modifier onlyAuthority(){
        require (Helpers.addressArrayContains(authorities, msg.sender));
        _;
    }

    function exchange(address _to,uint _amount) public{
        require(balances[msg.sender]>=_amount);
        balances[msg.sender] -= _amount;
        emit Exchange(msg.sender, _to, _amount);
    }
    
    
    function pay(uint8[] vs, bytes32[] rs, bytes32[] ss, bytes message) public {
        uint num = (message.length-32) / (32+20);

        // check that at least `requiredSignatures` `authorities` have signed `message`
        require(Helpers.hasEnoughValidSignatures(message, vs, rs, ss, authorities, requiredSignatures));

        address [] memory recipients = Message.getRecipients(message,num);
        uint256 [] memory values = Message.getValues(message, num);
        bytes32 hash = Message.getTransactionHash(message);


        require(!pays[hash]);
        // Order of operations below is critical to avoid TheDAO-like re-entry bug
        pays[hash] = true;
        for(uint i=0;i<num;++i){
            balances[recipients[i]]+=values[i];
        }
        // pay out recipient
        emit Pay(hash);
    }
    
}
