# Migrating from v1 to v2

Version 2 intentionally removes the root service locator and unsafe legacy
cryptography. Update the module path in every import:

The minimum supported toolchain is Go 1.26.6.

```diff
- github.com/thiagozs/go-xutils/email
+ github.com/thiagozs/go-xutils/v2/email
```

## Root facade

Import focused packages and call package functions:

```diff
- utils := xutils.New()
- ok := utils.CPF().IsValid(value)
+ ok := cpf.IsValid(value)
```

The same pattern applies to `bools`, `cep`, `cnpj`, `email`, `geo`, `hash`,
`ip`, and `structs`.

Conversion and calculation helpers were simplified:

```diff
- convs.NewConverter[int32]().StringToType(value)
+ convs.Parse[int32](value)

- calc.New().CalculateLimitAndOffset(page, size)
+ calc.LimitOffset(page, size)
```

## Cryptography

AES-CBC and RSA PKCS#1 v1.5 were removed. Ciphertexts produced by those APIs
must be migrated explicitly before upgrading.

```go
cipher, err := aes.NewCipher(key)
encoded, err := cipher.Encrypt(plaintext)
plaintext, err := cipher.Decrypt(encoded)
```

RSA v2 uses OAEP with SHA-256:

```go
privateKey, publicKey, err := rsa.GenerateKeyPair(3072)
encoded, err := rsa.EncryptOAEP(publicKey, plaintext)
plaintext, err := rsa.DecryptOAEP(privateKey, encoded)
```

PEM helpers are now `rsa.ExportPrivateKey`, `rsa.ExportPublicKey`,
`rsa.ParsePrivateKey`, and `rsa.ParsePublicKey`.

## Corrected behavior

- `slices.ContainsAll(values, required)` replaces `AreKeysValid` and correctly
  requires every required key to be present.
- Slice transforms are package functions, do not mutate their input, and use
  intention-revealing names such as `slices.Normalize`, `slices.Unique`, and
  `slices.Compact`.
- Pure text transforms are package functions. Random strings use
  `strings.NewGenerator`; `EscapeString` was replaced by the context-specific
  `strings.EscapeSQLLike`.
- `files.RemoveDir` removes only empty directories; use `RemoveAllDir` for
  recursive deletion.
- CSV parsing rejects duplicate/empty headers and mismatched field counts.
- CPF, CNPJ, and CEP validation accepts only canonical raw or formatted input.
- Unsupported countries no longer produce Brazilian phone numbers silently.
- Phone generation methods are now `GenerateMobile`, `GenerateLandline`,
  `GenerateMobileWithMask`, and `GenerateLandlineWithMask`; they generate
  Brazilian numbers and no longer accept a misleading country argument.
