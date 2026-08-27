# Changelog

## [2.0.0-rc.1] - Unreleased

### Added

- Authenticated AES-GCM encryption.
- RSA-OAEP with SHA-256 and safe PEM helpers.
- Package-level APIs for stateless validation and transformation.
- Non-mutating string-slice transformations and injectable random string
  generation.
- Stream-based CSV parsing and strict structural validation.
- Concurrency-safe pseudo-random source.
- CI checks for tests, race detection, vet, and lint.
- A vulnerability gate validated with `govulncheck`.

### Changed

- Module path now ends in `/v2`.
- Minimum Go version is 1.26.6.
- Excelize is 2.11.0 and `golang.org/x/text` is 0.39.0 or newer to include
  upstream security fixes.
- Filesystem operations preserve modes and symbolic links where applicable.
- Document, email, IP, phone, reflection, and numeric validation is stricter.

### Removed

- Root `xutils.New()` service-locator facade.
- AES-CBC compatibility API.
- RSA PKCS#1 v1.5 compatibility API.
- Stateless receiver wrappers for validation packages.
- Ambiguous text and slice receiver APIs, including the context-free escaping
  helper.
