package license

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type License struct {
	Nonce string `json:"nonce"`
	Time  int64  `json:"time"`
	Use   int64  `json:"use"`
}

func (l *License) String() string {
	var e error
	var j []byte
	if j, e = json.Marshal(l); e != nil {
		beego.Error("license to string ", e)
		os.Exit(-1)
	}
	return string(j)
}

var FileLicense = "conf/license"
var FileSignature = "conf/signature"

const TimeLimit = int64(3600) * 24 * 15
const TimeSpan  = 30
const UseLimit = int64(150)

func InitLicense(pathLicense string, pathSignature string) error {
	InitPath(pathLicense, pathSignature)
	var e error
	var u uuid.UUID
	var l *License
	var public *rsa.PublicKey
	var private *rsa.PrivateKey

	if public, private, e = LoadKey(); e != nil {
		beego.Error(e)
		return e
	}
	var client = PkcsClient{Private: private, Public: public}

	l = new(License)
	if e = l.ReadLicense(); e != nil {
		beego.Error(e)
		return e
	}

	l.Use = 0
	l.Time = 0

	if u, e = uuid.NewV1(); e != nil {
		beego.Error(e)
		return e
	}

	l.Nonce = u.String()

	if e = l.WriteLicense(); e != nil {
		beego.Error(e)
		return e
	}

	var signature []byte
	if signature, e = client.Sign([]byte(l.String()), crypto.SHA256); e != nil {
		beego.Error(e)
		return e
	}

	if e = WriteSignature(signature); e != nil {
		beego.Error(e)
		return e
	}

	return nil
}

func InitPath(pathLicense string, pathSignature string) {
	FileLicense = pathLicense
	FileSignature = pathSignature
}

func LoadKey() (*rsa.PublicKey, *rsa.PrivateKey, error) {
	var e error
	var private *rsa.PrivateKey
	if private, e = LoadPrivateKey(); e != nil {
		return nil, nil, e
	}

	var public *rsa.PublicKey
	if public, e = LoadPublicKey(); e != nil {
		return nil, nil, e
	}

	return public, private, e
}

func Daemon(out chan bool, pathLicense string, pathSignature string) error {
	InitPath(pathLicense, pathSignature)
	var e error
	var u uuid.UUID
	var l *License

	var public *rsa.PublicKey
	var private *rsa.PrivateKey
	if public, private, e = LoadKey(); e != nil {
		return e
	}

	var client = PkcsClient{Private: private, Public: public}
	var launch = true
	l = new(License)
Main:
	for {
		select {
		case value := <-out:
			if value {
				e = errors.New("select out")
				break Main
			}
		default:

		}

		if u, e = uuid.NewV1(); e != nil {
			break
		}

		if e = l.ReadLicense(); e != nil {
			break
		}

		var signature []byte
		signature, e = ReadSignature()

		if e = client.Verify([]byte(l.String()), signature, crypto.SHA256); e != nil {
			break
		}

		lastTime := l.Time
		l.Nonce = u.String()
		l.Time = lastTime + TimeSpan

		if launch {
			lastUse := l.Use
			l.Use = lastUse + 1
			launch = false
		}

		if l.Time > TimeLimit {
			e = errors.New("license out of date")
			break
		}

		if l.Use > UseLimit {
			e = errors.New("license out of use limit")
			break
		}

		time.Sleep(TimeSpan * time.Second)

		if e = l.WriteLicense(); e != nil {
			break
		}

		if signature, e = client.Sign([]byte(l.String()), crypto.SHA256); e != nil {
			break
		}

		if e = WriteSignature(signature); e != nil {
			break
		}
	}
	beego.Error(e)
	os.Exit(-1)
	return nil
}

func (l *License) WriteLicense() error {
	var e error
	var file *os.File
	if file, e = os.OpenFile(FileLicense, os.O_RDWR, os.ModePerm); e != nil {
		return e
	}
	defer file.Close()

	if e = file.Truncate(0); e != nil {
		return e
	}
	if _, e = file.Seek(0, 0); e != nil {
		return e
	}
	if _, e = file.WriteString(l.String()); e != nil {
		return e
	}
	return nil
}

