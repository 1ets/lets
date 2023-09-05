package lets

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Handle load and save keys to memory/storage.
type RsaKeys struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey

	PrivateKeyFile string
	PublicKeyFile  string
}

////////////
// LOADER //
////////////

// Load Keys file.
func (r *RsaKeys) Load() (err error) {
	err = r.LoadPrivateKey()
	if err != nil {
		return
	}

	err = r.LoadPublicKey()
	if err != nil {
		return
	}

	return
}

// Load file and setup private key.
func (r *RsaKeys) LoadPrivateKey() (err error) {
	privateKey, err := os.ReadFile(r.PrivateKeyFile)
	if err != nil {
		return
	}

	return r.SetPrivateKey(privateKey)
}

// Load file and setup public key.
func (r *RsaKeys) LoadPublicKey() (err error) {
	publicKey, err := os.ReadFile(r.PublicKeyFile)
	if err != nil {
		return
	}

	return r.SetPublicKey(publicKey)
}

///////////
// SAVER //
///////////

// Save all keys into storage.
func (r *RsaKeys) SavePKCS1() (err error) {
	LogI("SAVED: %s", r.PrivateKeyFile)
	err = r.SavePrivateKeyPKCS1()
	if err != nil {
		return
	}

	err = r.SavePublicKeyPKCS1()
	if err != nil {
		return
	}

	return
}

// Save PrivateKey to storage in PKCS1.
func (r *RsaKeys) SavePrivateKeyPKCS1() (err error) {
	pemFile, err := os.Create(r.PrivateKeyFile)
	if err != nil {
		return
	}

	pemKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(r.PrivateKey),
	}

	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return
	}

	pemFile.Close()
	return
}

// Save PublicKey to storage in PKCS1.
func (r *RsaKeys) SavePublicKeyPKCS1() (err error) {
	pemFile, err := os.Create(r.PublicKeyFile)
	if err != nil {
		return
	}

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(r.PublicKey),
	}

	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return
	}

	pemFile.Close()
	return
}

// Save all keys into storage in PKCS8.
func (r *RsaKeys) SavePKCS8() (err error) {
	LogI("SAVED: %s", r.PrivateKeyFile)

	err = r.SavePrivateKeyPKCS8()
	if err != nil {
		return
	}

	err = r.SavePublicKeyPKIX()
	if err != nil {
		return
	}

	return
}

// Save PrivateKey to storage in PKCS8 format.
func (r *RsaKeys) SavePrivateKeyPKCS8() (err error) {
	pemFile, err := os.Create(r.PrivateKeyFile)
	if err != nil {
		return
	}

	privateKeyPKCS8, err := x509.MarshalPKCS8PrivateKey(r.PrivateKey)
	if err != nil {
		return
	}

	pemKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyPKCS8,
	}

	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return
	}

	pemFile.Close()
	return
}

// Save PublicKey to storage in PKIX format.
func (r *RsaKeys) SavePublicKeyPKIX() (err error) {
	pemFile, err := os.Create(r.PublicKeyFile)
	if err != nil {
		return
	}

	publicKeyPKIX, err := x509.MarshalPKIXPublicKey(r.PublicKey)
	if err != nil {
		return
	}

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPKIX,
	}

	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return
	}

	pemFile.Close()
	return
}

////////////
// SETUPS //
////////////

// Parses a PEM encoded private key.
func (r *RsaKeys) SetPrivateKey(privateKey []byte) (err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("invalid private key")
		return
	}

	r.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return r.setPrivateKeyPKCS8(privateKey)
	}
	return
}

// Parses a PEM encoded private key.
func (r *RsaKeys) setPrivateKeyPKCS8(privateKey []byte) (err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("invalid private key")
		return
	}

	parsedPKCS8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	fmt.Println(reflect.TypeOf(parsedPKCS8))

	// TODO: add if parsed using pkcs
	r.PrivateKey = parsedPKCS8.(*rsa.PrivateKey)

	return
}

// Set the private key string.
func (r *RsaKeys) SetPrivateKeyString(privateKey string) (err error) {
	return r.SetPrivateKey([]byte(privateKey))
}

// Parses a PEM encoded public key.
func (r *RsaKeys) SetPublicKey(publicKey []byte) (err error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("invalid public key")
		return
	}

	var key interface{}

	key, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		LogD("Cannot parse using PKCS1, trying parse using PKIX.")
		return r.setPublicKeyPKIX(publicKey)
	}

	switch keyType := key.(type) {
	case *rsa.PublicKey:
		r.PublicKey = keyType
	default:
		err = errors.New("invalid public key type")

	}

	return
}

func (r *RsaKeys) setPublicKeyPKIX(publicKey []byte) (err error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("invalid public key")
		return
	}

	var key interface{}

	key, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	switch keyType := key.(type) {
	case *rsa.PublicKey:
		r.PublicKey = keyType
	default:
		err = errors.New("invalid public key type")

	}

	return
}

// Set the public key string.
func (r *RsaKeys) SetPublicKeyString(publicKey string) (err error) {
	return r.SetPublicKey([]byte(publicKey))
}
