# go-xutils v2

Biblioteca Go com pacotes focados para validação, transformação de dados,
arquivos e criptografia.

Requer Go 1.26.6 ou superior.

> A v2 possui mudanças incompatíveis. Consulte [MIGRATION_V2.md](MIGRATION_V2.md)
> antes de atualizar uma aplicação v1.

## Instalação

```bash
go get github.com/thiagozs/go-xutils/v2@v2.0.0-rc.1
```

Importe apenas o pacote necessário:

```go
import (
    "github.com/thiagozs/go-xutils/v2/email"
    "github.com/thiagozs/go-xutils/v2/ip"
)

validEmail := email.IsValid("dev@example.com")
validIP := ip.IsValid("2001:db8::1")
```

Utilitários sem estado são funções de pacote, como `cpf.IsValid`,
`cnpj.Normalize`, `cep.Format`, `strings.CamelCase` e `geo.IsLatitude`.
Transformações de `slices` não alteram a entrada. Estado configurável fica em
tipos explícitos, como `strings.Generator`, `csv.Parser` e `files.Files`.

## Criptografia

AES utiliza GCM autenticado. O nonce aleatório é incluído no valor codificado:

```go
key := []byte("0123456789abcdef0123456789abcdef")
cipher, err := aes.NewCipher(key)
encoded, err := cipher.Encrypt([]byte("segredo"))
plaintext, err := cipher.Decrypt(encoded)
```

RSA utiliza OAEP com SHA-256. AES-CBC e RSA PKCS#1 v1.5 não fazem parte da v2.

## Pacotes

- Documentos e validação: `cep`, `cnpj`, `cpf`, `email`, `geo`, `ip`, `phone`.
- Transformação: `bools`, `calc`, `convs`, `slices`, `strings`, `structs`.
- Entrada e saída: `csv`, `files`, `xls`.
- Criptografia e codificação: `aes`, `hash`, `rsa`.

## Exemplos

Há um programa executável para cada pacote em [examples](examples/README.md).
Por exemplo:

```bash
go run ./examples/cpf
go run ./examples/files
go run ./examples/rsa
```

## Qualidade

```bash
go test ./...
go test -race ./...
go vet ./...
golangci-lint run ./...
go run golang.org/x/vuln/cmd/govulncheck@v1.7.0 ./...
```

Consulte [ARCHITECTURE.md](ARCHITECTURE.md) para as regras de design e
[CHANGELOG.md](CHANGELOG.md) para as mudanças da versão.

Licença MIT.
