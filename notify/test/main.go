package main

import (
	"context"
	"fmt"
	"kryptos_rai/notify/message_notify"
	"time"

	"google.golang.org/grpc"
)

func main() {
	dataAddr := "127.0.0.1:16891"
	ds_rpc, err := grpc.Dial(dataAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Duration(10)*time.Second))
	if err != nil {
		fmt.Printf("create lg rpc client failed %v", dataAddr)

	}
	ds_client := message_notify.NewMessageNotifyClient(ds_rpc)

	freq := new(message_notify.FEN_Request)
	freq.PoolAddress = "0x11111111"
	freq.Hash = "0x22222"
	r, err := ds_client.FlashEventNotify(context.Background(), freq)
	if err != nil {
		fmt.Printf("error freq:%v", err)
	}
	fmt.Printf("freq return:%+v", r)

	preq := new(message_notify.PN_Request)
	preq.PoolAddress = "0x11111111"
	preq.CurCheckBlockHeight = "11111"
	preq.LastCheckBlockHeight = "2222"
	preq.CurCheckBlockHeightTvl = "177793756426221604490886"
	preq.LastCheckBlockHeightTvl = "177777076241166242585031"
	sr, err := ds_client.PoolNotify(context.Background(), preq)
	if err != nil {
		fmt.Printf("error preq:%+v", err)
	}
	fmt.Printf("preq return:%+v", sr)

}
