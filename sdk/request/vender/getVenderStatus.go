package vender

import "github.com/rise-worlds/jos/sdk"

type GetVenderStatusRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetVenderStatusRequest() (req *GetVenderStatusRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.vender.getVenderStatus", Params: make(map[string]interface{})}
	req = &GetVenderStatusRequest{
		Request: &request,
	}
	return
}
