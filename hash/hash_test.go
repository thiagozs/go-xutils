package hash

import (
	"testing"
)

func TestMD5(t *testing.T) {
	hasher := New()
	str := "hello"
	expectedMD5 := "5d41402abc4b2a76b9719d911017c592" // MD5 for "hello"

	if md5 := hasher.MD5(str); md5 != expectedMD5 {
		t.Errorf("MD5(%s) = %s; want %s", str, md5, expectedMD5)
	}
}

func TestIsMD5(t *testing.T) {
	hasher := New()
	validMD5 := "5d41402abc4b2a76b9719d911017c592"
	invalidMD5 := "ZGVmZzg3ZjhlYmM0ZjI3NzZiOTcxOWQ5MTEwMTdjNTky"

	if !hasher.IsMD5(validMD5) {
		t.Errorf("IsMD5(%s) = false; want true", validMD5)
	}

	if hasher.IsMD5(invalidMD5) {
		t.Errorf("IsMD5(%s) = true; want false", invalidMD5)
	}
}

func TestIsBase64(t *testing.T) {
	hasher := New()
	validBase64 := "SGVsbG8sIHdvcmxkIQ=="
	invalidBase64 := "SGVsbG8sIHdvcmxkIQ"

	if !hasher.IsBase64(validBase64) {
		t.Errorf("IsBase64(%s) = false; want true", validBase64)
	}

	if hasher.IsBase64(invalidBase64) {
		t.Errorf("IsBase64(%s) = true; want false", invalidBase64)
	}
}

func TestIsBase64URL(t *testing.T) {
	hasher := New()
	validBase64URL := "SGVsbG8sIHdvcmxkIQ"
	invalidBase64URL := "SGVsbG8sIHdvcmxkIQ==="

	if !hasher.IsBase64URL(validBase64URL) {
		t.Errorf("IsBase64URL(%s) = false; want true", validBase64URL)
	}

	if hasher.IsBase64URL(invalidBase64URL) {
		t.Errorf("IsBase64URL(%s) = true; want false", invalidBase64URL)
	}
}

func TestIsHex(t *testing.T) {
	hasher := New()
	validHex := "deadBEEF"
	invalidHex := "nothex"

	if !hasher.IsHex(validHex) {
		t.Errorf("IsHex(%s) = false; want true", validHex)
	}

	if hasher.IsHex(invalidHex) {
		t.Errorf("IsHex(%s) = true; want false", invalidHex)
	}
}

func TestIsBin(t *testing.T) {
	hasher := New()
	validBin := "101010"
	invalidBin := "10201"

	if !hasher.IsBin(validBin) {
		t.Errorf("IsBin(%s) = false; want true", validBin)
	}

	if hasher.IsBin(invalidBin) {
		t.Errorf("IsBin(%s) = true; want false", invalidBin)
	}
}

func TestIsHexColor(t *testing.T) {
	hasher := New()
	validHexColor := "#fff"
	invalidHexColor := "not"

	if !hasher.IsHexColor(validHexColor) {
		t.Errorf("IsHexColor(%s) = false; want true", validHexColor)
	}

	if hasher.IsHexColor(invalidHexColor) {
		t.Errorf("IsHexColor(%s) = true; want false", invalidHexColor)
	}
}

func TestIsRGBColor(t *testing.T) {
	hasher := New()
	validRGBColor := "rgb(255, 255, 255)"
	invalidRGBColor := "rgb(255, 255, 256)"

	if !hasher.IsRGBColor(validRGBColor) {
		t.Errorf("IsRGBColor(%s) = false; want true", validRGBColor)
	}

	if hasher.IsRGBColor(invalidRGBColor) {
		t.Errorf("IsRGBColor(%s) = true; want false", invalidRGBColor)
	}
}
