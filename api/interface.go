package api

import (
	"strconv"

	"github.com/rise-worlds/jos/sdk"
)

// ApiKey jd APP key/secret
type ApiKey struct {
	Key    string
	Secret string
	Id     uint8  `json:"id,omitempty" codec:"id,omitempty"`
	Name   string `json:"name,omitempty" codec:"name,omitempty"`
}

type BaseRequest struct {
	Session  string
	AnApiKey *ApiKey `json:",omitempty" codec:",omitempty"`

	Debug bool `json:"-"`
}

type ErrorResponnse struct {
	Code   string `json:"code,omitempty" codec:"code,omitempty"`
	ZhDesc string `json:"zh_desc,omitempty" codec:"zh_desc,omitempty"`
	EnDesc string `json:"en_desc,omitempty" codec:"en_desc,omitempty"`
}

func (e ErrorResponnse) Error() string {
	return sdk.StringsJoin("Code:", e.Code, ", ZhDesc:", e.ZhDesc, ", EnDesc:", e.EnDesc)
}

type ApiResult struct {
	Success        bool   `json:"success,omitempty" codec:"success,omitempty"`
	EnglishErrCode string `json:"englishErrCode,omitempty" codec:"englishErrCode,omitempty"`
	ChineseErrCode string `json:"chineseErrCode,omitempty" codec:"chineseErrCode,omitempty"`
	NumberCode     int    `json:"numberCode,omitempty" codec:"numberCode,omitempty"`
}

func (e ApiResult) IsError() bool {
	return !e.Success
}

func (e ApiResult) Error() string {
	return sdk.StringsJoin("EnglishErrCode:", e.EnglishErrCode, ", ChineseErrCode:", e.ChineseErrCode, ", NumberCode:", strconv.Itoa(e.NumberCode))
}
