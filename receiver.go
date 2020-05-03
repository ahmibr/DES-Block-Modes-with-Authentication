package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const buffSize = 4096

func main() {
	file, _ := os.Open("key.txt")

	key := make([]byte, keyLength)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		copy(key, []byte(scanner.Text()))
	}

	file.Close()

	fmt.Println("Read key " + string(key))

	socket, _ := net.Dial("tcp", "localhost:12345")
	defer socket.Close()

	receivedData := make([]byte, buffSize)
	for {
		socket.Read(receivedData)

		mode := string(receivedData)[:3]
		fmt.Printf("User choose %v mode\n", mode)

		n, _ := socket.Read(receivedData)

		message := []byte(receivedData[:n-hashLength])
		messageMAC := []byte(receivedData[n-hashLength : n])

		fmt.Printf("Received encrypted text %q\n", message)

		if !validateHMAC(message, messageMAC, key) {
			fmt.Println("Invalid MAC, cannot decrypt")
			continue
		}

		iv := make([]byte, blockSize)
		decryptedText := ""

		switch mode {
		case "ECB":
			decryptedText = string(ecbDecrypt([]byte(message), key))
		case "CBC":
			decryptedText = string(cbcDecrypt([]byte(message), key, iv))
		case "CFB":
			decryptedText = string(cfbDecrypt([]byte(message), key, iv))
		case "CTR":
			decryptedText = string(ctrDecrypt([]byte(message), key))
		default:
			panic("An error has occured")
		}

		fmt.Printf("Decrypted text: %v\n", decryptedText)
	}

}
