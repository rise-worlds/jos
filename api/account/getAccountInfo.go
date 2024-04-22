package account

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/account"
)

type GetAccountInfoRequest struct {
	api.BaseRequest

	AccountType uint8  `json:"accountType" codec:"accountType"`
	AccountCode string `json:"accountCode" codec:"accountCode"`
}

type GetAccountInfoResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetAccountInfoData `json:"jingdong_pop_account_getAccountInfo_responce,omitempty" codec:"jingdong_pop_account_getAccountInfo_responce,omitempty"`
}

func (r GetAccountInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetAccountInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no bean account info"
}

type GetAccountInfoData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	Result *AccountInfo `json:"beanAccount,omitempty" codec:"beanAccount,omitempty"`
}

func (r GetAccountInfoData) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r GetAccountInfoData) Error() string {
	if r.Code != "0" {
		return sdk.ErrorString(r.Code, r.ErrorDesc)
	}
	if r.Result == nil {
		return "no bean account info."
	}
	return "unexpected error"
}

func GetAccountInfo(req *GetAccountInfoRequest) (*AccountInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := account.NewGetAccountInfoRequest()
	r.SetAccountType(req.AccountType)
	r.SetAccountCode(req.AccountCode)

	var response GetAccountInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Data.Result, nil
}
