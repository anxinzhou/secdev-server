#!/bin/bash
if [[ $1 == "deploy" ]]
then
	if [[ $2 == "public" ]]
	then
		node contractUtils deploy "../etc/contracts/publicBridgeToken.sol" GameToken -r http://127.0.0.1:8640 -k "../etc/chain/private/testnet/node/keystore/authority1" -p "321" -a '[100000,"GameToken","gtk"]'
	elif  [[ $2 == "private" ]]
	then
		node contractUtils deploy "../etc/contracts/privateBridgeToken.sol"  GameToken -r http://127.0.0.1:8540 -k "../etc/chain/private/testnet/node/keystore/authority1" -p "321" -a '[100000,"GameToken","gtk"]'

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