package main

import (
	"context"
	"cosmossdk.io/math"
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

func main() {
	addr := os.Args[1]
	token := os.Args[2]

	fillBlocks(token, addr)
}

func fillBlocks(token, addr string) {
	client, err := client.NewClient(context.Background(), addr, token)
	if err != nil {
		panic(err)
	}

	for {
		nID := make([]byte, 8)
		rand.Read(nID)
		data := make([]byte, 1000000) // todo configure this
		rand.Read(data)
		resp, err := client.State.SubmitPayForBlob(context.Background(), nID, data, math.NewInt(50000), 9000000)
		if err != nil {
			panic(err)
		}
		fmt.Println("resp code: ", resp.Code)
		fmt.Println("resp height: ", resp.Height, "\n")
		time.Sleep(15 * time.Second)
	}
}
