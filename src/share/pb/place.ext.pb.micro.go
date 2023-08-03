// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: place.ext.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PlaceServiceExt service

func NewPlaceServiceExtEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PlaceServiceExt service

type PlaceServiceExtService interface {
	// 获取有电影院的地点
	HotCitiesByCinema(ctx context.Context, in *HotCitiesByCinemaReq, opts ...client.CallOption) (*HotCitiesByCinemaRep, error)
}

type placeServiceExtService struct {
	c    client.Client
	name string
}

func NewPlaceServiceExtService(name string, c client.Client) PlaceServiceExtService {
	return &placeServiceExtService{
		c:    c,
		name: name,
	}
}

func (c *placeServiceExtService) HotCitiesByCinema(ctx context.Context, in *HotCitiesByCinemaReq, opts ...client.CallOption) (*HotCitiesByCinemaRep, error) {
	req := c.c.NewRequest(c.name, "PlaceServiceExt.HotCitiesByCinema", in)
	out := new(HotCitiesByCinemaRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlaceServiceExt service

type PlaceServiceExtHandler interface {
	// 获取有电影院的地点
	HotCitiesByCinema(context.Context, *HotCitiesByCinemaReq, *HotCitiesByCinemaRep) error
}

func RegisterPlaceServiceExtHandler(s server.Server, hdlr PlaceServiceExtHandler, opts ...server.HandlerOption) error {
	type placeServiceExt interface {
		HotCitiesByCinema(ctx context.Context, in *HotCitiesByCinemaReq, out *HotCitiesByCinemaRep) error
	}
	type PlaceServiceExt struct {
		placeServiceExt
	}
	h := &placeServiceExtHandler{hdlr}
	return s.Handle(s.NewHandler(&PlaceServiceExt{h}, opts...))
}

type placeServiceExtHandler struct {
	PlaceServiceExtHandler
}

func (h *placeServiceExtHandler) HotCitiesByCinema(ctx context.Context, in *HotCitiesByCinemaReq, out *HotCitiesByCinemaRep) error {
	return h.PlaceServiceExtHandler.HotCitiesByCinema(ctx, in, out)
}
