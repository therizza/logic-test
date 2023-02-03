package main

import (
	"fmt"
)

func main() {
	// Ler o número de entradas do dicionário
	var n int
	fmt.Scan(&n)

	// Criar um mapa para armazenar os pares código-caractere
	dict := make(map[string]byte)

	// Ler cada par código-caractere e adicioná-lo ao dicionário
	for i := 0; i < n; i++ {
		var code string
		var char byte
		fmt.Scan(&code, &char)
		dict[code] = char
	}

	// Ler a string codificada
	var encoded string
	fmt.Scan(&encoded)

	// Inicializar a string decodificada
	decoded := ""

	// Manter o rastro do índice inicial para o próximo código
	start := 0

	// Loop até que tenhamos processado toda a string codificada
	for start < len(encoded) {
		for k := 1; k <= len(encoded)-start; k++ {
			// Verificar se o código atual existe no dicionário
			if char, ok := dict[encoded[start:start+k]]; ok {
				// Se existir, adicionar o caractere correspondente à string decodificada
				decoded += string(char)

				// Mover o índice inicial para o próximo código
				start += k
				break
			}

			// Se chegarmos ao final da string sem encontrar uma correspondência, imprimir um erro e retornar
			if k == len(encoded)-start {
				fmt.Printf("DECODE FAIL AT INDEX %d", start)
				return
			}
		}
	}

	// Imprimir a string decodificada
	fmt.Println(decoded)
}
