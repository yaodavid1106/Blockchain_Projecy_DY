package blockchain

import (
	"time"
)

type Block struct {
	PreviousBlockHash []byte
	Timestamp         []byte
	Data              []byte
	Nonce             int
	Hash              []byte
}

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(data string, PreviousBlockHash []byte) *Block {
	block := &Block{PreviousBlockHash, []byte(time.Now().String()), []byte(data), 0, []byte{}}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("This is Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
