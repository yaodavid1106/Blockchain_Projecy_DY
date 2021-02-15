package Block

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	PreviousBlockHash []byte
	Timestamp         []byte
	Data              []byte
	Hash              []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) calculateHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousBlockHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, PreviousBlockHash []byte) *Block {
	block := &Block{PreviousBlockHash, []byte(time.Now().String()), []byte(data), []byte{}}
	block.calculateHash()
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
