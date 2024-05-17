package constant

import (
	"errors"

	"gorm.io/gorm"
)

type ErrCode struct {
	ErrCode int
	ErrMsg  string
}

func (r ErrCode) Error() string {
	return r.ErrMsg
}

func NewErr(errCode int, errMsg string) ErrCode {
	return ErrCode{ErrCode: errCode, ErrMsg: errMsg}
}

func NewCustomErr(errMsg string) ErrCode {
	return ErrCode{ErrCode: 10008, ErrMsg: errMsg}
}

var (
	ErrInternal       = NewErr(99999, "服务器繁忙")
	ErrUnauthorized   = NewErr(10000, "登录失效")
	ErrForbidden      = NewErr(10001, "无权限")
	ErrRPC            = NewErr(10002, "远程调用失败")
	ErrRecordNotFound = NewErr(10003, "数据不存在")
	ErrInvalidParams  = NewErr(10004, "参数错误")
	ErrInvalidCaptcha = NewErr(10005, "验证码错误")
	ErrCacheNil       = NewErr(10006, "缓存不存在")
	ErrUpdateFailed   = NewErr(10007, "更新失败")
)

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrRecordNotFound)
}
