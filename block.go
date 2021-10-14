/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/14 15:23
 */

package main

import (
	"time"
)

type Block struct {
	Timestamp 		int64 	// when the block is created
	Data 			[]byte	// valuable information
	PrevBlockHash 	[]byte	// hash of the previous block
	Hash 			[]byte	// hash of this block
	Counter			int		// counter
}


//func (b *Block) setHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
//	hash := sha256.Sum256(headers)
//
//	b.Hash = hash[:]
//}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := newProofOfWork(block)
	counter, hash := pow.Run()

	block.Hash = hash[:]
	block.Counter = counter

	return block
}




