package sku

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/api/ware"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/sku"
)

type SearchSkuListRequest struct {
	api.BaseRequest
	WareId            []uint64 `json:"wareId,omitempty" codec:"wareId,omitempty"`                       // 商品ID
	SkuId             []uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`                         // SKU ID
	SkuStatuValue     []int    `json:"skuStatuValue,omitempty" codec:"skuStatuValue,omitempty"`         // SKU状态：1:上架 2:下架 4:删除
	MaxStockNum       int      `json:"maxStockNum,omitempty" codec:"maxStockNum,omitempty"`             // 库存范围
	MinStockNum       int      `json:"minStockNum,omitempty" codec:"minStockNum,omitempty"`             // 库存范围
	EndCreatedTime    int64    `json:"endCreatedTime,omitempty" codec:"endCreatedTime,omitempty"`       // 创建时间范围
	EndModifiedTime   int64    `json:"endModifiedTime,omitempty" codec:"endModifiedTime,omitempty"`     // 创建时间范围
	StartCreatedTime  int64    `json:"startCreatedTime,omitempty" codec:"startCreatedTime,omitempty"`   // 修改时间范围
	StartModifiedTime int64    `json:"startModifiedTime,omitempty" codec:"startModifiedTime,omitempty"` // 修改时间范围
	OutId             []string `json:"outId,omitempty" codec:"outId,omitempty"`                         // 外部ID
	ColType           []string `json:"colType,omitempty" codec:"colType,omitempty"`                     // 合作类型
	ItemNum           string   `json:"itemNum,omitempty" codec:"itemNum,omitempty"`                     // 货号
	WareTitle         string   `json:"wareTitle,omitempty" codec:"wareTitle,omitempty"`                 // 商品名称
	OrderField        []string `json:"orderField,omitempty" codec:"orderField,omitempty"`               // 排序字段.目前支持skuId、stockNum
	OrderType         []string `json:"orderType,omitempty" codec:"orderType,omitempty"`                 // 排序类型：asc、desc
	PageNo            int      `json:"pageNo,omitempty" codec:"pageNo,omitempty"`                       // 页码
	PageSize          int      `json:"page_size,omitempty" codec:"page_size,omitempty"`                 // 每页条数
	Field             string   `json:"field,omitempty" codec:"field,omitempty"`                         // 自定义返回字段
}

type SearchSkuListResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SearchSkuListSubResponse `json:"jingdong_sku_read_searchSkuList_responce,omitempty" codec:"jingdong_sku_read_searchSkuList_responce,omitempty"`
}

func (r SearchSkuListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r SearchSkuListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type SearchSkuListSubResponse struct {
	Code      string             `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string             `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Page      *SearchSkuListPage `json:"page,omitempty" codec:"page,omitempty"`
}

func (r SearchSkuListSubResponse) IsError() bool {
	return r.Code != "0" || r.Page == nil
}

func (r SearchSkuListSubResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type SearchSkuListPage struct {
	Data      []ware.Sku `json:"data,omitempty" codec:"data,omitempty"`
	PageNo    int        `json:"pageNo,omitempty" codec:"pageNo,omitempty"`
	PageSize  int        `json:"pageSize,omitempty" codec:"pageSize,omitempty"`
	TotalItem int        `json:"totalItem,omitempty" codec:"totalItem,omitempty"`
}

// 根据条件检索订单信息 （仅适用于SOP、LBP，SOPL类型，FBP类型请调取FBP订单检索 ）
func SearchSkuList(req *SearchSkuListRequest) (*SearchSkuListPage, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := sku.NewSearchSkuListRequest()
	if len(req.WareId) > 0 {
		r.SetWareId(req.WareId)
	}
	if len(req.SkuId) > 0 {
		r.SetSkuId(req.SkuId)
	}
	if req.Field != "" {
		r.SetField(req.Field)
	}
	if req.PageNo > 0 {
		r.SetPageNo(req.PageNo)
	}
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	}
	if len(req.OrderField) > 0 {
		r.SetOrderField(req.OrderField)
	}
	if len(req.OrderType) > 0 {
		r.SetOrderType(req.OrderType)
	}
	if len(req.SkuStatuValue) > 0 {
		r.SetSkuStatusValue(req.SkuStatuValue)
	}
	if len(req.WareTitle) > 0 {
		r.SetWareTitle(req.WareTitle)
	}

	var response SearchSkuListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Page, nil
}
