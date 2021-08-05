入口为rpc_server 根据链名 每个链名的不同的项目的地址分析协程去处理

v0.1版本地址分析 增加链和项目 参考(仅参考):

0 目前只有bsc链pancake项目

1 如果增加链的支持
    address_analysis目录下新建 链名_analysis文件夹 比如eth_analysis
    加个eth.go switch case处理eth链上不同项目的逻辑
    rpc_server switch分支加个链分支处理
2 如果增加项目
    在链名_analysis文件夹加一个项目名_analysis的go文件 比如增加bsc链上venus的支持 就在bsc_analysis下面加一个venus_analysis.go文件 
    在bsc.go swith分支中加一个venus项目分支处理


