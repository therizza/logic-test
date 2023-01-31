package main

import (
    "bufio"
	"fmt"
    "os"
	"sort"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n := 0 
    fmt.Sscan(scanner.Text(), &n)
    cavalos := make([]int, n)
    for i := 0; i < n; i++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &cavalos[i])
    }
    sort.Ints(cavalos)
    mininoDiferenca :=  10000000
    for i :=1; i < n; i ++ {
        diferenca := cavalos[i] - cavalos[i-1]
        if diferenca < mininoDiferenca{
            mininoDiferenca = diferenca
        }
    }
    fmt.Println(mininoDiferenca)

}