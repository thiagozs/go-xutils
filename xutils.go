package xutils

import (
	"github.com/thiagozs/go-xutils/aes"
	"github.com/thiagozs/go-xutils/bools"
	"github.com/thiagozs/go-xutils/calc"
	"github.com/thiagozs/go-xutils/cnpj"
	"github.com/thiagozs/go-xutils/convs"
	"github.com/thiagozs/go-xutils/cpf"
	"github.com/thiagozs/go-xutils/csv"
	"github.com/thiagozs/go-xutils/email"
	"github.com/thiagozs/go-xutils/md5"
	"github.com/thiagozs/go-xutils/rsa"
	"github.com/thiagozs/go-xutils/slices"
	"github.com/thiagozs/go-xutils/strings"
	"github.com/thiagozs/go-xutils/structs"
	"github.com/thiagozs/go-xutils/xls"
)

type XUtils struct {
	calc    *calc.Calc
	cnpj    *cnpj.CNPJ
	email   *email.Email
	slices  *slices.Slices
	str     *strings.Strings
	md5     *md5.Md5
	rsa     *rsa.RSA
	rsaPem  *rsa.RSAPem
	aes     *aes.AES
	csv     *csv.CSV
	bools   *bools.Bools
	cpf     *cpf.CPF
	structs *structs.Structs
	convs   *convs.Convs
	xls     *xls.XLS
}

func New() *XUtils {
	return &XUtils{
		slices:  slices.New(strings.New()),
		str:     strings.New(),
		calc:    calc.New(),
		cnpj:    cnpj.New(),
		email:   email.New(),
		md5:     md5.New(),
		rsa:     rsa.New(),
		rsaPem:  rsa.NewPem(),
		aes:     aes.New(),
		csv:     csv.New(),
		bools:   bools.New(),
		cpf:     cpf.New(),
		structs: structs.New(),
		convs:   convs.New(),
		xls:     xls.New(),
	}
}

func (x *XUtils) Strings() *strings.Strings {
	return x.str
}

func (x *XUtils) Slices() *slices.Slices {
	return x.slices
}

func (x *XUtils) Calc() *calc.Calc {
	return x.calc
}

func (x *XUtils) CNPJ() *cnpj.CNPJ {
	return x.cnpj
}

func (x *XUtils) Email() *email.Email {
	return x.email
}

func (x *XUtils) MD5() *md5.Md5 {
	return x.md5
}

func (x *XUtils) RSA() *rsa.RSA {
	return x.rsa
}

func (x *XUtils) RSAPem() *rsa.RSAPem {
	return x.rsaPem
}

func (x *XUtils) AES() *aes.AES {
	return x.aes
}

func (x *XUtils) CSV() *csv.CSV {
	return x.csv
}

func (x *XUtils) Bools() *bools.Bools {
	return x.bools
}

func (x *XUtils) CPF() *cpf.CPF {
	return x.cpf
}

func (x *XUtils) Structs() *structs.Structs {
	return x.structs
}

func (x *XUtils) Convs() *convs.Convs {
	return x.convs
}

func (x *XUtils) XLS() *xls.XLS {
	return x.xls
}
