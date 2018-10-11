(
cd ../etc/contracts 

solc --abi private_slot.sol -o ./privateAbi --overwrite
solc --abi public_slot.sol -o ./publicAbi --overwrite
solc --bin public_slot.sol -o ./publicBin --overwrite
solc --bin private_slot.sol -o ./privateBin --overwrite

)

abigen --abi=../etc/contracts/privateAbi/PrivateSlot.abi --bin=../etc/contracts/privateBin/PrivateSlot.bin --pkg=privateSlot -out ../token/privateSlot/privateSlot.go

abigen --abi=../etc/contracts/publicAbi/PublicSlot.abi --bin=../etc/contracts/publicBin/PublicSlot.bin --pkg=publicSlot -out ../token/publicSlot/publicSlot.go
