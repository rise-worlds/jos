package goods

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/union/goods"
)

type BidFieldQueryRequest struct {
	api.BaseRequest
	SkuIds []uint64 `json:"skuIds"` // skuId集合
	Fields []string `json:"fields"` // 查询域集合，不填写则查询全部 ('categoryInfo','imageInfo','baseBigFieldInfo','bookBigFieldInfo','videoBigFieldInfo')
}

type BidFieldQueryResponse struct {
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty"`
	Data      *BidFieldQueryResponseData `json:"jd_union_open_goods_bigfield_query_responce,omitempty"`
}

func (r BidFieldQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == ""
}

func (r BidFieldQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type BidFieldQueryResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

type BidFieldQueryResult struct {
	Code    int64               `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    []BidFieldQueryResp `json:"data,omitempty"`
}

func (r BidFieldQueryResult) IsError() bool {
	return r.Code != 200
}

func (r BidFieldQueryResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

type BidFieldQueryResp struct {
	SkuId        int64              `json:"skuId,omitempty"`             // skuId
	SkuName      string             `json:"skuName,omitempty"`           // 商品名称
	MainSkuId    int64              `json:"mainSkuId,omitempty"`         // 自营主skuId
	ProductId    int64              `json:"productId,omitempty"`         // 非自营商品Id
	SkuStatus    int                `json:"skuStatus"`                   // sku上下架状态 1：上架(可搜索，可购买)， 0：下架(可通过skuid搜索，不可购买)， 2：可上架（可通过skuid搜索，不可购买）， 10：pop 删除（不可搜索，不可购买））
	Owner        string             `json:"owner"`                       // g=自营，p=pop
	DetailImages string             `json:"detailImages,omitempty"`      // 商详图
	ItemId       string             `json:"itemId,omitempty"`            // 联盟商品ID
	Category     *CategoryInfo      `json:"categoryInfo,omitempty"`      // 目录信息
	Image        *ImageInfo         `json:"imageInfo,omitempty"`         // 图片信心
	Base         *BaseBigFieldInfo  `json:"baseBigFieldInfo,omitempty"`  // 基础大字段信息
	Book         *BookBigFieldInfo  `json:"bookBigFieldInfo,omitempty"`  // 图书大字段信息
	Video        *VideoBigFieldInfo `json:"videoBigFieldInfo,omitempty"` // 影音大字段信息
}

type CategoryInfo struct {
	Cid1     uint64 `json:"cid1"`     // 一级类目ID
	Cid1Name string `json:"cid1Name"` // 一级类目名称
	Cid2     uint64 `json:"cid2"`     // 二级类目ID
	Cid2Name string `json:"cid2Name"` // 二级类目名称
	Cid3     uint64 `json:"cid3"`     // 三级类目ID
	Cid3Name string `json:"cid3Name"` // 三级类目名称
}

type ImageInfo struct {
	List []Url `json:"imageList,omitempty"` // 图片合集
}

type Url struct {
	Url string `json:"url"` // 图片链接地址，第一个图片链接为主图链接
}

type BaseBigFieldInfo struct {
	Wdis       string `json:"wdis,omitempty"`       // 商品介绍
	PropCode   string `json:"propCode,omitempty"`   // 规格参数
	WareQD     string `json:"wareQD,omitempty"`     // 包装清单(仅自营商品)
	PropGroups string `json:"propGroups,omitempty"` // 规格参数
}

type BookBigFieldInfo struct {
	Comments        string `json:"comments,omitempty"`        // 媒体评论
	Image           string `json:"image,omitempty"`           // 精彩文摘与插图(插图)
	ContentDesc     string `json:"contentDesc,omitempty"`     // 内容摘要(内容简介)
	RelatedProducts string `json:"relatedProducts,omitempty"` // 产品描述(相关商品)
	EditerDesc      string `json:"editerDesc,omitempty"`      // 编辑推荐
	Catalogue       string `json:"catalogue,omitempty"`       // 目录
	BookAbstract    string `json:"bookAbstract,omitempty"`    // 精彩摘要(精彩书摘)
	AuthorDesc      string `json:"authorDesc,omitempty"`      // 作者简介
	Introduction    string `json:"introduction,omitempty"`    // 前言(前言/序言)
	ProductFeatures string `json:"productFeatures,omitempty"` // 产品特色
}

type VideoBigFieldInfo struct {
	Comments            string `json:"comments,omitempty"`             // 媒体评论
	Image               string `json:"image,omitempty"`                // 精彩文摘与插图(插图)
	ContentDesc         string `json:"contentDesc,omitempty"`          // 内容摘要(内容简介)
	RelatedProducts     string `json:"relatedProducts,omitempty"`      // 产品描述(相关商品)
	EditerDesc          string `json:"editerDesc,omitempty"`           // 编辑推荐
	Catalogue           string `json:"catalogue,omitempty"`            // 目录
	BoxContents         string `json:"box_Contents,omitempty"`         // 包装清单
	MaterialDescription string `json:"material_Description,omitempty"` // 特殊说明
	Manual              string `json:"manual,omitempty"`               // 说明书
	ProductFeatures     string `json:"productFeatures,omitempty"`      // 产品特色
}

// 大字段商品查询接口
func BidFieldQuery(req *BidFieldQueryRequest) ([]BidFieldQueryResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewBigFieldQueryRequest()
	goodsReq := &goods.BigFieldGoodsReq{
		SkuIds: req.SkuIds,
		Fields: req.Fields,
	}
	r.SetGoodsReq(goodsReq)

	var response BidFieldQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	var ret BidFieldQueryResult
	if err := client.Logger().DecodeJSON([]byte(response.Data.Result), &ret); err != nil {
		return nil, err
	}
	if ret.IsError() {
		return nil, ret
	}
	return ret.Data, nil
}
