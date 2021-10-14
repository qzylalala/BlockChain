/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/14 15:39
 */

package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) addBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := newBlock(data, prevBlock.Hash)

	bc.blocks = append(bc.blocks, newBlock)
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}

func newBlockChain() *Blockchain {
	return &Blockchain{[]*Block{newGenesisBlock()}}
}

