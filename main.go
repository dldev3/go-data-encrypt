package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetters := string(runes[0 : len(letter)-key])
	return fmt.Sprintf(`%s %s`, lastLetterKey, leftOverLetters)
}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {

	}

	strings.Map(findOne, plainText)

}

func decrypt(key int, enctyptedText string) (result string) {

}

func main() {

	plainText := "HELLOWORLD"
	fmt.Println("Plain Text", plainText)

	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted text", encrypted)

	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted", decrypted)

}
