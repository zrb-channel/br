package br

import (
	"context"
	"errors"
	json "github.com/json-iterator/go"
	"io"
	"net/http"
)

// Notify
// @param ctx
// @param req
// @date 2022-09-21 14:47:59
func Notify(ctx context.Context, req *http.Request) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	res := &NotifyRequest{}
	if err = json.Unmarshal(body, res); err != nil {
		return errors.New("参数有误")
	}

	return nil
}
