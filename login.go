package br

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
	log "github.com/zrb-channel/utils/logger"
)

// Login
// @param ctx
// @param conf
// @param body
// @date 2022-09-21 14:19:29
func Login(ctx context.Context, conf *Config, body *LoginRequest) (*LoginResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req, err := NewRequest(body, conf.ApyKey)
	if err != nil {
		return nil, err
	}

	var resp *resty.Response
	resp, err = Request(ctx, conf.Username, conf.Password).SetBody(req).Post(LoginAddr)
	if err != nil {
		log.WithError(err).Error("[百融]-获取联登地址，请求失败", log.Fields(map[string]any{"conf": conf, "body": body}))
		return nil, err
	}

	result := &BaseResponse{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[百融]-获取联登地址，返回数据解析失败", log.Fields(map[string]any{"conf": conf, "body": body, "resp": resp.String()}))
		return nil, err
	}

	if result.Code != 0 {
		log.Error("[百融]-获取联登地址，返回数据状态有误", log.Fields(map[string]any{"conf": conf, "body": body, "result": result, "resp": resp.String()}))
		return nil, errors.New(result.Message)
	}

	var v []byte
	v, err = Base64Decrypt(result.Data, []byte(conf.ApyKey))
	if err != nil {
		log.WithError(err).Error("[百融]-获取联登地址，数据解密失败", log.Fields(map[string]any{"conf": conf, "body": body, "result": result, "resp": resp.String()}))
		return nil, err
	}

	u := &LoginResponse{}
	if err = json.Unmarshal(v, u); err != nil {
		log.WithError(err).Error("[百融]-获取联登地址，结果解密失败", log.Fields(map[string]any{"conf": conf, "body": body, "result": result, "resp": resp.String(), "response": string(v)}))
		return nil, err
	}

	return u, nil
}
