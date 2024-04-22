package order

import (
	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/union/order"
)

type UnionOrderQueryRequest struct {
	api.BaseRequest
	PageNo       uint   `json:"pageNo,omitempty"`       // 页码，返回第几页结果
	PageSize     uint   `json:"pageSize,omitempty"`     // 每页包含条数，上限为500
	Type         uint   `json:"type,omitempty"`         // 订单时间查询类型(1：下单时间，2：完成时间，3：更新时间)
	Time         string `json:"time,omitempty"`         // 查询时间，建议使用分钟级查询，格式：yyyyMMddHH、yyyyMMddHHmm或yyyyMMddHHmmss，如201811031212 的查询范围从12:12:00--12:12:59
	ChildUnionId string `json:"childUnionId,omitempty"` // 子站长ID（需要联系运营开通PID账户权限才能拿到数据），childUnionId和key不能同时传入
	Key          string `json:"key,omitempty"`          // 其他推客的授权key，查询工具商订单需要填写此项，childUnionid和key不能同时传入
}

type UnionOrderQueryResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty"`
	Data      *UnionOrderQueryResponseData `json:"jd_union_open_order_query_response,omitempty"`
}

func (r UnionOrderQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == ""
}

func (r UnionOrderQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type UnionOrderQueryResponseData struct {
	Result string `json:"result,omitempty"`
}

type UnionOrderQueryResult struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    []OrderResp `json:"data,omitempty"`
	HasMore bool        `json:"hasMore,omitempty"`
}

func (r UnionOrderQueryResult) IsError() bool {
	return r.Code != 200
}

func (r UnionOrderQueryResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

type OrderResp struct {
	FinishTime int64     `json:"finishTime,omitempty"`
	OrderTime  int64     `json:"orderTime,omitempty"`
	OrderEmt   uint      `json:"orderEmt,omitempty"`
	OrderId    uint64    `json:"orderId,omitempty"`
	ParentId   uint64    `json:"ParentId,omitempty"`
	Plus       uint      `json:"plus,omitempty"`
	PopId      uint64    `json:"popId,omitempty"`
	UnionId    uint64    `json:"unionId,omitempty"`
	Ext1       string    `json:"ext1,omitempty"`
	ValidCode  int       `json:"validCode,omitempty"`
	SkuList    []SkuInfo `json:"skuList"`
}

type SkuInfo struct {
	ActualCosPrice    float64 `json:"actualCosPrice,omitempty"`
	ActualFee         float64 `json:"actualFee,omitempty"`
	CommissionRate    float64 `json:"commissionRate,omitempty"`
	EstimateCosPrice  float64 `json:"estimateCosPrice,omitempty"`
	EstimateFee       float64 `json:"estimateFee,omitempty"`
	FinalRate         float64 `json:"finalRate,omitempty"`
	Cid1              uint64  `json:"cid1,omitempty"`
	Cid2              uint64  `json:"cid2,omitempty"`
	Cid3              uint64  `json:"cid3,omitempty"`
	FrozenSkuNum      uint    `json:"frozenSkuNum,omitempty"`
	Pid               string  `json:"pid,omitempty"`
	PositionId        uint64  `json:"positionId,omitempty"`
	Price             float64 `json:"price,omitempty"`
	SiteId            uint64  `json:"siteId,omitempty"`
	SkuId             uint64  `json:"skuId,omitempty"`
	SkuName           string  `json:"skuName,omitempty"`
	SkuNum            uint    `json:"skuNum,omitempty"`
	SkuReturnNum      uint    `json:"skuReturnNum,omitempty"`
	SubSideRate       float64 `json:"subSideRate,omitempty"`
	SubsidyRate       float64 `json:"subsidyRate,omitempty"`
	UnionAlias        string  `json:"unionAlias,omitempty"`
	UnionTag          string  `json:"unionTag,omitempty"`
	UnionTrafficGroup int     `json:"unionTrafficGroup,omitempty"`
	ValidCode         int     `json:"validCode,omitempty"`
	SubUnionId        string  `json:"subUnionId,omitempty"`
	TraceType         int     `json:"traceType,omitempty"`
	PopId             uint64  `json:"popId,omitempty"`
	Ext1              string  `json:"ext1,omitempty"`
	CpActId           uint64  `json:"cpActId,omitempty"`
	UnionRole         uint    `json:"unionRole,omitempty"`
}

// 订单查询接口
func UnionOrderQuery(req *UnionOrderQueryRequest) (bool, []OrderResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewUnionOrderQueryRequest()
	orderReq := &order.OrderReq{
		PageNo:       req.PageNo,
		PageSize:     req.PageSize,
		Type:         req.Type,
		Time:         req.Time,
		ChildUnionId: req.ChildUnionId,
		Key:          req.Key,
	}
	r.SetOrderReq(orderReq)

	var response UnionOrderQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, nil, err
	}
	var ret UnionOrderQueryResult
	if err := client.Logger().DecodeJSON([]byte(response.Data.Result), &ret); err != nil {
		return false, nil, err
	}

	if ret.IsError() {
		return false, nil, ret
	}

	return ret.HasMore, ret.Data, nil
}
