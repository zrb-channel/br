package br

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
)

// Login
// @param ctx
// @param conf
// @param body
// @date 2022-09-21 14:19:29
func Login(ctx context.Context, conf *Config, body *LoginRequest) (*LoginResponse, error) {

	req, err := NewRequest(body, conf.ApyKey)
	if err != nil {
		return nil, err
	}

	var resp *resty.Response
	resp, err = Request(ctx, conf.Username, conf.Password).SetBody(req).Post(LoginAddr)
	if err != nil {
		return nil, err
	}

	result := &BaseResponse{}

	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}

	var v []byte
	v, err = Base64Decrypt(result.Data, []byte(conf.ApyKey))
	if err != nil {
		return nil, err
	}

	u := &LoginResponse{}
	if err = json.Unmarshal(v, u); err != nil {
		return nil, err
	}

	return u, nil
}
