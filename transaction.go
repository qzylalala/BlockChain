/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/15 11:13
 */

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsity = 10

//  an input references a previous output
type TXInput struct {
	Txid      []byte	// ID of such transaction
	Vout      int		// an index of an output in the transaction
	ScriptSig string	// a script which provides data to be used in an output’s ScriptPubKey
}

//  Outputs are where “coins” are stored
type TXOutput struct {
	Value        int	// the value of "coins"
	ScriptPubKey string	// store an arbitrary string (user defined wallet address).
}

type Transaction struct {
	ID 			[]byte
	Vin 		[]TXInput
	Vout 		[]TXOutput
}

func (tx *Transaction) setID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func newCoinBaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsity, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.setID()

	return &tx
}

