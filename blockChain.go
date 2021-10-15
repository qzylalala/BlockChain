/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/14 15:39
 */

package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type Blockchain struct {
	tip 	[]byte		// the hash value of the last block of blockchain
	db 		*bolt.DB	// store a DB connection
}

type BlockchainIterator struct {
	currentHash []byte		// inspect current hash value
	db 			*bolt.DB	// store a DB connection
}

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

func (bc *Blockchain) addBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := newBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		// hash value -> block
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		// update the last block file number
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}

// 'l' -> 4-byte file number: the last block file number used (K-V pair)
func newBlockChain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile,0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {	// empty Blockchain
			genesis := newGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {	// Blockchain is not empty
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// returns the current block and point to the previous block
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

