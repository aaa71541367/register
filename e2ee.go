package goline

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	rand2 "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
	"strings"
	"time"

	curve "github.com/miguelsandro/curve25519-go/axlsign"
)

func randomBytes(size int) []uint8 {
	rand.Seed(time.Now().UTC().UnixNano())
	var High int = 255
	var seed = make([]uint8, size)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(rand.Int() % (High + 1))
	}
	return seed
}

func GenerateAsymmetricKeypair() curve.Keys {
	seed := randomBytes(32)
	return curve.GenerateKeyPair(seed)
}

func GenerateSharedKey(privateKey, publicKey []byte) (sharedKey []byte, err error) {
	if len(privateKey) != 32 || len(publicKey) != 32 {
		return sharedKey, fmt.Errorf("invalid length of key pair")
	}
	sharedKey = curve.SharedKey(privateKey, publicKey)
	return sharedKey, nil
}

func CreateSecretQuery(PublicKey []uint8) string {
	return url.QueryEscape(base64.StdEncoding.EncodeToString(PublicKey))
}

func HalfXorData(data []byte) []byte {
	lenbuf := len(data)
	res := make([]byte, lenbuf/2)
	for i := 0; i < lenbuf/2; i++ {
		res[i] = data[i] ^ data[lenbuf/2+i]
	}
	return res
}

func aesCBCEncrypt(encodeBytes, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	encodeBytes = PKCS7Padding(encodeBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)
	return crypted, nil
}

func aesCBCDecrypt(encryptedData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptedData, encryptedData)
	encryptedData, err = PKCS7UnPadding(encryptedData)
	if err != nil {
		return []byte{}, err
	}
	return encryptedData, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if (length-unpadding) < 0 || (length-unpadding) >= len(origData) {
		return nil, fmt.Errorf("invalid length: %d %d", length, unpadding)
	}
	return origData[:(length - unpadding)], nil
}

func EncryptPassword(passwd string, keypair curve.Keys, mynonce []uint8, public_key, nonce string) (string, error) {
	publicKey, _ := base64.StdEncoding.DecodeString(public_key)
	exchangedNonce, _ := base64.StdEncoding.DecodeString(nonce)
	sharedSecret, err := GenerateSharedKey(keypair.PrivateKey, publicKey)
	if err != nil {
		return "", err
	}
	masterKey := sha256.Sum256(append([]byte("master_key"), append(sharedSecret, append(mynonce, exchangedNonce...)...)...))
	keyAndIv := sha256.Sum256(append([]byte("aes_key"), masterKey[:]...))
	hmacKey := sha256.Sum256(append([]byte("hmac_key"), masterKey[:]...))
	ciphertext, err := aesCBCEncrypt([]byte(passwd), keyAndIv[:16], keyAndIv[16:32])
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, hmacKey[:])
	h.Write(ciphertext)
	return base64.StdEncoding.EncodeToString(append(ciphertext, h.Sum(nil)...)), nil
}

func createPinCode() (string, error) {
	pin, err := rand2.Int(rand2.Reader, big.NewInt(999999))
	if err != nil {
		return "", err
	}
	// padding
	return fmt.Sprintf("%06d", pin), nil
}

func xor(data []byte) (r []byte) {
	length := len(data) / 2
	r = make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = data[i] ^ data[length+i]
	}
	return r
}

func GetSignature(authKey string, rev int) string {
	split := strings.SplitN(authKey, ":", 2)
	if len(split) > 3 || len(split) < 2 {
		return ""
	}
	currentMillis := time.Now().UTC().UnixNano() / int64(time.Millisecond)
	var key []byte
	var lastID = len(split) - 1
	key, _ = base64.StdEncoding.DecodeString(split[lastID])
	msg, meta := "", ""
	if rev == 1 { // androidlite
		msg = fmt.Sprintf("issuedTo: %s\niat: %d\n", split[0], currentMillis)
		meta = base64.StdEncoding.EncodeToString([]byte("type: YWT\nalg: HMAC_SHA1\n"))
	} else { // android
		msg = fmt.Sprintf("iat: %d\n", currentMillis)
		meta = ""
	}
	split[lastID] = base64.StdEncoding.EncodeToString([]byte(msg)) + "." + meta
	//
	h := hmac.New(sha1.New, key)
	h.Write([]byte(split[lastID]))
	split[lastID] += "." + base64.StdEncoding.EncodeToString(h.Sum(nil))
	return strings.Join(split, ":")
}
