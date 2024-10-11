package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PasswordGenerator é uma estrutura que contém informações sobre a geração da senha
type PasswordGenerator struct {
	upperChars   []byte
	lowerChars   []byte
	numberChars  []byte
	specialChars []byte
}

// NewPasswordGenerator inicializa uma nova instância de PasswordGenerator
func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{
		upperChars:   []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		lowerChars:   []byte("abcdefghijklmnopqrstuvwxyz"),
		numberChars:  []byte("0123456789"),
		specialChars: []byte("!@#$%^&*()-_=+,.?/:;{}[]`~"),
	}
}

// Generate gera uma senha aleatória com base nos comprimentos mínimo e máximo
func (pg *PasswordGenerator) Generate(minLength, maxLength int) string {
	if maxLength < minLength {
		maxLength = minLength
	}

	// Define o comprimento da senha aleatoriamente entre minLength e maxLength
	length := rand.Intn(maxLength-minLength+1) + minLength
	newPassword := make([]byte, length)

	// Garante que a senha contenha pelo menos um de cada tipo de caractere
	newPassword[0] = pg.upperChars[rand.Intn(len(pg.upperChars))]
	newPassword[1] = pg.lowerChars[rand.Intn(len(pg.lowerChars))]
	newPassword[2] = pg.numberChars[rand.Intn(len(pg.numberChars))]
	newPassword[3] = pg.specialChars[rand.Intn(len(pg.specialChars))]

	// Preenche o restante da senha com caracteres aleatórios
	allChars := append(append(pg.upperChars, pg.lowerChars...), append(pg.numberChars, pg.specialChars...)...)

	for i := 4; i < length; i++ {
		newPassword[i] = allChars[rand.Intn(len(allChars))]
	}

	// Embaralha a senha
	rand.Shuffle(length, func(i, j int) {
		newPassword[i], newPassword[j] = newPassword[j], newPassword[i]
	})

	return string(newPassword)
}

// readInt solicita um número inteiro do usuário, lidando com entradas inválidas
func readInt(prompt string) int {
	var value int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, insira um número inteiro.")
			var discard string
			fmt.Scan(&discard) // Limpa o buffer de entrada
			continue
		}
		break
	}
	return value
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Cria uma nova instância de PasswordGenerator
	passwordGenerator := NewPasswordGenerator()

	// Lê comprimentos mínimo e máximo do usuário
	minLength := readInt("Por favor, informe o comprimento mínimo: ")
	maxLength := readInt("Por favor, informe o comprimento máximo: ")

	// Gera e exibe a senha
	password := passwordGenerator.Generate(minLength, maxLength)
	fmt.Printf("Esta é a sua senha gerada por este programa em Go: %v\n", password)
}