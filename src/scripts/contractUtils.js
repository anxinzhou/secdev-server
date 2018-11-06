const solc = require('solc')
const fs = require('fs-extra')
var program = require('commander')
var Web3 = require('web3');
const {
    promisify
} = require('util')
const readFile = promisify(fs.readFile)
const readJson = promisify(fs.readJson)

program
    .version('1.0.0')

program
    .command('deploy <file> <contract>')
    .description('deploy contract in  file')
    .option('-r, --rpcport <port>', 'rpc path (default:8545)')
    .option('-a, --arguments <arguments>', 'contract constructor arguments')
    .option('-k, --keystore <path>', 'keystore file to deploy contract')
    .option('-p, --password <password>', 'password to decrpt keystore')
    .action(function(file, contract, options) {
        if (!options.keystore) {
            throw new Error('no keystore file specified')
        }

        if (!options.password) {
            throw new Error('no password specified')
        }

        web3 = new Web3(new Web3.providers.HttpProvider(options.rpcport))
        readFile(file).then((input) => {
            output = solc.compile(input.toString(), 1)
            contractName = ':' + contract
            if (!output.contracts.hasOwnProperty(contractName)) {
                throw new Error('no such contract ' + contractName)
            }
            bytecode = output.contracts[contractName].bytecode
            abi = JSON.parse(output.contracts[contractName].interface)
            contract = new web3.eth.Contract(abi)
            args = JSON.parse(options.arguments)
            readJson(options.keystore).then(ks => {
                account = web3.eth.accounts.decrypt(ks, options.password)
                dp = contract.deploy({
                    data: '0x' + bytecode,
                    arguments: args
                })
                var tx = {
                    data: dp.encodeABI(),
                    from: account.address,
                    gas: "3000000",
                    gasPrice: '0',
                    value: "10000000000000000000000", // 1000 ether
                }
                return account.signTransaction(tx)
            }).then(signedTx => {
                console.log(110)
                return web3.eth.sendSignedTransaction(signedTx.rawTransaction)
            }).then(receipt => {
                if (receipt.status == '0x0') {
                    throw new Error("contract deploy revert")
                }
                console.log(receipt.contractAddress)
            }).catch(console.log)
        })
    })

program
    .command('genkey <password>')
    .description('use password to generate keystore')
    .option('-f, --file <path>', "path to store keystore")
    .action(function(password, options) {
        web3 = new Web3()
        account = web3.eth.accounts.create()
        ks = account.encrypt(password)
        console.log(account.address)
        if (options.file) {
            fs.writeJson(options.file, ks, () => console.log("write key store to", options.file))
        }
    })

program.parse(process.argv)
