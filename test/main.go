package main

import (
	"context"
	"fmt"
	"github.com/zrb-channel/br"
)

func main() {
	ctx := context.Background()

	conf := &br.Config{
		ApyKey:   "AS6DMGhi99ER6PCc",
		Username: "yueshengshi",
		Password: "rbbny69zmyp0a",
	}

	query(ctx, conf)
}

// query
// @param ctx
// @param conf
// @date 2022-09-21 14:56:59
func query(ctx context.Context, conf *br.Config) {
	req := &br.QueryRequest{
		OrderNo:    "2022091421340aa",
		ChannelId:  121418,
		CreditCode: "91440605MA4X6KJM2N",
	}
	fmt.Println(br.Query(ctx, conf, req))
}

// create
// @param ctx
// @param conf
// @date 2022-09-21 14:56:57
func create(ctx context.Context, conf *br.Config) {

	req := &br.LoginRequest{
		Pid:         1001,
		OrderNo:     "2022091421340aa",
		ChannelId:   121418,
		CompanyName: "广东粤省事智能科技有限公司",
		CreditCode:  "91440605MA4X6KJM2N",
		ApplyName:   "刘国琼",
		ApplyId:     "440881198506294828",
		Mobile:      "13189664629",
		DDAStatus:   99,
	}

	fmt.Println(br.Login(ctx, conf, req))
}
