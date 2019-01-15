#!/bin/bash
if [[ $1 == "deploy" ]]
then
	if [[ $2 == "public" ]]
	then
		node contractUtils deploy "../etc/contracts/PbcERC998.sol" ComposableTopDown -r "https://kovan.infura.io/v3/38c50b5e76aa4de9866c79935349bee8" -k "../etc/chain/private/testnet/node/keystore/authority1" -p "321" -a '["testSlot","ts"]' -s
	elif  [[ $2 == "private" ]]
	then
		node contractUtils deploy "../etc/contracts/PvcERC998.sol" ComposableTopDown -r "http://127.0.0.1:8540" -k "../etc/chain/private/testnet/node/keystore/authority1" -p "321" -a '["testSlot","ts"]' -s
	else
		echo $2
		echo 'deploy public or private'
	fi
elif [[ $1 == "genkey" ]]
then
	node contractUtils.js genkey 321 -f "../etc/keystore/zx"
else
	echo 'genkey or deploy'
fi


# address of keystore
#0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f


# address of owner
# "0x3c62aa7913bc303ee4b9c07df87b556b6770e3fc"
