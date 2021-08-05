# kryptos_rai v0.1 doc



## 程序结构

kryptos_rai

├── address_analysis  //地址余额查询
├── block_parse  //区块解析
│   ├── simple_internal_tx  //简易内部交易解析
│   └── simple_tx //简易交易解析
├── cmd
│   ├── address_analysis
│   ├── block_parse
│   │   ├── simple_internal_tx
│   │   └── simple_tx
│   ├── dispatching_center
│   ├── notify
│   ├── pool_watch
│   └── web
├── dispatching_center  //调度中心
├── notify //消息通知
├── pool_watch //池子监控
├── support_pools //支持的项目(池子)
├── use //公共
└── web //web



详见各个文件夹下面的readme.txt

## 需求/v0.1处理

### 需求

1 池子监控 动态

2 支持的项目(协议)  动态

3 地址分析前端界面返回小于3s



### v0.1处理

1 池子监控与消息通知pool_watch/notify 和 地址余额查询address_analysis没关联 各自独立运行

2 不同项目池子监控与余额查询处理方式不同  无法程序化自动增加 具体来说比如各个合约abi不同 需要手动放到程序中去 搞了调度中心dispatching_center曲线救国实现热更新那种方式 这种方式有一个好处就是会增加地址查询成功率 因为节点的不稳定性  不同机器去查比较好

3 支持的项目(池子)信息  和 链上交易数据  各自存库。对于输入地址  根据链上交易数据 过滤得到  我们支持的而且交互过的池子列表  最后协程查询余额



### 一些信息

1 存储在mongo中(每日插入少 查询多)

2 bsc 整条链  按照字符串直接存储   交易信息200G(8亿条数据) 估记internal_tx(BEP20/ERC20)大概是 200G  如果存储字段优化为byte会省硬盘

3 按照eth出块速度与单区块交易数据大小 如果只存过去2年的ETH链上数据 大约是bsc的一半数据量200G 

4 按照bsc和eth的出块数据 以2年记的话大约是2T应该是够了的 保险点上个3T

5  mongo查询时间(根据from为输入地址查所有交易列表 )是 单表1.9亿条数据是20ms  单表8亿条数据查询是300ms 后续跟节点交互都是协程处理的话大概率是3s内返回

6 数据存储的话 多进程分区块段去获取 最后整合在一起完整的链上数据 在本地mac上大概是需要一周 如果是单进程去搞的话要好几个月才能存储完





## 本地运行与部署

### 本地运行

1. mongo运行 运行mongo.txt的建库表语句
2. 进入项目目录(kryptos_rai)
3. 修改cmd目录下(7个目录的local.sh中的gitDir变量)sh start.sh



### 部署

交叉编译 把可执行文件与配置文件放到服务器上运行







## 后续优化

1. 库的优化 可以考虑TIDB  毕竟随着数据的增加查询时间也会增加 可以试下TIDB 可以尝试直接用TIDB存储
2. 程序设计优化 --- 各位大佬如果有好的处理方式 或者发现bug 或者api风格不对  请随时告诉我我改一下



## 一些地址

API文档地址:https://www.showdoc.com.cn/1480342907406885/7178061966851686

gitee地址:https://gitee.com/purplecity/kryptos_rai



