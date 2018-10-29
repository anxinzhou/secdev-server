pragma solidity ^0.4.25;

import './utils/gameToken.sol';
import './utils/helpers.sol';


contract BridgeToken is GameToken {

    // A signature collection
    struct SignaturesCollection {
        bytes message;
        // signed machines
        address[] authorizedMachines;
        // signatures of machines
        bytes[] signatures;
    }

    // Need required amount of signatures of machines to reach concensus.
    uint256 public requiredSignatures;

    mapping(bytes32 => address[]) internal deposits;

    mapping (bytes32 => SignaturesCollection) internal signatures;

    event WithdrawSignatureSubmitted(bytes32 messageHash);
    event CollectedSignatures(address authorityMachineResponsibleForRelay, bytes32 messageHash);
    event DepositConfirmation(address recipient, uint256 value, bytes32 transactoinHash);
    event Exchange(address user, uint amount);
    event Pay(address indexed user, uint amount , bytes32 transactoinHash);
    event ExchangeNFT(uint256 tokenID, address owner, uint256 gene, uint256 avatarLevel, bool weaponed, bool armored);

    constructor (uint256 totalSupply,
                string tokenName,
                string tokenSymbol,
                uint8 decimalUnits,
                uint _requiredSignatures
    ) GameToken(totalSupply,tokenName,tokenSymbol,decimalUnits) public {
        requiredSignatures = _requiredSignatures;
        _owner = msg.sender;
    }

    function setRequiredSignatures(uint newRequiredSignatures) public onlyOwner(){
        requiredSignatures = newRequiredSignatures;
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


    function exchange(address user, uint amount) public {
        _transfer(user, _owner, amount);
        emit Exchange(user, amount);
    }

    function pay(address user, uint amount, bytes32 transactionHash) public onlyAuthorizedMachine() {
        bytes32 hash = keccak256(abi.encodePacked(user, amount, transactionHash));

        require(!Helpers.addressArrayContains(deposits[hash], msg.sender));

        deposits[hash].push(msg.sender);

        if(deposits[hash].length == requiredSignatures) {
            _transfer(_owner, user, amount);
            emit Pay(user, amount, transactionHash);
        } else {
            emit DepositConfirmation(user, amount, transactionHash);
        }
    }

    function submitSignature(bytes message, bytes signature) public onlyAuthorizedMachine(){
        //require(msg.sender == MessageSigning.recoverAddressFromSignedMessage(signature, message));

        bytes32 hash = keccak256(message);

        // each authority can only provide one signature per message
        require(!Helpers.addressArrayContains(signatures[hash].authorizedMachines, msg.sender));

        signatures[hash].message = message;
        signatures[hash].authorizedMachines.push(msg.sender);
        signatures[hash].signatures.push(signature);

        if (signatures[hash].authorizedMachines.length == requiredSignatures) {
            emit CollectedSignatures(msg.sender, hash);
        } else {
            emit WithdrawSignatureSubmitted(hash);
        }
    }

    function signature(bytes32 messageHash, uint256 index) public view returns (bytes) {
        return signatures[messageHash].signatures[index];
    }

    function message(bytes32 message_hash) public view returns (bytes) {
        return signatures[message_hash].message;
    }
}
