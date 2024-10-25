package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func Crypt(texto string) (string, error) {
	// Criar um bloco AES com a chave fornecida
	bloco, err := aes.NewCipher([]byte(os.Getenv("AES_CYPHERING_KEY")))
	if err != nil {
		return "", err
	}

	// Gerar um nonce (número único que só será usado uma vez)
	gcm, err := cipher.NewGCM(bloco)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Criptografar o texto
	textoCriptografado := gcm.Seal(nonce, nonce, []byte(texto), nil)
	return hex.EncodeToString(textoCriptografado), nil
}

func Decrypt(textoCriptografado string) (string, error) {
	// Decodificar o texto criptografado do formato hexadecimal
	dados, err := hex.DecodeString(textoCriptografado)
	if err != nil {
		return "", err
	}

	// Criar um bloco AES com a chave fornecida
	bloco, err := aes.NewCipher([]byte(os.Getenv("AES_CYPHERING_KEY")))
	if err != nil {
		return "", err
	}

	// Criar GCM (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(bloco)
	if err != nil {
		return "", err
	}

	// Verificar o tamanho do nonce
	nonceSize := gcm.NonceSize()
	if len(dados) < nonceSize {
		return "", errors.New("dados criptografados são muito curtos")
	}

	// Separar o nonce dos dados criptografados
	nonce, dadosCriptografados := dados[:nonceSize], dados[nonceSize:]

	// Descriptografar os dados
	textoDescriptografado, err := gcm.Open(nil, nonce, dadosCriptografados, nil)
	if err != nil {
		return "", err
	}

	return string(textoDescriptografado), nil
}
