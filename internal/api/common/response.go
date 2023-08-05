package common

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type Result struct {
	Success  bool
	Code     int
	HttpCode int
	Message  string
}

var (
	ResultSuccess             = Result{true, 1000, fiber.StatusOK, "操作成功"}
	ResultFailed              = Result{false, 1001, fiber.StatusBadRequest, "操作失败"}
	ResultCreateUserFailed    = Result{false, 1002, fiber.StatusUnprocessableEntity, "创建用户失败"}
	ResultFormInvalid         = Result{false, 1003, fiber.StatusUnprocessableEntity, "表单验证失败"}
	ResultUnauthorized        = Result{false, 1004, fiber.StatusUnauthorized, "身份验证失败"}
	ResultUserOperationFailed = Result{false, 1005, fiber.StatusBadRequest, "用户操作失败"}
	ResultQueryFailed         = Result{false, 1006, fiber.StatusNotFound, "查询失败"}
	ResultPermissionDenied    = Result{false, 1007, fiber.StatusForbidden, "权限不足"}
	ResultNotFound            = Result{false, 4000, fiber.StatusNotFound, "资源不存在"}
	ResultNotSupported        = Result{false, 4001, fiber.StatusMethodNotAllowed, "不支持此操作"}
	ResultInternalError       = Result{false, 5000, fiber.StatusInternalServerError, "内部错误"}
	ResultUnknown             = Result{false, 9999, fiber.StatusInternalServerError, "未知错误"}
)

func MakeAPIResponseWithMsg(result Result, data interface{}, msg string) *APIResponse {
	return &APIResponse{
		Code:      result.Code,
		Success:   result.Success,
		Message:   msg,
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
}

func MakeAPIResponse(result Result, data interface{}) *APIResponse {
	return MakeAPIResponseWithMsg(result, data, result.Message)
}
func MakeSuccessAPIResponse(data interface{}) *APIResponse {
	return MakeAPIResponse(ResultSuccess, data)
}
