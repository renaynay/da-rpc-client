package main

import (
	"context"
	"cosmossdk.io/math"
	"crypto/rand"
	"fmt"

	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

func main() {
	token := ""
	client, err := client.NewClient(context.Background(), "", token)
	if err != nil {
		panic(err)
	}

	nID := make([]byte, 8)
	rand.Read(nID)
	data := make([]byte, 10)
	rand.Read(data)
	resp, err := client.State.SubmitPayForBlob(context.Background(), nID, data, math.NewInt(500), 700000)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp: ", resp)

}
