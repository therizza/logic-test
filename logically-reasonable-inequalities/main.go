package main

import (
	"bufio"   // importa a biblioteca bufio para a leitura de entrada
	"fmt"     // importa a biblioteca fmt para formatar a saída
	"os"      // importa a biblioteca os para acesso aos recursos do sistema operacional
)

const maxLetters = 26   // define o número máximo de letras como 26 (A-Z)

// declara as variáveis globais utilizadas no programa
var graph [maxLetters][maxLetters]bool  // representa o grafo de relações
var visited [maxLetters]bool             // indica se um nó foi visitado ou não
var stack [maxLetters]bool               // usado para detectar ciclos no grafo
var n int                                 // número de relações de entrada

// implementa a busca em profundidade (DFS) para verificar se há ciclos no grafo
func dfs(node int) bool {
	visited[node] = true   // marca o nó atual como visitado
	stack[node] = true     // adiciona o nó atual na pilha

	// itera por todos os nós adjacentes ao nó atual
	for i := 0; i < n; i++ {
		if graph[node][i] {   // verifica se há uma relação entre o nó atual e o nó i
			if !visited[i] && dfs(i) {  // se o nó i não foi visitado e há um ciclo ao visitar o nó i
				return true
			} else if stack[i] {  // se o nó i já foi visitado e está na pilha, há um ciclo
				return true
			}
		}
	}

	stack[node] = false   // remove o nó atual da pilha
	return false
}

func main() {
	fmt.Scanf("%d\n", &n)   // lê o número de relações de entrada

	reader := bufio.NewReader(os.Stdin)  // cria um leitor de entrada
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')  // lê cada linha de entrada
		var a, b rune  // declara as variáveis para as duas letras da relação
		fmt.Sscanf(line, "%c > %c\n", &a, &b)  // lê as duas letras da relação
		graph[int(a-'A')][int(b-'A')] = true  // adiciona a relação no grafo
	}

	// chama a função dfs para cada variável que ainda não foi visitada
	for i := 0; i < n; i++ {
		if !visited[i] {
			if dfs(i) {
				fmt.Println("contradiction")
				return
			}
		}
	}

	fmt.Println("consistent")
}
