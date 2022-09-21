package br

import (
	"github.com/shopspring/decimal"
)

type (
	Config struct {
		ApyKey   string
		Username string
		Password string
	}

	BaseRequest struct {
		Data string `json:"data"`
	}

	BaseResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
		Success bool   `json:"success"`
	}

	LoginRequest struct {
		Pid         int    `json:"pid"`
		OrderNo     string `json:"channelOrderNo"`
		ChannelId   int    `json:"channelId"`
		CompanyName string `json:"companyName,omitempty"`
		CreditCode  string `json:"creditCode,omitempty"`
		ApplyName   string `json:"applyName,omitempty"`
		ApplyId     string `json:"applyId,omitempty"`
		Mobile      string `json:"phone,omitempty"`
		DDAStatus   int    `json:"DDAStatus,omitempty"`
	}

	LoginResponse struct {
		Url string `json:"url"`
	}

	Order struct {
		LoanNo     string          `json:"loanNo"`
		LoanStatus int             `json:"loanStatus"`
		LoanTerm   int             `json:"loanTerm"`
		LoantLimit decimal.Decimal `json:"loantLimit"`
	}

	OrderInfo struct {
		OrderNo       string          `json:"channelOrderNo"`
		Status        int8            `json:"applyStatus"`
		ApplyNo       *string         `json:"applyNo,omitempty"`
		PreCreditTerm decimal.Decimal `json:"preCreditTerm,omitempty"`
		CreditLimit   decimal.Decimal `json:"creditLimit,omitempty"`
		RejectReason  string          `json:"rejectReason,omitempty"`
		ApplyTerm     int             `json:"applyTerm"`
		List          []*Order        `json:"loanInfo"`
	}

	NotifyRequest OrderInfo

	QueryResponse OrderInfo

	QueryRequest struct {
		OrderNo    string `json:"channelOrderNo"`
		ChannelId  int    `json:"channelId"`
		CreditCode string `json:"creditCode"`
	}
)
