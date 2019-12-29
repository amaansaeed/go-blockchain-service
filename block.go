package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      Data
	Hash      string
	PrevHash  string
}

type Data struct {
	Transaction string `json:"transaction"`
	Total       int    `json:"total"`
	IsGenesis   bool   `json:"is_genesis"`
}

func (b *Block) generateHash() {
	// convert the data into a string
	bytes, _ := json.Marshal(b.Data)
	// concatenate all the elements as strings
	data := strconv.Itoa(b.Index) + b.Timestamp.String() + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, data Data) *Block {
	b := &Block{}
	b.Index = prevBlock.Index + 1
	b.Timestamp = time.Now()
	b.Data = data
	b.PrevHash = prevBlock.Hash
	b.generateHash()

	return b
}

// func ProcessTransaction() int {

// }

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, Data{IsGenesis: true})
}
