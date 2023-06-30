package main

import (
	// O nome crand é um apelido para "crypto/rand"
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func generatePassword(minLength int, maxLength int) string {
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	var length = rand.Intn(minLength) + minLength

	newPassword := make([]byte, length)
	randomData := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(crand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPassword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPassword)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Print("Por favor, informe o comprimento mínimo: ")
	var minLength int
	fmt.Scanf("%d", &minLength)

	//fmt.Print("Por favor, informe o comprimento máximo: ")
	var maxLength int
	fmt.Scanf("%d", &maxLength)

	fmt.Printf("Esta é a sua senha gerada por este programa em Go: %v\n", generatePassword(minLength, maxLength))
}
