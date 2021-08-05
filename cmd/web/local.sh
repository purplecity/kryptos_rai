#! /bin/sh

#local
gitDir=$HOME/code
BaseDir=$gitDir/kryptos_rai/cmd/web
TargetDir=$HOME/test/kryptos_rai/web
ConfigPath=$HOME/test/kryptos_rai/web/
logfile=$HOME/test/kryptos_rai/web/web.log

ps -ef | grep "./kryptos_rai_web" | grep -v grep | awk '{print $2}' | xargs kill -9

find $TargetDir/ -name "kryptos_rai_web" | xargs rm -rf

go build $BaseDir/main.go
#$(CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $BaseDir/main.go)
mv main $TargetDir/kryptos_rai_web

nohup $TargetDir/kryptos_rai_web --config_path=$ConfigPath --log_file=$logfile >>$TargetDir/web_nohup.log 2>&1 &
