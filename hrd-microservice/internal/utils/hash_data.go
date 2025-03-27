package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt akan mengenkripsi data menggunakan AES dengan kunci yang ditentukan
func Encrypt(data string) string {
	// Kunci enkripsi (gunakan kunci yang aman dan rahasia)
	key := []byte("your-32-byte-long-encryption-key!") // Sesuaikan panjang kunci (32 byte untuk AES-256)

	// Inisialisasi cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Menambahkan padding ke data yang ingin dienkripsi agar panjangnya kelipatan blok (AES)
	dataBytes := []byte(data)
	blockSize := block.BlockSize()
	padding := blockSize - len(dataBytes)%blockSize
	paddedData := append(dataBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Inisialisasi IV (Initialization Vector)
	iv := make([]byte, blockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		panic(err)
	}

	// Mengenkripsi data
	ciphertext := make([]byte, len(paddedData))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, paddedData)

	// Gabungkan IV dan ciphertext menjadi satu byte slice untuk disimpan
	result := append(iv, ciphertext...)

	// Mengembalikan data terenkripsi dalam format base64 agar dapat disimpan dalam database
	return base64.StdEncoding.EncodeToString(result)
}

// Decrypt akan mendekripsi data yang sudah dienkripsi
func Decrypt(encryptedData string) (string, error) {
	// Kunci enkripsi (gunakan kunci yang aman dan rahasia yang sama dengan yang digunakan di enkripsi)
	key := []byte("your-32-byte-long-encryption-key!") // Pastikan ini sama dengan kunci yang digunakan untuk enkripsi

	// Mengubah data terenkripsi dari base64 menjadi byte slice
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// Mengecek apakah panjang data valid
	if len(data) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// Mengambil IV dari data terenkripsi (IV berada di depan ciphertext)
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	// Inisialisasi cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Dekripsi data
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// Menghilangkan padding (mengetahui padding berdasarkan panjang blok)
	padding := int(ciphertext[len(ciphertext)-1])
	if padding > aes.BlockSize || padding > len(ciphertext) {
		return "", errors.New("invalid padding")
	}
	ciphertext = ciphertext[:len(ciphertext)-padding]

	// Mengembalikan hasil dekripsi sebagai string
	return string(ciphertext), nil
}
