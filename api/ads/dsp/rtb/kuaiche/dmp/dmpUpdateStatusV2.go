package dmp

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/api/ads/dsp"
	"github.com/rise-worlds/jos/sdk"
	requestDsp "github.com/rise-worlds/jos/sdk/request/ads/dsp"
	"github.com/rise-worlds/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 启动/暂停人群
type KuaicheDmpUpdateStatusV2Request struct {
	api.BaseRequest
	Data   *dmp.KuaicheDmpUpdateStatusV2RequestData      `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheDmpUpdateStatusV2Response struct {
	Responce  *KuaicheDmpUpdateStatusV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_dmp_updateStatus_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_dmp_updateStatus_responce,omitempty"`
	ErrorResp *api.ErrorResponnse               `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpUpdateStatusV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpUpdateStatusV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpUpdateStatusV2Responce struct {
	Data *KuaicheDmpUpdateStatusV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDmpUpdateStatusV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheDmpUpdateStatusV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheDmpUpdateStatusV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheDmpUpdateStatusV2(req *KuaicheDmpUpdateStatusV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpUpdateStatusV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheDmpUpdateStatusV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Success, nil
}
