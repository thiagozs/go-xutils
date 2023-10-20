package main

import "github.com/thiagozs/go-xutils"

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

	println("MD5 Hash:", utils.MD5().MD5Hash("123456"))

}
