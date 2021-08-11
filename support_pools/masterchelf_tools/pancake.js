import { ethers } from 'ethers';


let symbolAbi = [{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"}];
let url = "https://bsc-dataseed1.binance.org/";
let provider = new ethers.providers.JsonRpcProvider(url);
let pidABI= [
    {"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"poolInfo","outputs":[{"internalType":"contract IBEP20","name":"lpToken","type":"address"},{"internalType":"uint256","name":"allocPoint","type":"uint256"},{"internalType":"uint256","name":"lastRewardBlock","type":"uint256"},{"internalType":"uint256","name":"accCakePerShare","type":"uint256"}],"stateMutability":"view","type":"function"}
]


/*
    "https://bsc-dataseed1.binance.org/",
    "https://bsc-dataseed2.binance.org/",
    "https://bsc-dataseed3.binance.org/",
    "https://bsc-dataseed4.binance.org/",
    "https://bsc-dataseed1.defibit.io/",
    "https://bsc-dataseed2.defibit.io/",
    "https://bsc-dataseed3.defibit.io/",
    "https://bsc-dataseed4.defibit.io/",
    "https://bsc-dataseed1.ninicoin.io/",
    "https://bsc-dataseed2.ninicoin.io/",
    "https://bsc-dataseed3.ninicoin.io/",
    "https://bsc-dataseed4.ninicoin.io/"
*/
//获取所有pid以及对应的lp地址 
//每个lp地址获取组成token的地址 每个地址获取symbol

let pidArray =  new Array()
let consoleArray = new Array()
consoleArray.push({Address:"0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82",PID:0,Symbol:"Cake",Precision:18})
consoleArray.push({Address:"0xE0388F5Fe45EE31d1838CBA963B92d49E9879fD9",PID:33,Symbol:"DPT",Precision:18})
consoleArray.push({Address:"0x2222227E22102Fe3322098e4CBfE18cFebD57c95",PID:126,Symbol:"TLM",Precision:18})
consoleArray.push({Address:"0xcA8825E9486ac20b17737E7231d67e1373B7FB4a",PID:137,Symbol:"LOTTO",Precision:18})
consoleArray.push({Address:"0xffaf3aF0Cb86BA126158CAb09DbD094996Bb4d24",PID:138,Symbol:"BURN",Precision:18})

let exceedList = [
    "0xffaf3aF0Cb86BA126158CAb09DbD094996Bb4d24",//单币质押 138
    "0xcA8825E9486ac20b17737E7231d67e1373B7FB4a",//单币质押 137 LOTTO
    "0x2222227E22102Fe3322098e4CBfE18cFebD57c95",//单币质押TLM  126
    "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82",//单币质押 CAKE
    "0xE0388F5Fe45EE31d1838CBA963B92d49E9879fD9", //单币质押 DPT
    "0xa96b5ac38b33af0fef7d40c200f2bac9d817ebca", //合约没开源 pid31
    "0x6aaD6e2ec5B3B12F1853518534Cc9C2e5b2177ae",//pid 63 token1没有symbol方法 是一个pancake的factory
    "0xBB2281cA5B4d07263112604D1F182AD0Ab26a252",//pid 105 只是一个地址不是一个合约
]


function exceed(addr)  {
    for (v of exceedList) {
        if (String(v).toLowerCase() == String(addr).toLowerCase()) {
            return true
        }
    }
    return false
}


let masterchelfAddress = "0x73feaa1eE314F8c655E354234017bE2193C9E24E"
let masterchelfContract = new ethers.Contract(masterchelfAddress,pidABI,provider)

//因为有的合约没有公开 所以查询可能会失败 所以只能分批次去搞了 不用一回查到440
for (let index = 370; index < 441; index++) {
    console.log(index)
    let addressinfo = await masterchelfContract.poolInfo(index)

    let item = {Address:addressinfo.lpToken,PID:index}

    if (!exceed(addressinfo.lpToken) && addressinfo.allocPoint>0) {
        pidArray.push(item)
    }
}


let tokenABI = [
    {"constant":true,"inputs":[],"name":"token0","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"token1","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"}
]
for (var v of pidArray){
    console.log(v.Address,v.PID)
    let item = {Address:String(v.Address),PID:v.PID}

    let pairContract = new ethers.Contract(String(v.Address), tokenABI, provider);
    console.log("pairContract",v.Address)
    let token0 =  await pairContract.token0();
    console.log("token0",v.Address,token0)
    let token1 =  await  pairContract.token1();
    console.log("token1",v.Address,token1)



    let symbol0Contract = new ethers.Contract(String(token0), symbolAbi, provider);
    let symbol0 =  await symbol0Contract.symbol();
    console.log("symbol0",v.Address,symbol0)

    let symbol1Contract = new ethers.Contract(String(token1), symbolAbi, provider);
    let symbol1 =  await symbol1Contract.symbol();
    console.log("symbol1",v.Address,symbol1)

    let lpSymbol = String(symbol0) + "-"+String(symbol1) + " LP"
    item.Symbol = lpSymbol
    //console.log(item)

    let decimals = await pairContract.decimals()
    console.log("decimals",v.Address,decimals)
    item.Precision = decimals
    consoleArray.push(item)
}

//console.log(consoleArray)



//   { Address = "0x22", PID = 1, Symbol = "hehe" },
for (var v of consoleArray){
    console.log("{"+ "Address="+"\""+String(v.Address)+"\""+","+"PID="+String(v.PID)+","+"Symbol="+"\""+String(v.Symbol)+"\""+","+"Precision="+v.Precision+"}" + ",")
}

//npm init
//node exec.js >>./result.txt