func (l *License) ReadLicense() error {
	var e error
	var file *os.File
	if file, e = os.OpenFile(FileLicense, os.O_RDWR, os.ModePerm); e != nil {
		return e
	}
	defer file.Close()

	var decoder = json.NewDecoder(file)
	for {
		if e = decoder.Decode(l); e == io.EOF {
			break
		} else if e != nil {
			return e
		}
	}
	return nil
}

func WriteSignature(signed []byte) error {
	var e error
	var file *os.File
	if file, e = os.OpenFile(FileSignature, os.O_RDWR, os.ModePerm); e != nil {
		return e
	}
	defer file.Close()

	if e = file.Truncate(0); e != nil {
		return e
	}
	if _, e = file.Seek(0, 0); e != nil {
		return e
	}
	if _, e = file.WriteString(hex.EncodeToString(signed)); e != nil {
		return e
	}
	return nil
}

func ReadSignature() ([]byte, error) {
	var e error
	var file *os.File
	if file, e = os.OpenFile(FileSignature, os.O_RDWR, os.ModePerm); e != nil {
		return nil, e
	}
	defer file.Close()

	var signatureHex []byte
	if signatureHex, e = ioutil.ReadAll(file); e != nil {
		return nil, e
	}

	var signature = make([]byte, len(signatureHex)/2)
	if _, e = hex.Decode(signature, signatureHex); e != nil {
		return nil, e
	}

	return signature, nil
}

func FileInit() {
	var t = "12345678"
	var e error

	var private *rsa.PrivateKey
	if private, e = LoadPrivateKey(); e != nil {
		beego.Error(e)
		return
	}
	var public *rsa.PublicKey
	if public, e = LoadPublicKey(); e != nil {
		beego.Error(e)
		return
	}

	var client = PkcsClient{Private: private, Public: public}

	var signed []byte
	if signed, e = client.Sign([]byte(t), crypto.SHA256); e != nil {
		beego.Error(e)
		return
	}

	fmt.Println(signed)
	fmt.Println(len(signed))

	var licenseFile *os.File

	if licenseFile, e = os.OpenFile("../../conf/license", os.O_RDWR, 0666); e != nil {
		return
	}

	_, _ = licenseFile.WriteString(hex.EncodeToString(signed))
	_ = licenseFile.Close()
}

func String() ([]byte, error) {
	var licenseFile *os.File
	var e error
	if licenseFile, e = os.OpenFile("../conf/license", os.O_RDWR, 0666); e != nil {
		return nil, e
	}
	var reader = bufio.NewReader(licenseFile)
	return ioutil.ReadAll(reader)
}

type PkcsClient struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

func (this *PkcsClient) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, this.Public, plaintext)
}

func (this *PkcsClient) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, this.Private, ciphertext)
}

func (this *PkcsClient) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, this.Private, hash, hashed)
}

func (this *PkcsClient) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(this.Public, hash, hashed, sign)
}

func GenRsaKey(bits int) error {
	// gen private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// gen public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func LoadPrivateKey() (*rsa.PrivateKey, error) {
	var e error
	privateKey := PrivateKey
	blockPri, _ := pem.Decode([]byte(privateKey))
	if blockPri == nil {
		return nil, errors.New("decode error")
	}

	priKey, e := x509.ParsePKCS1PrivateKey(blockPri.Bytes)
	if e != nil {
		return nil, e
	}
	return priKey, nil

}

func LoadPublicKey() (*rsa.PublicKey, error) {
	var e error
	var publicKey = PublicKey
	blockPub, _ := pem.Decode([]byte(publicKey))
	if blockPub == nil {
		fmt.Println("decode error")
		return nil, errors.New("decode error")
	}

	pubKey, e := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if e != nil {
		beego.Error(e)
		return nil, e
	}

	return pubKey.(*rsa.PublicKey), nil
}