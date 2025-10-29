package xutils

import (
	"github.com/thiagozs/go-xutils/aes"
	"github.com/thiagozs/go-xutils/bools"
	"github.com/thiagozs/go-xutils/calc"
	"github.com/thiagozs/go-xutils/cep"
	"github.com/thiagozs/go-xutils/cnpj"
	"github.com/thiagozs/go-xutils/convs"
	"github.com/thiagozs/go-xutils/cpf"
	"github.com/thiagozs/go-xutils/csv"
	"github.com/thiagozs/go-xutils/email"
	"github.com/thiagozs/go-xutils/files"
	"github.com/thiagozs/go-xutils/geo"
	"github.com/thiagozs/go-xutils/hash"
	"github.com/thiagozs/go-xutils/ip"
	"github.com/thiagozs/go-xutils/phone"
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
	hash    *hash.Hash
	rsa     *rsa.RSA
	rsaPem  *rsa.RSAPem
	aes     *aes.AES
	csv     *csv.CSV
	bools   *bools.Bools
	cpf     *cpf.CPF
	structs *structs.Structs
	convs   *convs.Convs
	xls     *xls.XLS
	phone   *phone.Phone
	ip      *ip.Ip
	geo     *geo.Geo
	cep     *cep.CEP
	files   *files.Files
}

func New() *XUtils {
	return &XUtils{
		slices:  slices.New(strings.New()),
		str:     strings.New(),
		calc:    calc.New(),
		cnpj:    cnpj.New(),
		email:   email.New(),
		hash:    hash.New(),
		rsa:     rsa.New(),
		rsaPem:  rsa.NewPem(),
		aes:     aes.New(),
		csv:     csv.New(),
		bools:   bools.New(),
		cpf:     cpf.New(),
		structs: structs.New(),
		convs:   convs.New(),
		xls:     xls.New(),
		phone:   phone.New(),
		ip:      ip.New(),
		geo:     geo.New(),
		cep:     cep.New(),
		files:   files.New(),
	}
}

// (RNG centralized in package randutil)

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

func (x *XUtils) Hash() *hash.Hash {
	return x.hash
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

func (x *XUtils) Phone() *phone.Phone {
	return x.phone
}

func (x *XUtils) Ip() *ip.Ip {
	return x.ip
}

func (x *XUtils) Geo() *geo.Geo {
	return x.geo
}

func (x *XUtils) CEP() *cep.CEP {
	return x.cep
}

func (x *XUtils) Files() *files.Files {
	return x.files
}
