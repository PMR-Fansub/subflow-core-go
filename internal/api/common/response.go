package common

type APIResponse struct {
	Code      ResultCode  `json:"code"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type ResultCode int

const (
	ResultCodeSuccess             ResultCode = 1000
	ResultCodeFailed              ResultCode = 1001
	ResultCodeRegisterFailed      ResultCode = 1002
	ResultCodeFormInvalid         ResultCode = 1003
	ResultCodeLoginFailed         ResultCode = 1004
	ResultCodeNotLogin            ResultCode = 1005
	ResultCodeUserOperationFailed ResultCode = 1006
	ResultCodeQueryFailed         ResultCode = 1007
	ResultCodePermissionDenied    ResultCode = 1008
	ResultCodeNotFound            ResultCode = 4000
	ResultCodeNotSupported        ResultCode = 4001
	ResultCodeUnknown             ResultCode = 9999
)
