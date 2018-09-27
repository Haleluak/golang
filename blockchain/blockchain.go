package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) addBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func newBlockchain() *Blockchain {
	return &Blockchain{[]*Block{newGenesisBlock()}}
}
