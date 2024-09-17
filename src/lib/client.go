package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Função para conectar ao servidor e enviar o hash
func checkHashOnServer(ip, port, hash string) (string, error) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", hash)

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}

	return message, nil
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "search" {
		fmt.Println("Uso correto: ./programa search hash_de_arquivo")
		return
	}

	hash := os.Args[2]
	servers := []struct {
		ip   string
		port string
	}{
		{"127.0.0.1", "9001"},
		{"127.0.0.1", "9002"},
		{"127.0.0.1", "9003"},
	}

	for _, server := range servers {
		response, err := checkHashOnServer(server.ip, server.port, hash)
		response = strings.TrimSpace(response)
		fmt.Printf("Resposta do servidor: '%s' a \n", response)
		if err != nil {
			fmt.Printf("Erro ao conectar ao servidor %s:%s: %v\n", server.ip, server.port, err)
			continue
		}

		if response == "found\n" {
			fmt.Printf("Arquivo encontrado no servidor %s:%s\n", server.ip, server.port)
		} else {
			fmt.Printf("Arquivo não encontrado no servidor %s:%s\n", server.ip, server.port)
		}
	}
}
