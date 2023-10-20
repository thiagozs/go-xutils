package rsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RSASuite struct {
	suite.Suite
	pem          *RSAPem
	rsa          *RSA
	publicKey    string
	privateKey   string
	strEncrypted string
	toEncrypt    string
}

func (suite *RSASuite) SetupTest() {
	suite.pem = NewPem()
	suite.rsa = New()

	priv, pub := suite.pem.RSAGenKeyPair()

	strPub, err := suite.pem.RSAExportPublicKeyAsPem(pub)
	if err != nil {
		suite.T().Error("rsa export public key error", err)
		return
	}

	strPriv := suite.pem.RSAExportPrivateKeyAsPem(priv)
	if err != nil {
		suite.T().Error("rsa export private key error", err)
		return
	}

	suite.publicKey = strPub
	suite.privateKey = strPriv
	suite.toEncrypt = "123456"

	strEncrypted, err := suite.rsa.PublicKey(strPub).Encrypt(suite.toEncrypt)
	if err != nil {
		suite.T().Error("rsa public encrypt error", err)
		return
	}

	suite.strEncrypted = strEncrypted
}

func (suite *RSASuite) TestEncrypt() {
	p := suite.rsa.PublicKey(suite.publicKey)
	str, err := p.Encrypt(suite.toEncrypt)
	if err != nil {
		suite.T().Error("rsa public encrypt error", err)
		return
	}
	suite.T().Log(str)
}

func (suite *RSASuite) TestDecrypt() {
	p := suite.rsa.PrivateKey(suite.privateKey)
	str, err := p.Decrypt(suite.strEncrypted)
	if err != nil {
		suite.T().Error("rsa private decrypt error", err)
		return
	}
	suite.T().Log(str)

	assert.Equal(suite.T(), suite.toEncrypt, str)
}

func TestRSASuite(t *testing.T) {
	suite.Run(t, new(RSASuite))
}
