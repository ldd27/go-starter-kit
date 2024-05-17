package types

import (
	"github.com/ldd27/go-starter-kit/internal/constant"
)

type PageReq struct {
	PageIndex int `json:"page_index" query:"page_index" validate:"required"`
	PageSize  int `json:"page_size" query:"page_size" validate:"required"`
}

type PageRes struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

type CursorPageReq struct {
	Cursor int `json:"cursor" query:"cursor" validate:"min=0"`
}

type CursorPageRes struct {
	Data       interface{} `json:"data"`
	NextCursor int         `json:"next_cursor"`
}

type Res struct {
	Success    bool        `json:"success"`
	ErrCode    int         `json:"err_code"`
	ErrMsg     string      `json:"err_msg"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"-"`
}

func NewErrResWithStatusCode(statusCode, errCode int, errMsg string) Res {
	return Res{StatusCode: statusCode, ErrCode: errCode, ErrMsg: errMsg}
}

func NewErrResWithErrCode(statusCode int, errCode constant.ErrCode) Res {
	return Res{StatusCode: statusCode, ErrCode: errCode.ErrCode, ErrMsg: errCode.ErrMsg}
}
