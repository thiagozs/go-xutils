package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils"
)

func main() {
	utils := xutils.New()

	limit, offset, err := utils.Calc().CalculateLimitAndOffsetStr("10", "10")
	if err != nil {
		panic(err)
	}

	println("Limit and offset:", limit, offset)

	aes := utils.AES().RegisterKeys("IgkibX71IEf382PT", "IgkibX71IEf382PT")
	aesc, err := aes.Encrypt("123456")
	if err != nil {
		panic(err)
	}

	println("AES EncryptKey:", aesc)

	aesdec, err := aes.Decrypt(aesc)
	if err != nil {
		panic(err)
	}

	println("AES DEcryptKey:", aesdec)

	priv, pub := utils.RSAPem().RSAGenKeyPair()

	privpem := utils.RSAPem().RSAExportPrivateKeyAsPem(priv)

	println("RSA PrivateKey:", privpem)

	pubpem, err := utils.RSAPem().RSAExportPublicKeyAsPem(pub)
	if err != nil {
		panic(err)
	}

	println("RSA PublicKey:", pubpem)

	b := utils.RSA().PublicKey(pubpem)
	rsae, err := b.Encrypt("123456")
	if err != nil {
		panic(err)
	}

	println("RSA EncryptKey:", rsae)

	a := utils.RSA().PrivateKey(privpem)
	rsad, err := a.Decrypt(rsae)
	if err != nil {
		panic(err)
	}

	println("RSA DEcryptKey:", rsad)

	println("MD5 Hash:", utils.Hash().MD5("123456"))

	bt, err := utils.Convs().ToBool("true")
	if err != nil {
		panic(err)
	}

	println("ToBool:", bt)

	num, err := utils.Convs().ToInt64("123456")
	if err != nil {
		panic(err)
	}

	println("ToInt64:", num)

	mm, err := utils.XLS().ParseToMap("./xls/data/test.xlsx")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	header := []string{}
	for k, v := range mm {
		for kk, vv := range v {
			println(k, kk, vv)
			if k == 0 {
				header = append(header, vv)
			}
		}
	}

	fmt.Printf("headers %v\n", header)

	phone, err := utils.Phone().Normalize("11999999999", "BR")
	if err != nil {
		panic(err)
	}

	println("Phone:", phone)

	phone, err = utils.Phone().Normalize("1936114444", "BR")
	if err != nil {
		panic(err)
	}

	println("Phone:", phone)
}
