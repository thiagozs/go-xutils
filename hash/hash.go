package hash

import (
	cmd5 "crypto/md5"
	"encoding/hex"
	"regexp"
)

var (
	base64Regex = regexp.MustCompile(
		`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|` +
			`[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`,
	)

	base64URLRegex = regexp.MustCompile(
		`^([A-Za-z0-9_-]{4})*([A-Za-z0-9_-]{2}(==)?|[A-Za-z0-9_-]{3}=?)?$`,
	)

	hexRegex = regexp.MustCompile(`^(#|0x)?[0-9a-fA-F]+$`)

	binRegex = regexp.MustCompile(`^(0b)?[01]+$`)

	hexColorRegex = regexp.MustCompile(`^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)

	rgbColorRegex = regexp.MustCompile(
		`^(rgb|RGB)\(\s*([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\s*,` +
			`\s*([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\s*,` +
			`\s*([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\s*\)$`,
	)
)

type Hash struct{}

func New() *Hash {
	return &Hash{}
}

func (h *Hash) MD5(str string) string {
	s := cmd5.New()
	_, _ = s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func (h *Hash) IsMD5(v string) bool {
	return len(v) == 32 && h.IsHex(v)
}

func (h *Hash) IsBase64(v string) bool {
	if len(v) == 0 {
		return false
	}

	return base64Regex.MatchString(v)
}

func (h *Hash) IsBase64URL(v string) bool {
	if len(v) == 0 {
		return false
	}

	return base64URLRegex.MatchString(v)
}

func (h *Hash) IsHex(v string) bool {
	return hexRegex.MatchString(v)
}

func (h *Hash) IsBin(v string) bool {
	return binRegex.MatchString(v)
}

func (h *Hash) IsHexColor(v string) bool {
	return hexColorRegex.MatchString(v)
}

func (h *Hash) IsRGBColor(v string) bool {
	return rgbColorRegex.MatchString(v)
}
