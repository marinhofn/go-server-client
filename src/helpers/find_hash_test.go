package helpers

import (
	"os"
	"testing"
	"fmt"
)

func TestFindHash(t *testing.T) {
	testDir := t.TempDir()
	fmt.Printf("Diretório de teste: %s\n", testDir)
	filePath := testDir + "/test_file.txt"
	fileContent := []byte("conteúdo de teste para gerar o hash")
	err := os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		t.Fatalf("Erro ao criar arquivo de teste: %v", err)
	}

	expectedHash, err := Sum(filePath)
	if err != nil {
		t.Fatalf("Erro ao calcular hash do arquivo de teste: %v", err)
	}

	found, err := FindHash(expectedHash, testDir)
	if err != nil {
		t.Errorf("Erro ao executar FindHash: %v", err)
	}

	if !found {
		t.Errorf("Esperava que o hash fosse encontrado, mas não foi")
	}

	found, err = FindHash("hash_que_nao_existe", testDir)
	if err != nil {
		t.Errorf("Erro ao executar FindHash: %v", err)
	}

	if found {
		t.Errorf("Não esperava que o hash fosse encontrado, mas foi")
	}
}
