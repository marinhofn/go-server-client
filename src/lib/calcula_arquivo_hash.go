package main

import (
    "fmt"
    "os"
    "go-cliente-servidor/src/helpers"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: ./programa <caminho_do_arquivo>")
        return
    }

    filePath := os.Args[1]
    hash, err := helpers.SumStringFile(filePath)
    if err != nil {
        fmt.Printf("Erro ao calcular o hash: %v\n", err)
        return
    }

    fmt.Printf("Hash do arquivo %s: %s\n", filePath, hash)
}
