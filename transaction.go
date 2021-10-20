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
	"encoding/hex"
	"fmt"
	"log"
)

const subsity = 10 		// the amount of reward

//  an input references a previous output
type TXInput struct {
	Txid      []byte	// ID of the transaction
	Vout      int		// Index of an output in the transaction
	ScriptSig string	// a script which provides data to be used in an output’s ScriptPubKey
}

//  Outputs are where “coins” are stored (indivisible)
type TXOutput struct {
	Value        int	// the value of "coins"
	ScriptPubKey string	// store an arbitrary string (user defined wallet address).
}

type Transaction struct {
	ID 			[]byte
	Vin 		[]TXInput
	Vout 		[]TXOutput
}

func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
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

func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.setID()

	return &tx
}