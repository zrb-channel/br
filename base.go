package br

import (
	"context"
	"encoding/base64"
	"github.com/forgoer/openssl"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
	"strings"
)

const (
	LoginAddr = BaseAddr + "/channel/getH5"

	QueryAddr = BaseAddr + "/channel/getResult"
)

// Request
// @param ctx
// @param username
// @param password
// @date 2022-09-21 13:44:50
func Request(ctx context.Context, username, password string) *resty.Request {

	timestamp := utils.Millisecond()

	nonce := strings.ReplaceAll(uuid.NewV4().String(), "-", "")

	sign := hash.MD5String(password + username + timestamp + nonce)

	headers := map[string]string{
		"Content-Type":   "application/json",
		"X-BR-Name":      username,
		"X-BR-Timestamp": timestamp,
		"X-BR-Nonce":     nonce,
		"X-BR-Signature": sign,
	}

	return utils.Request(ctx).SetHeaders(headers)
}

// NewRequest
// @param body
// @param apiKey
// @date 2022-09-21 13:44:53
func NewRequest[T any](body T, apiKey string) (*BaseRequest, error) {
	msgBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var msg string
	msg, err = EncryptToBase64(msgBytes, []byte(apiKey))

	return &BaseRequest{Data: msg}, nil
}

// EncryptToBase64
// @param data
// @param key
// @param iv
// @date 2022-09-21 13:56:19
func EncryptToBase64(data []byte, key []byte) (string, error) {
	msg, err := openssl.AesCBCEncrypt(data, key, key, openssl.PKCS5_PADDING)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(msg), nil
}

func Base64Decrypt(value string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}
	return openssl.AesCBCDecrypt(data, key, key, openssl.PKCS5_PADDING)
}
