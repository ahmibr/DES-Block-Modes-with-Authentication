package main

import (
	"crypto/des"
	"strconv"
)

func ecbEncrypt(text, key []byte) []byte {
	paddedText := paddText(text)

	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(paddedText))
	for i := 0; i < len(paddedText); i += blockSize {
		block.Encrypt(ret[i:], paddedText[i:i+blockSize])
	}
	return ret
}

func cbcEncrypt(text, key, iv []byte) []byte {
	paddedText := paddText(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(paddedText))
	tmp := make([]byte, blockSize)
	for i := 0; i < len(paddedText); i += blockSize {
		tmp = xorBlock(iv, paddedText[i:i+blockSize])
		block.Encrypt(ret[i:], tmp)
		iv = ret[i : i+blockSize]
	}
	return ret
}

func cfbEncrypt(text, key, iv []byte) []byte {
	paddedText := paddText(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(paddedText))
	tmp := make([]byte, blockSize)
	for i := 0; i < len(paddedText); i += blockSize {
		block.Encrypt(tmp, iv)
		copy(ret[i:i+blockSize], xorBlock(tmp, paddedText[i:i+blockSize]))
		iv = ret[i : i+blockSize]
	}
	return ret
}

func ctrEncrypt(text, key []byte) []byte {
	paddedText := paddText(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	ret := make([]byte, len(paddedText))
	tmp := make([]byte, blockSize)
	for i := 0; i < len(paddedText); i += blockSize {
		ctr := make([]byte, blockSize)
		copy(ctr[blockSize-len(strconv.Itoa(i/blockSize)):], []byte(strconv.Itoa(i/blockSize)))
		block.Encrypt(tmp, ctr)
		copy(ret[i:i+blockSize], xorBlock(tmp, paddedText[i:i+blockSize]))
	}
	return ret
}
