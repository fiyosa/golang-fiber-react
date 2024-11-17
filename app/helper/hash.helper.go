package helper

import (
	"encoding/base64"
	"fmt"
	"go-fiber-react/config"

	"github.com/speps/go-hashids/v2"
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

func (*Hash) Create(data string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (*Hash) Verify(check string, original string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(original), []byte(check)); err != nil {
		return false
	}
	return true
}

func (*Hash) EncodeId(data int) (string, error) {
	h := setupHD()
	encode, err := h.Encode([]int{data})
	if err != nil {
		return "", err
	}
	return encode, err
}

func (*Hash) DecodeId(data string) (int, error) {
	h := setupHD()
	decode, err := h.DecodeWithError(data)
	if err != nil {
		return -1, err
	}
	return decode[0], err
}

func EncodeStr(data string) (string, error) {
	result := base64.StdEncoding.EncodeToString([]byte(data))
	return result, nil
}

func DecodeStr(encode string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return "", err
	}
	return string(result), nil
}

func setupHD() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = config.APP_SECRET
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	return h
}
