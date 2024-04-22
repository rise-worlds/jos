package points

import "github.com/rise-worlds/jos/sdk"

type GetCouponInfoRequest struct {
	Request *sdk.Request
}

func NewGetCouponInfoRequest() (req *GetCouponInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.getCouponInfo", Params: nil}
	req = &GetCouponInfoRequest{
		Request: &request,
	}
	return
}
