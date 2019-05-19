package keypair

import (
	"encoding/base64"
	"sync"

	"golang.org/x/net/context"

	lbcf "github.com/lidstromberg/config"
	stor "github.com/lidstromberg/storage"

	jwt "github.com/dgrijalva/jwt-go"

	"crypto/rand"
	"crypto/rsa"
	"io/ioutil"
)

//pub holder for a public key
type pub struct {
	key *rsa.PublicKey
	mux sync.Mutex
}

//pri holder for a private key
type pri struct {
	key *rsa.PrivateKey
	mux sync.Mutex
}

//KeyPair holds jwt encryption/decryption keys
type KeyPair struct {
	pubKey *pub
	priKey *pri
}

//GetPubKey returns the public key
func (kp *KeyPair) GetPubKey() *rsa.PublicKey {
	kp.pubKey.mux.Lock()
	defer kp.pubKey.mux.Unlock()
	return kp.pubKey.key
}

//GetPriKey returns the private key
func (kp *KeyPair) GetPriKey() *rsa.PrivateKey {
	kp.priKey.mux.Lock()
	defer kp.priKey.mux.Unlock()
	return kp.priKey.key
}

//newBucketKeyPair creates a new signing keypair from private/public keys stored in a GCP bucket
func newBucketKeyPair(ctx context.Context, bc lbcf.ConfigSetting) (*KeyPair, error) {
	pk1 := new(pub)
	pk2 := new(pri)
	keyPair := &KeyPair{pubKey: pk1, priKey: pk2}

	sto, err := stor.NewMgr(ctx, bc)
	if err != nil {
		return nil, err
	}

	signBytes, err := sto.GetBucketFileData(ctx, bc.GetConfigValue(ctx, "EnvKpGcpBucket"), bc.GetConfigValue(ctx, "EnvKpPrivateKey"))
	if err != nil {
		return nil, err
	}

	keyPair.priKey.key, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword(signBytes, bc.GetConfigValue(ctx, "EnvKpPrivateKeyCredential"))
	if err != nil {
		return nil, err
	}

	verifyBytes, err := sto.GetBucketFileData(ctx, bc.GetConfigValue(ctx, "EnvKpGcpBucket"), bc.GetConfigValue(ctx, "EnvKpPublicKey"))
	if err != nil {
		return nil, err
	}

	keyPair.pubKey.key, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return keyPair, nil
}

//newLocalKeyPair creates a new signing keypair from private/public keys stored in local storage
func newLocalKeyPair(ctx context.Context, priKey, priKeyPass, pubKey string) (*KeyPair, error) {
	pk1 := new(pub)
	pk2 := new(pri)
	keyPair := &KeyPair{pubKey: pk1, priKey: pk2}

	signBytes, err := ioutil.ReadFile(priKey)
	if err != nil {
		return nil, err
	}

	keyPair.priKey.key, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword(signBytes, priKeyPass)
	if err != nil {
		return nil, err
	}

	verifyBytes, err := ioutil.ReadFile(pubKey)
	if err != nil {
		return nil, err
	}

	keyPair.pubKey.key, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return keyPair, nil
}

//NewKeyPair creates a new signing keypair from private/public keys based on config settings
func NewKeyPair(ctx context.Context, bc lbcf.ConfigSetting) (*KeyPair, error) {
	preflight(ctx, bc)

	var kpr *KeyPair

	if bc.GetConfigValue(ctx, "EnvKpType") == "local" {
		kp1, err := newLocalKeyPair(ctx, bc.GetConfigValue(ctx, "EnvKpPrivateKey"), bc.GetConfigValue(ctx, "EnvKpPrivateKeyCredential"), bc.GetConfigValue(ctx, "EnvKpPublicKey"))
		if err != nil {
			return nil, err
		}

		kpr = kp1
	}

	if bc.GetConfigValue(ctx, "EnvKpType") == "bucket" {
		kp1, err := newBucketKeyPair(ctx, bc)
		if err != nil {
			return nil, err
		}

		kpr = kp1
	}

	return kpr, nil
}

//EncryptBytes uses the keypair to encrypt a byte array
func (kp *KeyPair) EncryptBytes(ctx context.Context, val []byte) (string, error) {
	encmsg, err := rsa.EncryptPKCS1v15(rand.Reader, kp.GetPubKey(), val)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(encmsg)

	return str, nil
}

//DecryptString uses the keypair to decrypt a base64 encrypted string
func (kp *KeyPair) DecryptString(ctx context.Context, val string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}

	decmsg, err := rsa.DecryptPKCS1v15(rand.Reader, kp.GetPriKey(), data)
	if err != nil {
		return "", err
	}

	return string(decmsg), nil
}
