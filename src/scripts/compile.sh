(
cd ../etc/contracts 

solc --abi privateBridgeToken.sol -o ./privateAbi --overwrite
solc --abi publicBridgeToken.sol -o ./publicAbi --overwrite
solc --bin publicBridgeToken.sol -o ./publicBin --overwrite
solc --bin privateBridgeToken.sol -o ./privateBin --overwrite

)

abigen --abi=../etc/contracts/privateAbi/BridgeToken.abi --bin=../etc/contracts/privateBin/BridgeToken.bin --pkg=privateSlot -out ../token/privateSlot/privateSlot.go

abigen --abi=../etc/contracts/publicAbi/BridgeToken.abi --bin=../etc/contracts/publicBin/BridgeToken.bin --pkg=publicSlot -out ../token/publicSlot/publicSlot.go
