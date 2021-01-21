package main

import (
	"math/big"
	"testing"
)

func TestSubmit(t *testing.T) {
	// this data have to be filled manually
	extraNonce := "0a4e"
	height := uint64(0xb26a70)
	job := []string{"2c93d","82232565de6c6216a88e3d3b4a49c8ad5e5913b731e78c93c7e29e50b9f0743f","e045fdb8e5418242e7b4fde12c2bf1047e098d80c19ef75768a671c06861d01d"}
	submit := []string{"XXX.NH01","2c93d","ad1202e9d241"}
	// end of data

	headerHash := job[2]
	nonce := extraNonce + submit[2]
	_, hash, _ := GetEthash(height, headerHash, nonce)
	target1 := big.NewFloat(2.695994666715064e+67)
	hashBig, _, _ := new(big.Float).Parse(hash, 16)
	diffBig := new(big.Float).Quo(target1, hashBig)
	t.Log("diff", diffBig)
}
