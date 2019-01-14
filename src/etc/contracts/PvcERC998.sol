pragma solidity ^0.5.2;

import "IERC721.sol";
import "IERC20.sol";
import "Address.sol";
import "SafeMath.sol";

contract ComposableTopDown  {
    using SafeMath for uint256;
    using Address for address;

    /**
    * For bridge development
    * Mapping tokenId from lock state
    */
    mapping(uint256 => bool) private _locked;

    // return this.rootOwnerOf.selector ^ this.rootOwnerOfChild.selector ^
    //   this.tokenOwnerOf.selector ^ this.ownerOfChild.selector;
    bytes4 constant ERC998_MAGIC_VALUE = 0xcd740db5;

    uint256 tokenCount = 0;
    // mapping from token ID to owner
    mapping(uint256 => address) private _tokenOwner;

    // Mapping from root token owner to tokenId to approved address
    mapping(address => mapping(uint256 => address)) private _tokenApprovals;

    // Mapping from token owner to token count
    mapping(address => uint256) private _ownedTokensCount;

    // Mapping from root token owner to operator approvals
    mapping(address => mapping(address => bool)) private _operatorApprovals;


    // Token name
    string private _name;

    // Token symbol
    string private _symbol;

    // Optional mapping for token URIs
    mapping(uint256 => string) private _tokenURIs;

    // ERC998 Event
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);
    event ReceivedERC20(address indexed _from, uint256 indexed _tokenId, address indexed _erc20Contract, uint256 _value);
    event TransferERC20(uint256 indexed _tokenId, address indexed _to, address indexed _erc20Contract, uint256 _value);
    event ReceivedChild(address indexed _from, uint256 indexed _tokenId, address indexed _childContract, uint256 _childTokenId);
    event TransferChild(uint256 indexed tokenId, address indexed _to, address indexed _childContract, uint256 _childTokenId);

    // Bridge Event
    event Lock(address operator ,address owner, uint256 tokenId);
    event Unlock(address operator, address owner, uint256 tokenId);

    /**
     * @dev Constructor function
     */
    constructor (string memory name, string memory symbol) public {
        _name = name;
        _symbol = symbol;
    }

    /**
     * for bridge developemnt
     *  query token lock state
     * if locked, not allowed to transfer or burn
    */
    function isLocked(uint256 tokenId) external view returns (bool) {
        return _locked[tokenId];
    }

    /**
    * for bridge developemnt
    *  lock
    * called by user
   */
    function lock(uint256 tokenId) external {
        address operator = msg.sender;
        address tokenOwner = _tokenOwner[tokenId];
        require(tokenOwner !=address(0), "token not exist");
        require(!_locked[tokenId], "token already locked");
        _locked[tokenId] = true;
        emit Lock(operator , tokenOwner , tokenId);
    }

    /**
     * for bridge developemnt
     * unlock Token
     */
    function unlock(uint256 tokenId) external {
        address operator = msg.sender;
        address tokenOwner = _tokenOwner[tokenId];
        require(tokenOwner !=address(0), "token not exist");
        require(_locked[tokenId], "token already unlocked");
        _locked[tokenId] = false;
        emit Unlock(operator, tokenOwner, tokenId);
    }

    /**
     * for bridge development
     * The token mint in public chain is locked initially
     * The token mint in private chain is unlocked initially
    */
    function mint(address to, uint256 tokenId, string memory uri) public {
        _mint(to, tokenId);
        _setTokenURI(tokenId,uri);
        _locked[tokenId] = false;
    }

    /**
     * for bridge developemnt
     *  query token lock state
     * Require multisig
    */
    function changeOwner(address owner, uint256 tokenId) external {
        _tokenOwner[tokenId] = owner;
    }

    // for ERC721 child bridge transfer

    /**
     * for bridge developemnt
     *  mint child Token
     * Require Multisig
    */

    function mintERC721Child(uint256 parentTokenId, address childContract ,uint256 childTokenId, string memory uri) public {
        IERC721(childContract).mint(address(this), childTokenId, uri);
        address operator = msg.sender;
        receiveChild(operator, parentTokenId, childContract, childTokenId);
    }

    /**
     * for bridge developemnt
     *  burn child Token
     * Require Multisig
    */

    function burnERC721Child(uint256 parentTokenId ,address childContract, uint256 childTokenId) external {
        IERC721(childContract).burn(childTokenId);
        removeChild(parentTokenId, childContract, childTokenId);
    }

    /**
     * for bridge developemnt
     *  mint child Token
     * Require Multisig
    */

    function mintERC20Child(uint256 parentTokenId, address childContract ,uint256 value) public {
        IERC20(childContract).mint(address(this), value);
        address operator = msg.sender;
        erc20Received(operator, parentTokenId, childContract, value);
    }

    /**
     * for bridge developemnt
     *  burn child Token
     * Require Multisig
    */

    function burnERC20Child(uint256 parentTokenId ,address childContract, uint256 value) external {
        IERC20(childContract).burn(address(this),value);
        removeERC20(parentTokenId, childContract, value);
    }

    //Test if token exist

    function _exists(uint256 tokenId) internal view returns (bool) {
        address owner = _tokenOwner[tokenId];
        return owner != address(0);
    }

    /**
     * @dev Gets the token name
     * @return string representing the token name
     */
    function name() external view returns (string memory) {
        return _name;
    }

    /**
     * @dev Gets the token symbol
     * @return string representing the token symbol
     */
    function symbol() external view returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns an URI for a given token ID
     * Throws if the token ID does not exist. May return an empty string.
     * @param tokenId uint256 ID of the token to query
     */
    function tokenURI(uint256 tokenId) external view returns (string memory) {
        require(_exists(tokenId));
        return _tokenURIs[tokenId];
    }

    /**
     * @dev Internal function to set the token URI for a given token
     * Reverts if the token ID does not exist
     * @param tokenId uint256 ID of the token to set its URI
     * @param uri string URI to assign
     */
    function _setTokenURI(uint256 tokenId, string memory uri) internal {
        require(_exists(tokenId));
        _tokenURIs[tokenId] = uri;
    }

    // for template development
    function setTokenURI(uint256 tokenId, string memory uri) public {
        require(_exists(tokenId));
        _tokenURIs[tokenId] = uri;
    }

    /**
     * for template develop use
    */
    function burn(uint256 tokenId) external {
        _burn(tokenId);
    }

    /**
     * @dev Internal function to mint a new token
     * Reverts if the given token ID already exists
     * @param to The address that will own the minted token
     * @param tokenId uint256 ID of the token to be minted
     */
    function _mint(address to, uint256 tokenId) internal {
        require(to != address(0));
        require(!_exists(tokenId));

        _tokenOwner[tokenId] = to;
        _ownedTokensCount[to] = _ownedTokensCount[to].add(1);

        emit Transfer(address(0), to, tokenId);
    }

    function _clearApproval(uint256 tokenId) internal {
        address owner = ownerOf(tokenId);
        _tokenApprovals[owner][tokenId]=address(0);
    }

    /**
     * @dev Internal function to burn a specific token
     * Reverts if the token does not exist
     * Deprecated, use _burn(uint256) instead.
     * @param owner owner of the token to burn
     * @param tokenId uint256 ID of the token being burned
     */
    function _burn(address owner, uint256 tokenId) internal {
        require(ownerOf(tokenId) == owner);

        _clearApproval(tokenId);

        _ownedTokensCount[owner] = _ownedTokensCount[owner].sub(1);
        _tokenOwner[tokenId] = address(0);

        if (bytes(_tokenURIs[tokenId]).length != 0) {
            delete _tokenURIs[tokenId];
        }

        emit Transfer(owner, address(0), tokenId);
    }

    /**
     * @dev Internal function to burn a specific token
     * Reverts if the token does not exist
     * @param tokenId uint256 ID of the token being burned
     */
    function _burn(uint256 tokenId) internal {
        _burn(ownerOf(tokenId), tokenId);
    }

    //from zepellin ERC721.sol
    bytes4 constant ERC721_RECEIVED = 0x150b7a02;

    ////////////////////////////////////////////////////////
    // ERC721 implementation
    ////////////////////////////////////////////////////////

    function rootOwnerOf(uint256 _tokenId) public view returns (bytes32) {
        return rootOwnerOfChild(address(0), _tokenId);
    }

    // returns the owner at the top of the tree of composables
    // Use Cases handled:
    // Case 1: Token owner is this contract and token.
    // Case 2: Token owner is other top-down composable
    // Case 3: Token owner is other contract
    // Case 4: Token owner is user
    function rootOwnerOfChild(address _childContract, uint256 _childTokenId) public view returns (bytes32) {
        bytes32 rootOwner;
        address rootOwnerAddress;
        if (_childContract != address(0)) {
            (rootOwnerAddress, _childTokenId) = _ownerOfChild(_childContract, _childTokenId);
        }
        else {
            rootOwnerAddress = _tokenOwner[_childTokenId];
        }
        // Case 1: Token owner is this contract and token.
        while (rootOwnerAddress == address(this)) {
            (rootOwnerAddress, _childTokenId) = _ownerOfChild(rootOwnerAddress, _childTokenId);
        }

        bool callSuccess;
        bytes memory _calldata;
        // 0xed81cdda == rootOwnerOfChild(address,uint256)
        _calldata = abi.encodeWithSelector(0xed81cdda, address(this), _childTokenId);
        assembly {
            callSuccess := staticcall(gas, rootOwnerAddress, add(_calldata, 0x20), mload(_calldata), _calldata, 0x20)
            if callSuccess {
                rootOwner := mload(_calldata)
            }
        }
        if(callSuccess == true && rootOwner >> 224 == ERC998_MAGIC_VALUE) {
            // Case 2: Token owner is other top-down composable
            return rootOwner;
        }
        else {
            // Case 3: Token owner is other contract
            // Or
            // Case 4: Token owner is user
            return bytes32(ERC998_MAGIC_VALUE) << 224 | bytes32(bytes20(rootOwnerAddress));
        }
    }


    // returns the owner at the top of the tree of composables

    function ownerOf(uint256 _tokenId) public view returns (address) {
        address tokenOwner = _tokenOwner[_tokenId];
        require(tokenOwner != address(0));
        return tokenOwner;
    }

    function balanceOf(address owner) external view returns (uint256) {
        require(owner != address(0));
        return _ownedTokensCount[owner];
    }


    function approve(address _approved, uint256 _tokenId) external {
        address rootOwner = address(bytes20(rootOwnerOf(_tokenId)));
        require(rootOwner == msg.sender || _operatorApprovals[rootOwner][msg.sender]);
        _tokenApprovals[rootOwner][_tokenId] = _approved;
        emit Approval(rootOwner, _approved, _tokenId);
    }

    function getApproved(uint256 _tokenId) public view returns (address)  {
        address rootOwner = address(bytes20(rootOwnerOf(_tokenId)));
        return _tokenApprovals[rootOwner][_tokenId];
    }

    function setApprovalForAll(address _operator, bool _approved) external {
        require(_operator != address(0));
        _operatorApprovals[msg.sender][_operator] = _approved;
        emit ApprovalForAll(msg.sender, _operator, _approved);
    }

    function isApprovedForAll(address _owner, address _operator) external view returns (bool)  {
        require(_owner != address(0));
        require(_operator != address(0));
        return _operatorApprovals[_owner][_operator];
    }


    function _transferFrom(address _from, address _to, uint256 _tokenId) private {
        require(_from != address(0));
        require(_tokenOwner[_tokenId] == _from);
        require(_to != address(0));

        if(msg.sender != _from) {
            bytes32 rootOwner;
            bool callSuccess;
            // 0xed81cdda == rootOwnerOfChild(address,uint256)
            bytes memory _calldata = abi.encodeWithSelector(0xed81cdda, address(this), _tokenId);
            assembly {
                callSuccess := staticcall(gas, _from, add(_calldata, 0x20), mload(_calldata), _calldata, 0x20)
                if callSuccess {
                    rootOwner := mload(_calldata)
                }
            }
            if(callSuccess == true) {
                require(rootOwner >> 224 != ERC998_MAGIC_VALUE, "Token is child of other top down composable");
            }
            require(_operatorApprovals[_from][msg.sender] ||
            _tokenApprovals[_from][_tokenId] == msg.sender);
        }

        // clear approval
        if (_tokenApprovals[_from][_tokenId] != address(0)) {
            delete _tokenApprovals[_from][_tokenId];
            emit Approval(_from, address(0), _tokenId);
        }

        // remove and transfer token
        if (_from != _to) {
            assert(_ownedTokensCount[_from] > 0);
            _ownedTokensCount[_from]--;
            _tokenOwner[_tokenId] = _to;
            _ownedTokensCount[_to]++;
        }
        emit Transfer(_from, _to, _tokenId);

    }

    function transferFrom(address _from, address _to, uint256 _tokenId) external {
        _transferFrom(_from, _to, _tokenId);
    }

    function safeTransferFrom(address _from, address _to, uint256 _tokenId) public {
        _transferFrom(_from, _to, _tokenId);
        if (_to.isContract()) {
            bytes4 retval = IERC721(_to).onERC721Received(msg.sender, _from, _tokenId, "");
            require(retval == ERC721_RECEIVED);
        }
    }

    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes memory _data) public {
        _transferFrom(_from, _to, _tokenId);
        if (_to.isContract()) {
            bytes4 retval = IERC721(_to).onERC721Received(msg.sender, _from, _tokenId, _data);
            require(retval == ERC721_RECEIVED);
        }
    }

    ////////////////////////////////////////////////////////
    // ERC998ERC721 and ERC998ERC721Enumerable implementation
    ////////////////////////////////////////////////////////

    // tokenId => child contract
    mapping(uint256 => address[]) private childContracts;

    // tokenId => (child address => contract index+1)
    mapping(uint256 => mapping(address => uint256)) private childContractIndex;

    // tokenId => (child address => array of child tokens)
    mapping(uint256 => mapping(address => uint256[])) private childTokens;

    // tokenId => (child address => (child token => child index+1)
    mapping(uint256 => mapping(address => mapping(uint256 => uint256))) private childTokenIndex;

    // child address => childId => tokenId
    mapping(address => mapping(uint256 => uint256)) internal childTokenOwner;


    function removeChild(uint256 _tokenId, address _childContract, uint256 _childTokenId) private {
        uint256 tokenIndex = childTokenIndex[_tokenId][_childContract][_childTokenId];
        require(tokenIndex != 0, "Child token not owned by token.");

        // remove child token
        uint256 lastTokenIndex = childTokens[_tokenId][_childContract].length - 1;
        uint256 lastToken = childTokens[_tokenId][_childContract][lastTokenIndex];
        if (_childTokenId == lastToken) {
            childTokens[_tokenId][_childContract][tokenIndex - 1] = lastToken;
            childTokenIndex[_tokenId][_childContract][lastToken] = tokenIndex;
        }
        childTokens[_tokenId][_childContract].length--;
        delete childTokenIndex[_tokenId][_childContract][_childTokenId];
        delete childTokenOwner[_childContract][_childTokenId];

        // remove contract
        if (lastTokenIndex == 0) {
            uint256 lastContractIndex = childContracts[_tokenId].length - 1;
            address lastContract = childContracts[_tokenId][lastContractIndex];
            if (_childContract != lastContract) {
                uint256 contractIndex = childContractIndex[_tokenId][_childContract];
                childContracts[_tokenId][contractIndex] = lastContract;
                childContractIndex[_tokenId][lastContract] = contractIndex;
            }
            childContracts[_tokenId].length--;
            delete childContractIndex[_tokenId][_childContract];
        }
    }

    function safeTransferChild(uint256 _fromTokenId, address _to, address _childContract, uint256 _childTokenId) public {
        uint256 tokenId = childTokenOwner[_childContract][_childTokenId];
        require(tokenId > 0 || childTokenIndex[tokenId][_childContract][_childTokenId] > 0);
        require(tokenId == _fromTokenId);
        require(_to != address(0));
        address rootOwner = address(bytes20(rootOwnerOf(tokenId)));
        require(rootOwner == msg.sender || _operatorApprovals[rootOwner][msg.sender] ||
        _tokenApprovals[rootOwner][tokenId] == msg.sender);
        removeChild(tokenId, _childContract, _childTokenId);
        IERC721(_childContract).safeTransferFrom(address(this), _to, _childTokenId);
        emit TransferChild(tokenId, _to, _childContract, _childTokenId);
    }

    function safeTransferChild(uint256 _fromTokenId, address _to, address _childContract, uint256 _childTokenId, bytes memory _data) public {
        uint256 tokenId = childTokenOwner[_childContract][_childTokenId];
        require(tokenId > 0 || childTokenIndex[tokenId][_childContract][_childTokenId] > 0);
        require(tokenId == _fromTokenId);
        require(_to != address(0));
        address rootOwner = address(bytes20(rootOwnerOf(tokenId)));
        require(rootOwner == msg.sender || _operatorApprovals[rootOwner][msg.sender] ||
        _tokenApprovals[rootOwner][tokenId] == msg.sender);
        removeChild(tokenId, _childContract, _childTokenId);
        IERC721(_childContract).safeTransferFrom(address(this), _to, _childTokenId, _data);
        emit TransferChild(tokenId, _to, _childContract, _childTokenId);
    }

    function transferChild(uint256 _fromTokenId, address _to, address _childContract, uint256 _childTokenId) external {
        uint256 tokenId = childTokenOwner[_childContract][_childTokenId];
        require(tokenId > 0 || childTokenIndex[tokenId][_childContract][_childTokenId] > 0);
        require(tokenId == _fromTokenId);
        require(_to != address(0));
        address rootOwner = address(bytes20(rootOwnerOf(tokenId)));
        require(rootOwner == msg.sender || _operatorApprovals[rootOwner][msg.sender] ||
        _tokenApprovals[rootOwner][tokenId] == msg.sender);
        removeChild(tokenId, _childContract, _childTokenId);
        //this is here to be compatible with cryptokitties and other old contracts that require being owner and approved
        // before transferring.
        //does not work with current standard which does not allow approving self, so we must let it fail in that case.
        //0x095ea7b3 == "approve(address,uint256)"
        bytes memory _calldata = abi.encodeWithSelector(0x095ea7b3, this, _childTokenId);
        assembly {
            let success := call(gas, _childContract, 0, add(_calldata, 0x20), mload(_calldata), _calldata, 0)
        }
        IERC721(_childContract).transferFrom(address(this), _to, _childTokenId);
        emit TransferChild(tokenId, _to, _childContract, _childTokenId);
    }

    // this contract has to be approved first in _childContract
    function getChild(address _from, uint256 _tokenId, address _childContract, uint256 _childTokenId) external {
        receiveChild(_from, _tokenId, _childContract, _childTokenId);
        require(_from == msg.sender ||
        IERC721(_childContract).isApprovedForAll(_from, msg.sender) ||
        IERC721(_childContract).getApproved(_childTokenId) == msg.sender);
        IERC721(_childContract).transferFrom(_from, address(this), _childTokenId);
    }

    function onERC721Received(address _from, uint256 _childTokenId, bytes memory _data) public returns (bytes4) {
        require(_data.length > 0, "_data must contain the uint256 tokenId to transfer the child token to.");
        // convert up to 32 bytes of_data to uint256, owner nft tokenId passed as uint in bytes
        uint256 tokenId;
        assembly {tokenId := calldataload(132)}
        if (_data.length < 32) {
            tokenId = tokenId >> 256 - _data.length * 8;
        }
        receiveChild(_from, tokenId, msg.sender, _childTokenId);
        require(IERC721(msg.sender).ownerOf(_childTokenId) != address(0), "Child token not owned.");
        return ERC721_RECEIVED;
    }


    function onERC721Received(address _operator, address _from, uint256 _childTokenId, bytes memory _data) public returns (bytes4) {
        require(_data.length > 0, "_data must contain the uint256 tokenId to transfer the child token to.");
        require(_operator!=address(0));
        // convert up to 32 bytes of_data to uint256, owner nft tokenId passed as uint in bytes
        uint256 tokenId;
        assembly {tokenId := calldataload(164)}
        if (_data.length < 32) {
            tokenId = tokenId >> 256 - _data.length * 8;
        }

        receiveChild(_from, tokenId, msg.sender, _childTokenId);
        require(IERC721(msg.sender).ownerOf(_childTokenId) != address(0), "Child token not owned.");
        return ERC721_RECEIVED;
    }


    function receiveChild(address _from, uint256 _tokenId, address _childContract, uint256 _childTokenId) private {
        require(_tokenOwner[_tokenId] != address(0), "_tokenId does not exist.");
        require(childTokenIndex[_tokenId][_childContract][_childTokenId] == 0, "Cannot receive child token because it has already been received.");
        uint256 childTokensLength = childTokens[_tokenId][_childContract].length;
        if (childTokensLength == 0) {
            childContractIndex[_tokenId][_childContract] = childContracts[_tokenId].length;
            childContracts[_tokenId].push(_childContract);
        }
        childTokens[_tokenId][_childContract].push(_childTokenId);
        childTokenIndex[_tokenId][_childContract][_childTokenId] = childTokensLength + 1;
        childTokenOwner[_childContract][_childTokenId] = _tokenId;
        emit ReceivedChild(_from, _tokenId, _childContract, _childTokenId);
    }

    function _ownerOfChild(address _childContract, uint256 _childTokenId) internal view returns (address parentTokenOwner, uint256 parentTokenId) {
        parentTokenId = childTokenOwner[_childContract][_childTokenId];
        require(parentTokenId > 0 || childTokenIndex[parentTokenId][_childContract][_childTokenId] > 0);
        return (_tokenOwner[parentTokenId], parentTokenId);
    }

    function ownerOfChild(address _childContract, uint256 _childTokenId) external view returns (bytes32 parentTokenOwner, uint256 parentTokenId) {
        parentTokenId = childTokenOwner[_childContract][_childTokenId];
        require(parentTokenId > 0 || childTokenIndex[parentTokenId][_childContract][_childTokenId] > 0);
        return (bytes32(ERC998_MAGIC_VALUE) << 224 | bytes32(bytes20(_tokenOwner[parentTokenId])), parentTokenId);
    }

    function childExists(address _childContract, uint256 _childTokenId) external view returns (bool) {
        uint256 tokenId = childTokenOwner[_childContract][_childTokenId];
        return childTokenIndex[tokenId][_childContract][_childTokenId] != 0;
    }

    function totalChildContracts(uint256 _tokenId) external view returns (uint256) {
        return childContracts[_tokenId].length;
    }

    function childContractByIndex(uint256 _tokenId, uint256 _index) external view returns (address childContract) {
        require(_index < childContracts[_tokenId].length, "Contract address does not exist for this token and index.");
        return childContracts[_tokenId][_index];
    }

    function totalChildTokens(uint256 _tokenId, address _childContract) external view returns (uint256) {
        return childTokens[_tokenId][_childContract].length;
    }

    function childTokenByIndex(uint256 _tokenId, address _childContract, uint256 _index) external view returns (uint256 childTokenId) {
        require(_index < childTokens[_tokenId][_childContract].length, "Token does not own a child token at contract address and index.");
        return childTokens[_tokenId][_childContract][_index];
    }

    ////////////////////////////////////////////////////////
    // ERC998ERC20 implementation
    ////////////////////////////////////////////////////////

    // tokenId => token contract
    mapping(uint256 => address[]) erc20Contracts;

    // tokenId => (token contract => token contract index)
    mapping(uint256 => mapping(address => uint256)) erc20ContractIndex;

    // tokenId => (token contract => balance)
    mapping(uint256 => mapping(address => uint256)) erc20Balances;

    function balanceOfERC20(uint256 _tokenId, address _erc20Contract) external view returns (uint256) {
        return erc20Balances[_tokenId][_erc20Contract];
    }

    function removeERC20(uint256 _tokenId, address _erc20Contract, uint256 _value) private {
        if (_value == 0) {
            return;
        }
        uint256 erc20Balance = erc20Balances[_tokenId][_erc20Contract];
        require(erc20Balance >= _value, "Not enough token available to transfer.");
        uint256 newERC20Balance = erc20Balance - _value;
        erc20Balances[_tokenId][_erc20Contract] = newERC20Balance;
        if (newERC20Balance == 0) {
            uint256 lastContractIndex = erc20Contracts[_tokenId].length - 1;
            address lastContract = erc20Contracts[_tokenId][lastContractIndex];
            if (_erc20Contract != lastContract) {
                uint256 contractIndex = erc20ContractIndex[_tokenId][_erc20Contract];
                erc20Contracts[_tokenId][contractIndex] = lastContract;
                erc20ContractIndex[_tokenId][lastContract] = contractIndex;
            }
            erc20Contracts[_tokenId].length--;
            delete erc20ContractIndex[_tokenId][_erc20Contract];
        }
    }


    function transferERC20(uint256 _tokenId, address _to, address _erc20Contract, uint256 _value) external {
        require(_to != address(0));
        address rootOwner = address(bytes20(rootOwnerOf(_tokenId)));
        require(rootOwner == msg.sender || _operatorApprovals[rootOwner][msg.sender] ||
        _tokenApprovals[rootOwner][_tokenId] == msg.sender);
        removeERC20(_tokenId, _erc20Contract, _value);
        require(IERC20(_erc20Contract).transfer(_to, _value), "ERC20 transfer failed.");
        emit TransferERC20(_tokenId, _to, _erc20Contract, _value);
    }

    // this contract has to be approved first by _erc20Contract
    function getERC20(address _from, uint256 _tokenId, address _erc20Contract, uint256 _value) public {
        bool allowed = _from == msg.sender;
        if (!allowed) {
            uint256 remaining;
            // 0xdd62ed3e == allowance(address,address)
            bytes memory _calldata = abi.encodeWithSelector(0xdd62ed3e, _from, msg.sender);
            bool callSuccess;
            assembly {
                callSuccess := staticcall(gas, _erc20Contract, add(_calldata, 0x20), mload(_calldata), _calldata, 0x20)
                if callSuccess {
                    remaining := mload(_calldata)
                }
            }
            require(callSuccess, "call to allowance failed");
            require(remaining >= _value, "Value greater than remaining");
            allowed = true;
        }
        require(allowed, "not allowed to getERC20");
        erc20Received(_from, _tokenId, _erc20Contract, _value);
        require(IERC20(_erc20Contract).transferFrom(_from, address(this), _value), "ERC20 transfer failed.");
    }

    function erc20Received(address _from, uint256 _tokenId, address _erc20Contract, uint256 _value) private {
        require(_tokenOwner[_tokenId] != address(0), "_tokenId does not exist.");
        if (_value == 0) {
            return;
        }
        uint256 erc20Balance = erc20Balances[_tokenId][_erc20Contract];
        if (erc20Balance == 0) {
            erc20ContractIndex[_tokenId][_erc20Contract] = erc20Contracts[_tokenId].length;
            erc20Contracts[_tokenId].push(_erc20Contract);
        }
        erc20Balances[_tokenId][_erc20Contract] += _value;
        emit ReceivedERC20(_from, _tokenId, _erc20Contract, _value);
    }

    function erc20ContractByIndex(uint256 _tokenId, uint256 _index) external view returns (address) {
        require(_index < erc20Contracts[_tokenId].length, "Contract address does not exist for this token and index.");
        return erc20Contracts[_tokenId][_index];
    }

    function totalERC20Contracts(uint256 _tokenId) external view returns (uint256) {
        return erc20Contracts[_tokenId].length;
    }

}