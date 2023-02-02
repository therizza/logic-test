package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestSequence é uma função que recebe uma string de bits como entrada
// e retorna o comprimento da maior sequência de 1's contíguos.
func longestSequence(bits string) int {
	// maxLen é a variável que armazena o comprimento da maior sequência de 1's até o momento.
	maxLen := 0
	// Laço for que percorre toda a string de bits.
	for i := range bits {
		// Verifica se o bit na posição i é 0.
		if bits[i] == '0' {
			// Variáveis que armazenam a quantidade de 1's à esquerda e à direita do bit 0.
			left, right := 0, 0
			// Laço que percorre os bits à esquerda de i até encontrar um 0 ou chegar no início da string.
			for j := i - 1; j >= 0 && bits[j] == '1'; j-- {
				left++
			}
			// Laço que percorre os bits à direita de i até encontrar um 0 ou chegar no final da string.
			for j := i + 1; j < len(bits) && bits[j] == '1'; j++ {
				right++
			}
			// curLen é o comprimento da sequência de 1's contíguos que inclui o bit 0.
			curLen := left + right + 1
			// Verifica se curLen é maior do que maxLen e, se sim, atualiza maxLen.
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	// Se maxLen é 0, isso significa que todos os bits são 1's. Neste caso, retornamos o comprimento da string.
	if maxLen == 0 {
		return len(bits)
	}
	// Retorna o comprimento da maior sequência de 1's.
	return maxLen
}

func main() {
	// Cria um novo scanner para ler a entrada padrão.
	scanner := bufio.NewScanner(os.Stdin)
	// Lê a primeira linha da entrada padrão e a armazena em bits.
	scanner.Scan()
	bits := scanner.Text()
	// Imprime o comprimento da maior sequência de 1's.
	fmt.Println(longestSequence(bits))
}
