package main

import (
	"log"
	"time"

	"github.com/vsc-eco/hivego"
)

func main() {
	hrpc := hivego.NewHiveClient(1, "https://api.hive.blog")

	props, _ := hrpc.GetGlobalProps()
	log.Println(props.HeadBlockNumber)
	r, _ := hrpc.GetBlock(props.HeadBlockNumber)
	log.Println("block=", r.BlockNumber)
	go hrpc.StreamBlocks()
	for {
		log.Println("blocks streaming...")
		time.Sleep(1 * time.Minute)

	}
}
