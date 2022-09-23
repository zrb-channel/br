package br

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
)

// Query
// @param ctx
// @param conf
// @param body
// @date 2022-09-21 14:53:54
func Query(ctx context.Context, conf *Config, body *QueryRequest) (*QueryResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req, err := NewRequest(body, conf.ApyKey)
	if err != nil {
		return nil, errors.New("[百融]-创建查询请求失败")
	}

	var resp *resty.Response
	resp, err = Request(ctx, conf.Username, conf.Password).SetBody(req).Post(QueryAddr)

	if err != nil {
		return nil, errors.New("[百融]-查询请求失败")
	}

	result := &BaseResponse{}

	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, errors.New("[百融]-查询数据响应数据解析失败")
	}

	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}

	v, err := Base64Decrypt(result.Data, []byte(conf.ApyKey))
	if err != nil {
		return nil, errors.New("[百融]-查询响应数据解密失败")
	}
	data := &QueryResponse{}
	if err = json.Unmarshal(v, data); err != nil {
		return nil, errors.New("[百融]-查询响应数据解析为结果失败")
	}

	return data, nil
}
