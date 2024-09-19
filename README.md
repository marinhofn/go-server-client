Este repositório é um simulador de transferência de arquivos de maneira peer-to-peer. Os arquivos que você está disposto a compartilhar devem estar na pasta 'tmp/dataset'. Disponibilizamos 3 funções principais:

### Para rodar o código do servidor

```bash
go run src/lib/server.go {porta_desejada}
```

### Para rodar o código do cliente:

```bash
go run src/lib/client.go search {hash_do_arquivo}
```

### Para calcular o hash de um arquivo:

```bash
go run src/lib/calcula_arquivo_hash.go {caminho_do_arquivo}
```