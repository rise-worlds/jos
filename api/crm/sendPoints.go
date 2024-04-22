package crm

import (
	"strconv"
	"time"

	"github.com/rise-worlds/jos/api"
	"github.com/rise-worlds/jos/sdk"
	"github.com/rise-worlds/jos/sdk/request/crm"
)

const (
	POINTS_SOURCE_TYPE_EXCHANGE_REAL   = 24 // 实物兑换扣减积分
	POINTS_SOURCE_TYPE_EVALUATION      = 21 // 评价送积分
	POINTS_SOURCE_TYPE_BELOW_LINE      = 16 // 线下积分变动
	POINTS_SOURCE_TYPE_SHOP_GIFT       = 15 // 店铺礼包
	POINTS_SOURCE_TYPE_SHOP_SEND       = 14 // 商家发积分
	POINTS_SOURCE_TYPE_GAME            = 13 // 游戏积分
	POINTS_SOURCE_TYPE_SIGNIN          = 10 // 签到送积分
	POINTS_SOURCE_TYPE_EXCHANGE_COUPON = 5  // 积分兑换优惠券
	POINTS_SOURCE_TYPE_FOLLOW          = 11 // 关注送积分
)

type SendPointsRequest struct {
	api.BaseRequest

	SysName     string `json:"sys_name,omitempty" codec:"sys_name,omitempty"`         // 系统名称
	CustomerPin string `json:"customer_pin,omitempty" codec:"customer_pin,omitempty"` // 用户pin
	ResId       string `json:"res_id,omitempty" codec:"res_id,omitempty"`             // 来源标识
	SourceType  int    `json:"source_type,omitempty" codec:"source_type,omitempty"`   // 来源类型
	Points      int64  `json:"points,omitempty" codec:"points,omitempty"`             // 积分数
}

type SendPointsResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SendPointsData     `json:"jingdong_pop_crm_sendPoints_responce,omitempty" codec:"jingdong_pop_crm_sendPoints_responce,omitempty"`
}

func (r SendPointsResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r SendPointsResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type SendPointsData struct {
	Code   string `json:"code,omitempty" codec:"code,omitempty"`
	Result int64  `json:"sendpoints_result,omitempty" codec:"sendpoints_result,omitempty"`
}

func SendPoints(req *SendPointsRequest) (int64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewSendPointsRequest()
	r.SetCustomerPin(req.CustomerPin)
	if req.SysName != "" {
		r.SetSysName(req.SysName)
	} else {
		r.SetSysName("jdvip")
	}
	if req.ResId != "" {
		r.SetResId(req.ResId)
	} else {
		r.SetResId(sdk.StringsJoin("jdvip", time.Now().Format("20060102150405")))
	}
	r.SetSourceType(strconv.Itoa(req.SourceType))
	r.SetPoints(req.Points)

	var response SendPointsResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result, nil
}

func GetPointsSourceTypeDesc(sourceType int) string {
	switch sourceType {
	case POINTS_SOURCE_TYPE_EXCHANGE_REAL:
		return "实物兑换扣减积分"
	case POINTS_SOURCE_TYPE_EVALUATION:
		return "评价送积分"
	case POINTS_SOURCE_TYPE_BELOW_LINE:
		return "线下积分变动"
	case POINTS_SOURCE_TYPE_SHOP_GIFT:
		return "店铺礼包"
	case POINTS_SOURCE_TYPE_SHOP_SEND:
		return "商家发积分"
	case POINTS_SOURCE_TYPE_GAME:
		return "游戏积分"
	case POINTS_SOURCE_TYPE_SIGNIN:
		return "签到送积分"
	case POINTS_SOURCE_TYPE_EXCHANGE_COUPON:
		return "积分兑换优惠券"
	case POINTS_SOURCE_TYPE_FOLLOW:
		return "关注送积分"
	}

	return ""
}
