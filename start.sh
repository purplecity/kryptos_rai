#! /bin/sh

#前提是得进入kcryptos_debank
testWebDir=$HOME/test/kryptos_rai/web
testAADir=$HOME/test/kryptos_rai/address_analysis
testDCDir=$HOME/test/kryptos_rai/dispatching_center
testPWDir=$HOME/test/kryptos_rai/pool_watch
testNotiDir=$HOME/test/kryptos_rai/notify
testSTDir=$HOME/test/kryptos_rai/simple_tx/bsc
testSITDir=$HOME/test/kryptos_rai/simple_internal_tx/bsc

mkdir -p $testWebDir $testAADir $testDCDir $testPWDir $testNotiDir $testSTDir $testSITDir

cp ./block_parse/simple_tx/bsc/config/config.toml $testSTDir/
cp ./block_parse/simple_internal_tx/bsc/config/config.toml $testSITDir/
cp ./notify/config/config.toml $testNotiDir/
cp ./pool_watch/config/config.toml $testPWDir/
cp ./address_analysis/config/config.toml $testAADir/
cp ./dispatching_center/config/config.toml $testDCDir/
cp ./web/config/config.toml $testWebDir/

# ps -ef | grep "./kryptos_rai_address_analysis" | grep -v grep | awk '{print $2}' | xargs kill -9
# ps -ef | grep "./kryptos_rai_simple_tx_bsc" | grep -v grep | awk '{print $2}' | xargs kill -9
# ps -ef | grep "./kryptos_rai_simple_internal_tx_bsc" | grep -v grep | awk '{print $2}' | xargs kill -9
# ps -ef | grep "./kryptos_rai_dispatching_center" | grep -v grep | awk '{print $2}' | xargs kill -9
# ps -ef | grep "./kryptos_rai_pool_watch" | grep -v grep | awk '{print $2}' | xargs kill -9
# ps -ef | grep "./kryptos_rai_web" | grep -v grep | awk '{print $2}' | xargs kill -9
# ps -ef | grep "./kryptos_rai_notify" | grep -v grep | awk '{print $2}' | xargs kill -9

#启动顺序
sh ./cmd/block_parse/simple_tx/bsc/local.sh
sh ./cmd/block_parse/simple_internal_tx/bsc/local.sh

sh ./cmd/pool_watch/local.sh
sh ./cmd/notify/local.sh

sh ./cmd/address_analysis/local.sh
sh ./cmd/dispatching_center/local.sh
sh ./cmd/web/local.sh

s=$(ps -ef | grep kryptos | wc -l)
echo $s
if [ $s != 8 ]; then
    echo "start failed check which one failed"
fi
