// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ShopInterfaceHTTPServer interface {
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
}

func RegisterShopInterfaceHTTPServer(s *http.Server, srv ShopInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/shop/{id}", _ShopInterface_GetGoods0_HTTP_Handler(srv))
}

func _ShopInterface_GetGoods0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/GetGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGoods(ctx, req.(*GetGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGoodsReply)
		return ctx.Result(200, reply)
	}
}

type ShopInterfaceHTTPClient interface {
	GetGoods(ctx context.Context, req *GetGoodsRequest, opts ...http.CallOption) (rsp *GetGoodsReply, err error)
}

type ShopInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewShopInterfaceHTTPClient(client *http.Client) ShopInterfaceHTTPClient {
	return &ShopInterfaceHTTPClientImpl{client}
}

func (c *ShopInterfaceHTTPClientImpl) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...http.CallOption) (*GetGoodsReply, error) {
	var out GetGoodsReply
	pattern := "/v1/shop/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/GetGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
