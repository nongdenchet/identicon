package encode

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hasher.Sum(nil)

	size := hex.EncodedLen(len(hash))
	result := make([]byte, size)
	hex.Encode(result, hash)

	return result
}
