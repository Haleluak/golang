package main

import (
	"fmt"
)

func main() {
	bc := newBlockchain()

	bc.addBlock("Send 1 BTC to HDJoker")
	bc.addBlock("Send 2 BTC to Hack4ever")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.prevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
