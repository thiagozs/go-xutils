# Architecture

## Principles

`go-xutils/v2` is a collection of focused Go packages, not a framework.

1. Consumers import only the packages they use; the root has no facade.
2. Stateless behavior is exposed as package functions.
3. Constructors exist only for configured state, such as an encryption key,
   random source, phone generator, or I/O service.
4. Core parsers accept streams when practical; filesystem paths are adapters.
5. Invalid external input returns an error or `false`, never a panic.
6. Sentinel errors support `errors.Is`; contextual errors wrap their cause.
7. Security-sensitive APIs use authenticated modern primitives.
8. Shared mutable state must be concurrency-safe.

## Package boundaries

- Validation and Brazilian documents: `email`, `ip`, `geo`, `cpf`, `cnpj`,
  `cep`, `phone`.
- Text and data transformations: `strings`, `slices`, `convs`, `structs`,
  `bools`, `calc`.
- I/O adapters: `files`, `csv`, `xls`.
- Cryptography and encodings: `aes`, `rsa`, `hash`.

Packages never depend on the module root. Cross-package dependencies stay
explicit, and standard-library facilities are preferred over local wrappers.

## Compatibility

Breaking API changes require a new major module path. Behavior corrections that
can affect persisted data or destructive operations must be called out in the
changelog and migration guide. New legacy compatibility layers are not added to
the v2 module.
