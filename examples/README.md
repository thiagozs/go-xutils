# Exemplos

Cada diretório contém um programa independente que demonstra a API pública de
um pacote da v2.

Execute um exemplo a partir da raiz do projeto:

```bash
go run ./examples/aes
go run ./examples/csv
go run ./examples/strings
```

Para apenas compilar e validar todos os exemplos:

```bash
go test ./examples/...
```

Os exemplos de `files` e `xls` trabalham em diretórios temporários. O exemplo
de `rsa` gera uma chave de 2048 bits durante a execução e, por isso, pode levar
um pouco mais de tempo.
