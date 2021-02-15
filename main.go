package main

import (
	"fmt"

	"github.com/yaodavid1106/Blockchain_Projecy_DY/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("This is the second block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash : %x\n", block.PreviousBlockHash)
		fmt.Printf("Time : %x\n", block.Timestamp)
		fmt.Printf("Data : %x\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Println()
	}
}
