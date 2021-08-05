#! /bin/sh

#local
gitDir=$HOME/code
BaseDir=$gitDir/kryptos_rai/cmd/notify
TargetDir=$HOME/test/kryptos_rai/notify
ConfigPath=$HOME/test/kryptos_rai/notify/
logfile=$HOME/test/kryptos_rai/notify/notify.log

ps -ef | grep "./kryptos_rai_notify" | grep -v grep | awk '{print $2}' | xargs kill -9

find $TargetDir/ -name "kryptos_rai_notify" | xargs rm -rf

go build $BaseDir/main.go
#$(CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $BaseDir/main.go)
mv main $TargetDir/kryptos_rai_notify

nohup $TargetDir/kryptos_rai_notify --config_path=$ConfigPath --log_file=$logfile >>$TargetDir/notify_nohup.log 2>&1 &
