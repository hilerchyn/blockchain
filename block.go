package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type Block struct {
	Index     int    // is the position of the data record in the blockchain
	Timestamp string // is automatically determined and is the time the data is written
	BPM       int    // or beats per minute, is your pulse rate
	Hash      string // is a SHA256 identifier representing this data record
	PrevHash  string // is the SHA256 identifier of the previous record in the chain
}

// Blockchain is a series of validated Blocks
var Blockchain []Block
var mutex = &sync.Mutex{}

func calculateHash(block Block) string {
	record := fmt.Sprint(block.Index) + block.Timestamp + fmt.Sprint(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock

}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
