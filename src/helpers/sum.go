package helpers

import (
	"io"
	"os"
)

func Sum(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sum := 0
	buf := make([]byte, 1024) // Lê o arquivo em blocos de 1024 bytes

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}

		// Processa os bytes lidos
		for _, b := range buf[:n] {
			if b >= '0' && b <= '9' { // Verifica se é um dígito
				sum += int(b - '0') // Converte o byte para o número e soma
			}
		}
	}

	return sum, nil
}
