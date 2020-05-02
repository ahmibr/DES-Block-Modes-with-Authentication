package main

const blockSize = 8
const hashLength = 32

func paddText(text []byte) []byte {
	paddingCount := blockSize - len(text)%blockSize
	for i := 0; i < paddingCount; i++ {
		text = append(text, byte(paddingCount))
	}
	return text
}

func unpaddText(text []byte) []byte {
	paddingCount := int(text[len(text)-1])
	return text[:(len(text) - paddingCount)]
}

func xorBlock(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		panic("Lengthes aren't equal")
	}

	ret := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		ret[i] = b1[i] ^ b2[i]
	}

	return ret
}
