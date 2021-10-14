/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/14 15:43
 */

package main

import "fmt"

func main() {
	bc := newBlockChain()

	bc.addBlock("Send 1 BTC to qzylalalla")
	bc.addBlock("Send 2 more BTC to qzylalala")

	for _, block := range bc.blocks {
		fmt.Printf("Previous block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
