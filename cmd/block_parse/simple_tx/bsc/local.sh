#! /bin/sh

#local
gitDir=$HOME/code
BaseDir=$gitDir/kryptos_rai/cmd/block_parse/simple_tx/bsc
TargetDir=$HOME/test/kryptos_rai/simple_tx/bsc
ConfigPath=$HOME/test/kryptos_rai/simple_tx/bsc/
logfile=$HOME/test/kryptos_rai/simple_tx/bsc/simple_tx_bsc.log

ps -ef | grep "./kryptos_rai_simple_tx_bsc" | grep -v grep | awk '{print $2}' | xargs kill -9

find $TargetDir/ -name "kryptos_rai_simple_tx_bsc" | xargs rm -rf

go build $BaseDir/main.go
#$(CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $BaseDir/main.go)
mv main $TargetDir/kryptos_rai_simple_tx_bsc

nohup $TargetDir/kryptos_rai_simple_tx_bsc --config_path=$ConfigPath --log_file=$logfile >>$TargetDir/simple_tx_bsc_nohup.log 2>&1 &
