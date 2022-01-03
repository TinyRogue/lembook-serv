package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

const (
	kB                = 1 << 10
	argonParamsAmount = 6
	argonAlgType      = "argon2id"
)

var (
	InvalidHash           = errors.New("the encoded hash is not in the correct format")
	IncompatibleVersion   = errors.New("incompatible version of argon2")
	IncompatibleAlgorithm = errors.New("incompatible algorithm used of argon2xx")
)

type argonTune struct {
	kBmem       uint32
	iter        uint32
	parallelism uint8
	saltLen     uint32
	keyLen      uint32
}

func BeautifyPassword(password string, tuning *argonTune) (string, error) {
	if tuning == nil {
		tuning = getDefaultTuning()
	}

	salt, err := buySomeSalt(tuning.saltLen)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, tuning.iter, tuning.kBmem, tuning.parallelism, tuning.keyLen)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, tuning.kBmem, tuning.iter, tuning.parallelism, encodedSalt, encodedHash)

	return encodedHash, nil
}

func Compare(password, encodedHash string) (match bool, err error) {
	tuning, salt, hash, err := DecodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(password), salt, tuning.iter, tuning.kBmem, tuning.parallelism, tuning.keyLen)

	if subtle.ConstantTimeCompare(hash, hashToCompare) == 1 {
		return true, nil
	}
	return false, nil
}

func DecodeHash(encodedHash string) (tuning *argonTune, salt, hash []byte, err error) {
	splitHash := strings.Split(encodedHash, "$")
	if len(splitHash) != argonParamsAmount {
		return nil, nil, nil, InvalidHash
	}

	var algType string
	var version int
	tuning = &argonTune{}
	_, err = fmt.Sscanf(splitHash[1], "%s", &algType)
	if err != nil || algType != argonAlgType {
		return nil, nil, nil, IncompatibleAlgorithm
	}

	_, err = fmt.Sscanf(splitHash[2], "v=%d", &version)
	if err != nil || version != argon2.Version {
		return nil, nil, nil, IncompatibleVersion
	}

	_, err = fmt.Sscanf(splitHash[3], "m=%d,t=%d,p=%d", &tuning.kBmem, &tuning.iter, &tuning.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(splitHash[4])
	if err != nil {
		return nil, nil, nil, err
	}
	tuning.saltLen = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(splitHash[5])
	if err != nil {
		return nil, nil, nil, err
	}
	tuning.keyLen = uint32(len(hash))

	return tuning, salt, hash, nil
}

func buySomeSalt(length uint32) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	return salt, err
}

func getDefaultTuning() *argonTune {
	return &argonTune{
		kBmem:       uint32(64 * kB), //kb * kb * multiplier = MB * multiplier
		iter:        10,
		parallelism: 8,
		saltLen:     16,
		keyLen:      32,
	}
}
