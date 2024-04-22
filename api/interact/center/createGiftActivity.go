package center

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/interact/center"
)

type CreateGiftActivityRequest struct {
	api.BaseRequest
	AppName      string `json:"appName,omitempty" codec:"appName,omitempty"` // 调用方应用名称，新接口接入必须联系产品，出现问题概不负责，且有权利追求责任及接口降级
	AppId        uint64 `json:"appId,omitempty" codec:"appId,omitempty"`
	Channel      uint8  `json:"channel,omitempty" codec:"channel,omitempty"`
	IsPrize      uint8  `json:"isPrize,omitempty" codec:"isPrize,omitempty"`
	PrizeType    uint8  `json:"prizeType,omitempty" codec:"prizeType,omitempty"`
	Source       uint8  `json:"source,omitempty" codec:"source,omitempty"`
	Type         uint8  `json:"type,omitempty" codec:"type,omitempty"`
	StartTime    string `json:"startTime,omitempty" codec:"startTime,omitempty"`
	EndTime      string `json:"endTime,omitempty" codec:"endTime,omitempty"`
	Creator      string `json:"creator,omitempty" codec:"creator,omitempty"`
	Name         string `json:"name,omitempty" codec:"name,omitempty"`
	SourceName   string `json:"sourceName,omitempty" codec:"sourceName,omitempty"`
	CouponId     uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`
	Discount     string `json:"discount,omitempty" codec:"discount,omitempty"`
	Quota        string `json:"quota,omitempty" codec:"quota,omitempty"`
	SendCount    string `json:"send_count"`
	PrizeEndTime string `json:"prize_end_time"`
	ValidateDay  string `json:"validate_day"`
}

type CreateGiftActivityResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CreateGiftActivityData `json:"jingdong_interact_center_api_service_write_createGiftActivity_responce,omitempty" codec:"jingdong_interact_center_api_service_write_createGiftActivity_responce,omitempty"`
}

func (r CreateGiftActivityResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CreateGiftActivityResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CreateGiftActivityData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *CreateGiftActivityResult `json:"giftActivityResults,omitempty" codec:"giftActivityResults,omitempty"`
}

func (r CreateGiftActivityData) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r CreateGiftActivityData) Error() string {
	if r.Code != "0" {
		return sdk.ErrorString(r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type CreateGiftActivityResult struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Code uint32 `json:"code" codec:"code"`
	Msg  string `json:"msg" codec:"msg"`
}

func CreateGiftActivity(req *CreateGiftActivityRequest) (*CreateGiftActivityResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCreateGiftActivityRequest()
	r.SetAppName(req.AppName)
	r.SetAppId(req.AppId)
	r.SetChannel(req.Channel)
	r.SetIsPrize(req.IsPrize)
	r.SetPrizeType(req.PrizeType)
	r.SetSource(req.Source)
	r.SetType(req.Type)
	r.SetStartTime(req.StartTime)
	r.SetEndTime(req.EndTime)
	r.SetCreator(req.Creator)
	r.SetName(req.Name)
	r.SetSourceName(req.SourceName)
	r.SetCouponId(req.CouponId)

	if len(req.Quota) > 0 {
		r.SetQuota(req.Quota)
	}

	if len(req.Discount) > 0 {
		r.SetDiscount(req.Discount)
	}

	if len(req.SendCount) > 0 {
		r.SetSendCount(req.SendCount)
	}

	if len(req.PrizeEndTime) > 0 {
		r.SetPrizeEndTime(req.PrizeEndTime)
	}
	if len(req.ValidateDay) > 0 {
		r.SetValidateDay(req.ValidateDay)
	}

	var response CreateGiftActivityResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
