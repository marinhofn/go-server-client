
## Equipe

**Membros**:
- Huandrey Souza Pontes
- Igor Correia da Silva
- José Marinho Falcão Neto
- Thayane Stheffany Silva Barros

---
Este repositório é um simulador de transferência de arquivos utilizando uma abordagem peer-to-peer (P2P). O objetivo deste projeto é permitir o compartilhamento de arquivos entre nós da rede de maneira distribuída, onde cada nó pode atuar tanto como servidor quanto como cliente.

Os arquivos que você deseja compartilhar devem ser colocados na pasta tmp/dataset. O sistema oferece três funções principais:

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

### Configuração dos IPs no Cliente

Antes de rodar o cliente, você deve configurar os IPs dos servidores no arquivo `client.go`. Localize a seguinte estrutura no código:

```go
servers := []struct {
    ip   string
    port string
}{
    {"150.165.74.64", "9000"}, 
}