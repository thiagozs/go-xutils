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

// MD5 returns the hexadecimal MD5 checksum. It must not be used for passwords
// or cryptographic integrity; it exists for legacy checksum interoperability.
func MD5(str string) string {
	s := cmd5.New()
	_, _ = s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func IsMD5(v string) bool { return len(v) == 32 && IsHex(v) }

func IsBase64(v string) bool {
	if len(v) == 0 {
		return false
	}

	return base64Regex.MatchString(v)
}

func IsBase64URL(v string) bool {
	if len(v) == 0 {
		return false
	}

	return base64URLRegex.MatchString(v)
}

func IsHex(v string) bool { return hexRegex.MatchString(v) }

func IsBin(v string) bool { return binRegex.MatchString(v) }

func IsHexColor(v string) bool { return hexColorRegex.MatchString(v) }

func IsRGBColor(v string) bool { return rgbColorRegex.MatchString(v) }
