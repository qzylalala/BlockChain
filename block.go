/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/14 15:23
 */

package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp 		int64 	// when the block is created
	Data 			[]byte	// valuable information
	PrevBlockHash 	[]byte	// hash of the previous block
	Hash 			[]byte	// hash of this block
}


func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}




