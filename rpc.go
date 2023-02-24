package main

import (
	"context"
	"cosmossdk.io/math"
	"crypto/rand"
	"fmt"
	"github.com/celestiaorg/celestia-node/api/rpc/client"
	"os"
	"time"
)

func main() {
	addr := os.Args[1]
	token := os.Args[2]
	addr2 := os.Args[3]
	token2 := os.Args[4]

	fillBlocks(token, addr, token2, addr2)
}

// 256 size block
func fillBlocks(token, addr, token2, addr2 string) {
	client1, err := client.NewClient(context.Background(), addr, token)
	if err != nil {
		panic(err)
	}
	client2, err := client.NewClient(context.Background(), addr2, token2)
	if err != nil {
		panic(err)
	}

	for {
		go func() {
			nID := make([]byte, 8)
			rand.Read(nID)
			data := make([]byte, 1000000) // todo configure this
			rand.Read(data)
			resp, err := client1.State.SubmitPayForBlob(context.Background(), nID, data, math.NewInt(30000), 9000000)
			if err != nil {
				panic(err)
			}
			fmt.Println("client1: resp code: ", resp.Code)
			fmt.Println("client1: resp height: ", resp.Height, "\n")
		}()
		go func() {
			nID := make([]byte, 8)
			rand.Read(nID)
			data := make([]byte, 1005000) // todo configure this
			rand.Read(data)
			resp, err := client2.State.SubmitPayForBlob(context.Background(), nID, data, math.NewInt(500000), 90000000)
			if err != nil {
				panic(err)
			}
			fmt.Println("client2: resp code: ", resp.Code)
			fmt.Println("client2: resp height: ", resp.Height, "\n")
		}()
		time.Sleep(15 * time.Second)
	}
}
