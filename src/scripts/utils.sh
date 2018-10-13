#!/bin/bash
if [[ $1 == "deploy" ]] 
then 
	if [[ $2 == "public" ]]
	then 
		node contractUtils deploy "../etc/contracts/all/public_slot.sol" PublicSlot -r localhost:8640 -k "../etc/keystore/zx" -p "321" -a [10000,1,["0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f"]]
	elif  [[ $2 == "private" ]]
	then 
		node contractUtils deploy "../etc/contracts/all/private_slot.sol"  PrivateSlot -r http://127.0.0.1:8540 -k "../etc/keystore/zx" -p "321" -a '[10000,1,["0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f"]]'

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
