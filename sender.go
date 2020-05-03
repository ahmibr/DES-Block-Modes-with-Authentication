package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("key.txt")

	key := make([]byte, keyLength)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		copy(key, []byte(scanner.Text()))
	}

	file.Close()

	fmt.Println("Read key " + string(key))

	listener, error := net.Listen("tcp", ":12345")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Waiting connection")
	socket, _ := listener.Accept()

	defer socket.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		encryptionModes := map[string]string{
			"1": "ECB",
			"2": "CBC",
			"3": "CFB",
			"4": "CTR",
		}

		fmt.Println("Choose mode of operation")
		for i := 1; i <= 4; i++ {
			fmt.Printf("%v %v\n", i, encryptionModes[strconv.Itoa(i)])
		}

		choosenMode := ""

		for {
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			if choosenMode = encryptionModes[text]; choosenMode != "" {
				break
			}
			fmt.Println("Choose correct mode")
		}

		socket.Write([]byte(choosenMode))

		fmt.Println("Write message to be encrypted")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		encryptedText := []byte{}

		iv := make([]byte, blockSize)

		fmt.Println("Message to be encrypted: " + text)
		switch choosenMode {
		case "ECB":
			encryptedText = ecbEncrypt([]byte(text), []byte(key))
		case "CBC":
			encryptedText = cbcEncrypt([]byte(text), []byte(key), iv)
		case "CFB":
			encryptedText = cfbEncrypt([]byte(text), []byte(key), iv)
		case "CTR":
			encryptedText = ctrEncrypt([]byte(text), []byte(key))
		default:
			panic("An error has occured")
		}

		// prepare message for sending
		encryptedText = createHMAC(encryptedText, []byte(key))

		socket.Write(encryptedText)
	}
}
