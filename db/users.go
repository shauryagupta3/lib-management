package db

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"libraryManagement/types"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var defaultParams = &argonParams{
	memory:      12288,
	iterations:  2,
	parallelism: 1,
	saltLength:  16,
	keyLength:   32,
}

func CreateUser(u *types.User) error {

	if _, err := getUserPasswordByEmail(u); err == nil {
		return errors.New("user already exsists")
	}

	if err := generateHashFromPass(u); err != nil {
		return err
	}

	if _, err := dbConn.Exec(context.Background(), "insert into users(email,password) values($1,$2)", u.Email, u.Password); err != nil {
		return err
	}

	return nil
}

func AuthUser(u *types.User) error {
	passFromDB, err := getUserPasswordByEmail(u)
	if err != nil || passFromDB == "" {
		return err
	}
	passFromDBArr := strings.Split(passFromDB, " ")

	check, err := compareHashAndPass(u.Password, passFromDBArr[0], passFromDBArr[1])
	if err != nil {
		return err
	}
	if !check {
		return errors.New("wrong password")
	}
	return nil
}

func getUserPasswordByEmail(u *types.User) (string, error) {
	var pass string
	err := dbConn.QueryRow(context.Background(), "select password from users where email=$1", u.Email).Scan(&pass)
	if err != nil {
		return "", err
	}
	return pass, nil

}

func generateHashFromPass(u *types.User) error {
	salt, err := generateSalt(defaultParams.saltLength)
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(u.Password), salt, defaultParams.iterations, defaultParams.memory, defaultParams.parallelism, defaultParams.keyLength)

	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	// encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)
	encodedHash := fmt.Sprintf("%s %s", b64Hash, b64Salt)
	u.Password = encodedHash
	return nil
}

func generateSalt(i uint32) ([]byte, error) {
	randomBytes := make([]byte, i)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

func compareHashAndPass(pass string, hash string, salt string) (bool, error) {
	saltByte, err := base64.RawStdEncoding.Strict().DecodeString(salt)
	if err != nil {
		return false, err
	}

	hashByte, err := base64.RawStdEncoding.Strict().DecodeString(hash)
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(pass), saltByte, defaultParams.iterations, defaultParams.memory, defaultParams.parallelism, defaultParams.keyLength)

	if subtle.ConstantTimeCompare(newHash, hashByte) == 1 {
		return true, nil
	}

	return false, nil
}
