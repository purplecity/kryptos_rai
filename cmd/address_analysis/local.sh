#! /bin/sh

#local
gitDir=$HOME/code
BaseDir=$gitDir/kryptos_rai/cmd/address_analysis
TargetDir=$HOME/test/kryptos_rai/address_analysis
ConfigPath=$HOME/test/kryptos_rai/address_analysis/
logfile=$HOME/test/kryptos_rai/address_analysis/address_analysis.log

ps -ef | grep "./kryptos_rai_address_analysis" | grep -v grep | awk '{print $2}' | xargs kill -9

find $TargetDir/ -name "kryptos_rai_address_analysis" | xargs rm -rf

go build $BaseDir/main.go
#$(CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $BaseDir/main.go)
mv main $TargetDir/kryptos_rai_address_analysis

nohup $TargetDir/kryptos_rai_address_analysis --config_path=$ConfigPath --log_file=$logfile >>$TargetDir/address_analysis_nohup.log 2>&1 &
