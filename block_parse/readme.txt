
simple_tx 与 simple_internal_tx 是为了更新最新的区块数据(普通交易与内部交易) 因为节点的不稳定问题所以没用ws链接节点 
定时拉取 目前bsc是4秒一次


对于历史数据 多进程分区块段去拉取最后整合所有数据在一起就好了

v0.1 增加链的参考(仅参考)

比如要存eth链的数据
在simple_internal_tx和simple_tx都加一个eth文件夹 写逻辑开始拉取最新数据