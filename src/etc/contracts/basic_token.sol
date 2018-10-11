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
