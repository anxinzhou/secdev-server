trap " kill -9 0 " SIGINT
if [ ! -e "node/geth" ]
then 
	geth --datadir node init genesis.json
fi

geth --gasprice 0 --txpool.pricelimit 0  --mine --datadir node --syncmode 'full' --port 30321 --rpc --rpcaddr 'localhost' \
 	--rpcport 8640 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --rpccorsdomain '*'  \
 	--ws --wsport 8550 --wsapi="personal,db,eth,net,web3,txpool,miner" --wsorigins "*" --targetgaslimit '9000000000000'  --networkid 1515  \
  	--etherbase '0x3c62aa7913bc303ee4b9c07df87b556b6770e3fc' --unlock '0x3c62aa7913bc303ee4b9c07df87b556b6770e3fc' --password password.txt  
