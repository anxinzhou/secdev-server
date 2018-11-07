(
cd ../etc/contracts 

solc --abi privateBridgeToken.sol -o ./privateAbi --overwrite
solc --bin privateBridgeToken.sol -o ./privateBin --overwrite

solc --abi publicBridgeToken.sol -o ./publicAbi --overwrite
solc --bin publicBridgeToken.sol -o ./publicBin --overwrite

)

abigen --abi=../etc/contracts/privateAbi/GameToken.abi --bin=../etc/contracts/privateBin/GameToken.bin --pkg=privateSlot -out ../token/privateSlot/privateSlot.go &&
abigen --abi=../etc/contracts/publicAbi/GameToken.abi --bin=../etc/contracts/publicBin/GameToken.bin --pkg=publicSlot -out ../token/publicSlot/publicSlot.go