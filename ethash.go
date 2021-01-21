package main

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/valar999/stratum-submit/ethash"
	"strconv"
)

func GetEthash(blockNum uint64, headerHash, nonce string) (mixDigest, hash string, err error) {
	cache := ethash.SharedEthash.Cache(blockNum)
	size := ethash.DatasetSize(blockNum)
	hashNoNonce := common.HexToHash(headerHash)
	uintNonce, err := strconv.ParseUint(nonce, 16, 64)
	if err != nil {
		return
	}
	digest, hashBinary := ethash.HashimotoLight(size, cache.Cache,
		hashNoNonce.Bytes(), uintNonce)
	mixDigest = hex.EncodeToString(digest[:])
	hash = hex.EncodeToString(hashBinary[:])
	return
}
