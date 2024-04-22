package keyword

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/api/ads/dsp"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/ads/dsp/rtb/kuaiche/keyword"
)

// 快车腰带店铺关键词管理列表
type KuaicheKeywordShopListRequest struct {
	api.BaseRequest
	PageSize             int    `json:"pageSize,omitempty"`             // 每页条数
	ClickOrOrderDay      uint   `json:"clickOrOrderDay"`                // 当天:0;昨天：1;最近15天：15
	GiftFlag             uint   `json:"giftFlag"`                       // 包含赠品 0不含赠品 1含赠品
	CampaignId           uint64 `json:"campaignId,omitempty"`           // 计划id
	PutType              uint   `json:"putType,omitempty"`              // 投放类型 3商品 4活动
	EndDay               string `json:"endDay"`                         // 结束时间
	GroupId              uint64 `json:"groupId,omitempty"`              // 单元id
	StartDay             string `json:"startDay"`                       // 开始时间
	Page                 int    `json:"page"`                           // 页码
	Platform             string `json:"platform,omitempty"`             // 设备类型 空：全部 0 pc 1 无线
	ClickOrOrderCaliber  uint   `json:"clickOrOrderCaliber"`            // 点击口径：0 订单口径：1
	OrderStatusCategory  uint   `json:"orderStatusCategory,omitempty"`  // 订单类型 空：全部订单 1：成交订单
	NameLike             string `json:"nameLike,omitempty"`             // 计划名称
	AccessPin            string `json:"accessPin,omitempty"`            // 被免密访问的pin
	AuthType             string `json:"authType,omitempty"`             // 授权模式
	PlatformBusinessType string `json:"platformBusinessType,omitempty"` // 平台业务类型，DST_JZT：京准通
}

type KuaicheKeywordShopListResponse struct {
	Responce  *KuaicheKeywordShopListResponce `json:"jingdong_ads_dsp_rtb_shopkeyword_list_responce,omitempty" codec:"jingdong_ads_dsp_rtb_shopkeyword_list_responce,omitempty"`
	ErrorResp *api.ErrorResponnse             `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheKeywordShopListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordShopListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordShopListResponce struct {
	ReturnType *KuaicheKeywordShopListResponseData `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r KuaicheKeywordShopListResponce) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r KuaicheKeywordShopListResponce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type KuaicheKeywordShopListResponseData struct {
	Data *KuaicheKeywordShopListResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheKeywordShopListResponseDataData struct {
	Keywords  []dsp.KeywordData   `json:"datas,omitempty" codec:"datas,omitempty"`
	Ext       *dsp.KeywordExtData `json:"ext,omitempty" codec:"ext,omitempty"`
	Paginator *dsp.Paginator      `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheKeywordShopList(req *KuaicheKeywordShopListRequest) (*KuaicheKeywordShopListResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordShopListRequest()
	r.SetPageSize(req.PageSize)
	r.SetClickOrOrderDay(req.ClickOrOrderDay)
	r.SetGiftFlag(req.GiftFlag)
	r.SetEndDay(req.EndDay)
	r.SetStartDay(req.StartDay)
	r.SetPage(req.Page)
	r.SetClickOrOrderCaliber(req.ClickOrOrderCaliber)

	if req.CampaignId > 0 {
		r.SetCampaignId(req.CampaignId)
	}
	if req.GroupId > 0 {
		r.SetGroupId(req.GroupId)
	}
	if req.Platform != "" {
		r.SetPlatform(req.Platform)
	}
	if req.OrderStatusCategory > 0 {
		r.SetOrderStatusCategory(req.OrderStatusCategory)
	}
	if req.NameLike != "" {
		r.SetNameLike(req.NameLike)
	}
	if req.AccessPin != "" {
		r.SetAccessPin(req.AccessPin)
	}
	if req.AuthType != "" {
		r.SetAuthType(req.AuthType)
	}
	if req.PlatformBusinessType != "" {
		r.SetPlatformBusinessType(req.PlatformBusinessType)
	}

	var response KuaicheKeywordShopListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.ReturnType.Data, nil
}
