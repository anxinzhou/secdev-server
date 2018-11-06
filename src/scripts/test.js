const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:4000');
var Web3 = require('web3');
web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8540"))


//done
var signinReq = {
    gcuid: 888,
    act: "login",
    login_code: "6469D2959C1AA0D2AA9A6ACDF35BEA32524B5142684948576C4B31586D317475695364587A425552334A34"
}

//done
var signoutReq = {
    gcuid: 100,
    act: "sign_out",
    guuid: "gu9"
}

//done
var walletAndMachineReq = {
    gcuid: 101,
    act: "get_wallets_and_machine_status"
}

//done
var walletReq = {
    gcuid:102,
    act: "get_wallets"
}

var ethRecordReq = {
    gcuid: 103,
    amount: 10,
    after: 5,
    act: "get_transactions",
    from: 1
}

var tokenRecordReq = {
    gcuid: 103,
    amount: 10,
    after: 5,
    act: "get_transactions",
    from: 2
}

//done
var exchangeRateReq = {
    gcuid: 104,
    exchange_type : 1,
    act: "get_exchange_rate"
}

//done
var ethToSlotReq= {
    gcuid: 105,
    exchange_type: 1,
    amount: "1.21",
    act: "post_exchange"
}

//done
var SlotToEthReq = {
    gcuid: 105,
    exchange_type: 2,
    amount: "2.5",
    act: "post_exchange"
}

//done
var qrCodeReq = {
    gcuid: 106,
    qrcode: "0x1",
    act: "post_qrcode"
}

//done
var logoutReq = {
    gcuid: 109,
    act: "machine_logout",
    machine_id: "0x11"
}

//done
var tokenReward = {
    gcuid: 111,
    type: 2,
    amount: "100",
    act: "post_token_exchange"
}

//done
var tokenUse = {
    gcuid: 111,
    type: 1,
    amount: "1.9",
    act: "post_token_exchange"
}

ws.on('open', function open() {
    // ws.send(JSON.stringify(qrCodeReq))
    // ws.send(JSON.stringify(logoutReq))
    // ws.send(JSON.stringify(walletAndMachineReq))
    // //ws.send(JSON.stringify(tokenUse))
    // ws.send(JSON.stringify(tokenReward))
    // ws.send(JSON.stringify(exchangeRateReq))
    // ws.send(JSON.stringify(SlotToEthReq))
    //
    // ws.send(JSON.stringify(ethToSlotReq))
    ws.send(JSON.stringify(ethRecordReq))
    ws.send(JSON.stringify(tokenRecordReq))
});

ws.on('message', function incoming(data) {
    console.log("")
    console.log(data)
});