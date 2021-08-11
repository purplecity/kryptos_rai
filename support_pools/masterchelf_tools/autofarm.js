import { ethers } from 'ethers';


let symbolAbi = [{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"}];
let url = "https://bsc-dataseed1.binance.org/";
let provider = new ethers.providers.JsonRpcProvider(url);
let pidABI= [
    {"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"poolInfo","outputs":[{"internalType":"contract IERC20","name":"want","type":"address"},{"internalType":"uint256","name":"allocPoint","type":"uint256"},{"internalType":"uint256","name":"lastRewardBlock","type":"uint256"},{"internalType":"uint256","name":"accAUTOPerShare","type":"uint256"},{"internalType":"address","name":"strat","type":"address"}],"stateMutability":"view","type":"function"}]



let pidArray =  new Array()

let consoleArray = new Array()
consoleArray.push({Address:"0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82",PID:7,Symbol:"Cake",Precision:18})
consoleArray.push({Address:"0xa8Bb71facdd46445644C277F9499Dd22f6F0A30C",PID:338,Symbol:"beltBNB",Precision:18})
consoleArray.push({Address:"0x51bd63F240fB13870550423D208452cA87c44444",PID:339,Symbol:"beltBTC",Precision:18})
consoleArray.push({Address:"0xAA20E8Cb61299df2357561C2AC2e1172bC68bc25",PID:340,Symbol:"beltETH",Precision:18})
consoleArray.push({Address:"0x603c7f932ED1fc6575303D8Fb018fDCBb0f39a95",PID:348,Symbol:"BANANA",Precision:18})
consoleArray.push({Address:"0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",PID:380,Symbol:"WBNB",Precision:18})
consoleArray.push({Address:"0x9C65AB58d8d978DB963e63f2bfB7121627e3a739",PID:415,Symbol:"MDX",Precision:18})
consoleArray.push({Address:"0xa9c41A46a6B3531d28d5c32F6633dd2fF05dFB90",PID:370,Symbol:"WEX",Precision:18})
consoleArray.push({Address:"0x0487b824c8261462F88940f97053E65bDb498446",PID:386,Symbol:"WINGS",Precision:18})
consoleArray.push({Address:"0x9cb73F20164e399958261c289Eb5F9846f4D1404",PID:341,Symbol:"4Belt",Precision:18})





let exceedList = [
    "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82",//单币质押 
    "0xa8Bb71facdd46445644C277F9499Dd22f6F0A30C",
    "0x51bd63F240fB13870550423D208452cA87c44444",
    "0xAA20E8Cb61299df2357561C2AC2e1172bC68bc25",
    "0x603c7f932ED1fc6575303D8Fb018fDCBb0f39a95", 
    "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", 
    "0x9C65AB58d8d978DB963e63f2bfB7121627e3a739",
    "0xa9c41A46a6B3531d28d5c32F6633dd2fF05dFB90",
    "0x0487b824c8261462F88940f97053E65bDb498446",
    "0x9cb73F20164e399958261c289Eb5F9846f4D1404"
]


function exceed(addr)  {
    for (v of exceedList) {
        if (String(v).toLowerCase() == String(addr).toLowerCase()) {
            return true
        }
    }
    return false
}

let masterchelfAddress = "0x0895196562C7868C5Be92459FaE7f877ED450452"
let masterchelfContract = new ethers.Contract(masterchelfAddress,pidABI,provider)

//因为有的合约没有公开 所以查询可能会失败 所以只能分批次去搞了 不用一回查到440
for (let index = 1; index < 417; index++) {
    let addressinfo = await masterchelfContract.poolInfo(index)

    let item = {Address:addressinfo.want,PID:index,Strat:addressinfo.strat}
    
    if (String(addressinfo.want).toLowerCase() == String("0x763a05bdb9f8946d8c3fa72d1e0d3f5e68647e5c").toLowerCase()) {

    }

    if (addressinfo.allocPoint >0 && !exceed(addressinfo.want)) {
        pidArray.push(item)
    }


    //没暂停 合约公开
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


for (var v of consoleArray){
    console.log("{"+ "Address="+"\""+String(v.Address)+"\""+","+"PID="+String(v.PID)+","+"Symbol="+"\""+String(v.Symbol)+"\""+","+"Precision="+v.Precision+"}" + ",")
}



