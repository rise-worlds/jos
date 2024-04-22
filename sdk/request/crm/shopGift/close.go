package crm

import (
	"github.com/rise-worlds/jos/sdk"
)

type ShopGiftCloseRequest struct {
	Request *sdk.Request
}

// create new request
func NewShopGiftCloseRequest() (req *ShopGiftCloseRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.shopGift.closeShopGiftCallBack", Params: make(map[string]interface{}, 1)}
	req = &ShopGiftCloseRequest{
		Request: &request,
	}
	return
}

func (req *ShopGiftCloseRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *ShopGiftCloseRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}
