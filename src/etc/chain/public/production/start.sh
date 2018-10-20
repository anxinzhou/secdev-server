trap " kill -9 0 " SIGINT
node=$1
if [ ! -e "node1/geth" ]
then 
	geth --datadir node1 init genesis.json
fi

if [ ! -e "node2/geth" ]
then 
	geth --datadir node2 init genesis.json
fi

if [ ! -e "node3/geth" ]
then 
	geth --datadir node3 init genesis.json
fi

if [ $node = "node1" ]
then
geth --gasprice 0 --txpool.pricelimit 0  --mine --datadir node1/ --syncmode 'full' --port 30321 --rpc --rpcaddr 'localhost' \
 	--rpcport 8640 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --rpccorsdomain '*'  \
 	--ws --wsport 8550 --wsapi="personal,db,eth,net,web3,txpool,miner" --wsorigins "*" --targetgaslimit '9000000000000'  --networkid 1515  \
  	--etherbase '0x3c62aa7913bc303ee4b9c07df87b556b6770e3fc' --unlock '0x3c62aa7913bc303ee4b9c07df87b556b6770e3fc' --password password.txt  --bootnodes "enode://0c9ac4595299a4ca179fac9ed5f9a8f2fc508065b931741dbb315b470c8adcfd983dbe44a99f31f6693489a3c086881dbc22eb8d16073c28e70163d64b3f0f96@13.251.200.1:30311"
elif [ $node = "node2" ]
then
geth --gasprice 0 --txpool.pricelimit 0  --mine --datadir node2/ --syncmode 'full' --port 30322 --rpc --rpcaddr 'localhost' \
 --rpcport 8641 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --rpccorsdomain '*'  \
 --ws --wsport 8551 --wsapi="personal,db,eth,net,web3,txpool,miner" --wsorigins "*" --targetgaslimit '9000000000000'  --networkid 1515  \
 --etherbase "0x752befae2efee656811eceeeea46a2d6d9621733" --unlock "0x752befae2efee656811eceeeea46a2d6d9621733" --password password.txt 
elif [ $node = "node3" ]
then
 geth --gasprice 0 --txpool.pricelimit 0 --mine --datadir node3/ --targetgaslimit '9000000000000' --syncmode 'full' --port 30323 --rpc --rpcaddr 'localhost' \
  --rpcport 8642 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --rpccorsdomain '*'  \
 --ws --wsport 8552 --wsapi="personal,db,eth,net,web3,txpool,miner" --wsorigins "*" --targetgaslimit '9000000000000'  --networkid 1515  \
 --etherbase "0x81a5c09bb2f15f4548f458b7f3b4d49080e0eb4a" --unlock '0x81a5c09bb2f15f4548f458b7f3b4d49080e0eb4a' --password password.txt 
else
	echo 'wrong target node file'
	exit
fi
