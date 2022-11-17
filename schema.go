package br

import (
	"github.com/shopspring/decimal"
)

// https://docs.qq.com/doc/DTnhNTEFpYXNncnFH

type (
	// Config 百融配置参数
	Config struct {
		// ApyKey 加密/解密AES key
		ApyKey string

		// Username 账号
		Username string

		// Password 密码,用于签名
		Password string

		ChannelId string
	}

	// BaseRequest 公共请求参数
	BaseRequest struct {
		// Data 业务参数json字符串aes加密后的内容，详细字段见各接口描述
		Data string `json:"data"`
	}

	// BaseResponse 公共返回数据
	BaseResponse struct {
		// Code 成功:0
		//500	服务异常
		//10003	参数缺失
		//10004	参数不合规
		//20001	签名验证失败
		//60001	解密错误
		//60002	接口次数超限
		//60003	接口未开通
		Code int `json:"code"`

		// Message 描述，可填写失败原因或拒绝原因
		Message string `json:"message"`

		// Data 需要aes解密，解密后为各接口业务响应参数json
		Data string `json:"data"`

		// 是否返回成功
		Success bool `json:"success"`
	}

	// LoginRequest 联登获取H5地址请求参数
	// 每次获取H5地址如果同一客户申请同一产品，
	// 借款结果状态为终态即applyStatus=4、7、9、10，11，31，需要更换channelOrderNo申请
	// 其余状态为非终态，同一客户申请同一产品channelOrderNo不可变
	LoginRequest struct {
		// Pid 平台指定
		Pid int `json:"pid"`

		// OrderNo 订单编号 32位以内
		OrderNo string `json:"channelOrderNo"`

		// 平台指定，每个渠道唯一，对应平台营销经理id
		ChannelId int `json:"channelId"`

		// CompanyName 企业名称
		CompanyName string `json:"companyName,omitempty"`

		// CreditCode 统一信用代码
		CreditCode string `json:"creditCode,omitempty"`

		// ApplyName 申请人姓名
		ApplyName string `json:"applyName,omitempty"`

		// ApplyId 申请人身份证号
		ApplyId string `json:"applyId,omitempty"`

		// Mobile 手机号
		Mobile string `json:"phone,omitempty"`

		// DDAStatus 数据认证情况
		//0：未认证任何数据
		//1：发票数据已认证
		//2：税务数据已认证
		//99：自动查验
		DDAStatus int `json:"DDAStatus,omitempty"`
	}

	// LoginResponse 联登获取H5地址响应参数
	LoginResponse struct {
		// Url 跳转H5地址
		Url string `json:"url"`
	}

	// Order 借款详情
	Order struct {
		// LoanNo 平台放款单号
		LoanNo string `json:"loanNo"`

		// LoanStatus 状态
		// 21：放款中
		//22：放款失败
		//23：放款成功，使用中
		//24：订单完结
		LoanStatus int `json:"loanStatus"`

		// LoanDate 放款时间 yyyy-MM-dd hh:mm:ss
		LoanDate string `json:"loanDate"`

		// LoanTerm 借款期限（单位月）
		LoanTerm int64 `json:"loanTerm"`

		// LoantLimit 放款金额（单位元）
		LoantLimit decimal.Decimal `json:"loantLimit"`
	}

	// OrderInfo 查询结果
	OrderInfo struct {
		// OrderNo 订单编号
		OrderNo string `json:"channelOrderNo"`

		// Status 授信申请结果
		// 1：申请中
		//2：初审中，用户提交申请，初步审核或预授信审核
		//3：初审通过，或预授信通过
		//4：初审拒绝，或预授信被拒
		//5：终审中
		//6：终审通过，或授信通过
		//7：终审拒绝
		//8：额度冻结
		//9：额度失效
		//10：订单取消
		//11：订单完结
		//31：前筛拒绝
		//33：申请失败
		//99：系统异常
		Status int8 `json:"applyStatus"`

		// ApplyNo 平台贷款申请单号
		ApplyNo *string `json:"applyNo,omitempty"`

		// PreCreditTerm 预授信金额（单位元）
		PreCreditTerm decimal.Decimal `json:"preCreditTerm,omitempty"`

		// CreditLimit 审批金额（单位元）
		CreditLimit decimal.Decimal `json:"creditLimit,omitempty"`

		// RejectReason 拒绝原因
		RejectReason string `json:"rejectReason,omitempty"`

		// ApplyTerm 申请期限（单位月）
		ApplyTerm int `json:"applyTerm"`

		// List 借款详情
		List []*Order `json:"loanInfo"`
	}

	NotifyRequest OrderInfo

	// QueryResponse 借款结果查询响应参数
	QueryResponse OrderInfo

	// QueryRequest 借款结果查询请求参数
	QueryRequest struct {
		// OrderNo 订单编号
		OrderNo string `json:"channelOrderNo"`

		// ChannelId 平台指定，每个渠道唯一，对应平台营销经理id
		ChannelId int `json:"channelId"`

		// CreditCode 统一信用代码
		CreditCode string `json:"creditCode"`
	}
)
