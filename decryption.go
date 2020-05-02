package main

import (
	"crypto/des"
	"strconv"
)

func ecbDecrypt(text, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(text))
	for i := 0; i < len(text); i += blockSize {
		block.Decrypt(ret[i:], text[i:i+blockSize])
	}

	return unpaddText(ret)
}

func cbcDecrypt(text, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(text))
	tmp := make([]byte, blockSize)
	for i := 0; i < len(text); i += blockSize {
		block.Decrypt(tmp, text[i:i+blockSize])
		tmp := xorBlock(iv, tmp)
		copy(ret[i:i+blockSize], tmp)
		iv = text[i : i+blockSize]
	}

	return unpaddText(ret)
}

func cfbDecrypt(text, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(text))
	tmp := make([]byte, blockSize)

	for i := 0; i < len(text); i += blockSize {
		block.Encrypt(tmp, iv)
		tmp := xorBlock(tmp, text[i:i+blockSize])
		copy(ret[i:i+blockSize], tmp)
		iv = text[i : i+blockSize]
	}

	return unpaddText(ret)
}

func ctrDecrypt(text, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(text))
	tmp := make([]byte, blockSize)

	for i := 0; i < len(text); i += blockSize {
		ctr := make([]byte, blockSize)
		copy(ctr[blockSize-len(strconv.Itoa(i/blockSize)):], []byte(strconv.Itoa(i/blockSize)))
		block.Encrypt(tmp, ctr)
		copy(ret[i:i+blockSize], xorBlock(tmp, text[i:i+blockSize]))
	}

	return unpaddText(ret)
}
