package main

import (
	"fmt"
	"github.com/pudgekim/gocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println("Data: ", block.Data)
	}
}
