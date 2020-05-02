package main

func main() {
	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// var lines []string
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }

	// key := lines[0]
	// text := lines[1]

	// fmt.Printf("Key = %v\n", key)
	// fmt.Printf("Original text = %v\n", text)

	// encryptedText := paddText([]byte(text))
	// encryptedText = desEncrypt(encryptedText, []byte(key))

	// decryptedText := desDecrypt(encryptedText, []byte(key))
	// decryptedText = unpaddText(decryptedText)

	// encryptedText := cbcEncrypt([]byte(text), []byte(key), make([]byte, blockSize))
	// decryptedText := cbcDecrypt(encryptedText, []byte(key), make([]byte, blockSize))

	// encryptedText := cfbEncrypt([]byte(text), []byte(key), make([]byte, blockSize))
	// decryptedText := cfbDecrypt(encryptedText, []byte(key), make([]byte, blockSize))

	// encryptedText := ctrEncrypt([]byte(text), []byte(key))
	// decryptedText := ctrDecrypt(encryptedText, []byte(key))

	// messageWithMAC := createHMAC(encryptedText, []byte(key))

	// if len(messageWithMAC) < hashLength {
	// 	fmt.Println("Invalid MAC")
	// 	return
	// }

	// message := messageWithMAC[:len(messageWithMAC)-hashLength]
	// hashMAC := messageWithMAC[len(messageWithMAC)-hashLength:]

	// if !validateHMAC(message, hashMAC, []byte(key)) {
	// 	fmt.Println("Invalid MAC")
	// 	return
	// }

	// fmt.Printf("Encrypted Text: %q\n", string(encryptedText))
	// fmt.Printf("Decrypyed Text: %v", string(decryptedText))

	// f, _ := os.Create("output.txt")

	// defer f.Close()

	// f.Write(encryptedText)

}
