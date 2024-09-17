
### Para rodar o código do servidor

```bash
go run src/lib/server.go {porta_desejada}
```

### Para rodar o código do cliente:

```bash
go run src/lib/client.go search ae5e34f70239c2d2634ec7780f38d526a797a2b0095b1128cb4a066b0be110c1
```



### Ao mudar o arquivo do cliente ou servidor é necessário buildar o arquivo, para isso:

```bash
go build -o bin/server src/lib/server.go
```

```bash
go build -o bin/client src/lib/client.go
```
