package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data Data) {
	// get previous block
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// create new block
	block := CreateBlock(prevBlock, data)
	bc.blocks = append(bc.blocks, block)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
