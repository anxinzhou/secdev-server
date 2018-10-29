pragma solidity ^0.4.25;

import './utils/gameToken.sol';
import './utils/helpers.sol';

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

