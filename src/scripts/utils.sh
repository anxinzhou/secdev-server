#!/bin/bash
if [[ $1 == "deploy" ]] 
then 
	if [[ $2 == "public" ]]
	then 
		node contractUtils deploy "../etc/contracts/all/publicBridgeToken.sol" BridgeToken -r http://127.0.0.1:8640 -k "../etc/keystore/zx" -p "321" -a '[100000,"bridgeToken","btk",0,1]'
	elif  [[ $2 == "private" ]]
	then 
		node contractUtils deploy "../etc/contracts/all/privateBridgeToken.sol"  BridgeToken -r http://127.0.0.1:8540 -k "../etc/keystore/zx" -p "321" -a '[100000,"bridgeToken","btk",0,1]'

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


#0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f