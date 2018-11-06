(
cd ../etc/contracts 

solc --abi privateBridgeToken.sol -o ./privateAbi --overwrite
solc --bin privateBridgeToken.sol -o ./privateBin --overwrite

) && abigen --abi=../etc/contracts/privateAbi/GameToken.abi --bin=../etc/contracts/privateBin/GameToken.bin --pkg=privateSlot -out ../token/privateSlot/privateSlot.go
