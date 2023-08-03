// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: film.ext.proto

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

// Api Endpoints for FilmServiceExt service

func NewFilmServiceExtEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FilmServiceExt service

type FilmServiceExtService interface {
	// 获取正在售票的影片
	HotPlayMovies(ctx context.Context, in *HotPlayMoviesReq, opts ...client.CallOption) (*HotPlayMoviesRep, error)
	// 电影详情
	MovieDetail(ctx context.Context, in *MovieDetailReq, opts ...client.CallOption) (*MovieDetailRep, error)
	// 获取导演与演员信息
	MovieCreditsWithTypes(ctx context.Context, in *MovieCreditsWithTypesReq, opts ...client.CallOption) (*MovieCreditsWithTypesRep, error)
	// 获取海报信息
	ImageAll(ctx context.Context, in *ImageAllReq, opts ...client.CallOption) (*ImageAllRep, error)
	// 正在热映的影片
	LocationMovies(ctx context.Context, in *LocationMoviesReq, opts ...client.CallOption) (*LocationMoviesRep, error)
	// 即将上映的影片
	MovieComingNew(ctx context.Context, in *MovieComingNewReq, opts ...client.CallOption) (*MovieComingNewRep, error)
	// 图片搜索
	Search(ctx context.Context, in *SearchReq, opts ...client.CallOption) (*SearchRep, error)
	// 根据影院id和时间获取正在销售的影片信息
	GetFilmsByCidADay(ctx context.Context, in *GetFilmsByCidADayReq, opts ...client.CallOption) (*GetFilmsByCidADayRsp, error)
}

type filmServiceExtService struct {
	c    client.Client
	name string
}

func NewFilmServiceExtService(name string, c client.Client) FilmServiceExtService {
	return &filmServiceExtService{
		c:    c,
		name: name,
	}
}

func (c *filmServiceExtService) HotPlayMovies(ctx context.Context, in *HotPlayMoviesReq, opts ...client.CallOption) (*HotPlayMoviesRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.HotPlayMovies", in)
	out := new(HotPlayMoviesRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) MovieDetail(ctx context.Context, in *MovieDetailReq, opts ...client.CallOption) (*MovieDetailRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.MovieDetail", in)
	out := new(MovieDetailRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) MovieCreditsWithTypes(ctx context.Context, in *MovieCreditsWithTypesReq, opts ...client.CallOption) (*MovieCreditsWithTypesRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.MovieCreditsWithTypes", in)
	out := new(MovieCreditsWithTypesRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) ImageAll(ctx context.Context, in *ImageAllReq, opts ...client.CallOption) (*ImageAllRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.ImageAll", in)
	out := new(ImageAllRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) LocationMovies(ctx context.Context, in *LocationMoviesReq, opts ...client.CallOption) (*LocationMoviesRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.LocationMovies", in)
	out := new(LocationMoviesRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) MovieComingNew(ctx context.Context, in *MovieComingNewReq, opts ...client.CallOption) (*MovieComingNewRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.MovieComingNew", in)
	out := new(MovieComingNewRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) Search(ctx context.Context, in *SearchReq, opts ...client.CallOption) (*SearchRep, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.Search", in)
	out := new(SearchRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmServiceExtService) GetFilmsByCidADay(ctx context.Context, in *GetFilmsByCidADayReq, opts ...client.CallOption) (*GetFilmsByCidADayRsp, error) {
	req := c.c.NewRequest(c.name, "FilmServiceExt.GetFilmsByCidADay", in)
	out := new(GetFilmsByCidADayRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FilmServiceExt service

type FilmServiceExtHandler interface {
	// 获取正在售票的影片
	HotPlayMovies(context.Context, *HotPlayMoviesReq, *HotPlayMoviesRep) error
	// 电影详情
	MovieDetail(context.Context, *MovieDetailReq, *MovieDetailRep) error
	// 获取导演与演员信息
	MovieCreditsWithTypes(context.Context, *MovieCreditsWithTypesReq, *MovieCreditsWithTypesRep) error
	// 获取海报信息
	ImageAll(context.Context, *ImageAllReq, *ImageAllRep) error
	// 正在热映的影片
	LocationMovies(context.Context, *LocationMoviesReq, *LocationMoviesRep) error
	// 即将上映的影片
	MovieComingNew(context.Context, *MovieComingNewReq, *MovieComingNewRep) error
	// 图片搜索
	Search(context.Context, *SearchReq, *SearchRep) error
	// 根据影院id和时间获取正在销售的影片信息
	GetFilmsByCidADay(context.Context, *GetFilmsByCidADayReq, *GetFilmsByCidADayRsp) error
}

func RegisterFilmServiceExtHandler(s server.Server, hdlr FilmServiceExtHandler, opts ...server.HandlerOption) error {
	type filmServiceExt interface {
		HotPlayMovies(ctx context.Context, in *HotPlayMoviesReq, out *HotPlayMoviesRep) error
		MovieDetail(ctx context.Context, in *MovieDetailReq, out *MovieDetailRep) error
		MovieCreditsWithTypes(ctx context.Context, in *MovieCreditsWithTypesReq, out *MovieCreditsWithTypesRep) error
		ImageAll(ctx context.Context, in *ImageAllReq, out *ImageAllRep) error
		LocationMovies(ctx context.Context, in *LocationMoviesReq, out *LocationMoviesRep) error
		MovieComingNew(ctx context.Context, in *MovieComingNewReq, out *MovieComingNewRep) error
		Search(ctx context.Context, in *SearchReq, out *SearchRep) error
		GetFilmsByCidADay(ctx context.Context, in *GetFilmsByCidADayReq, out *GetFilmsByCidADayRsp) error
	}
	type FilmServiceExt struct {
		filmServiceExt
	}
	h := &filmServiceExtHandler{hdlr}
	return s.Handle(s.NewHandler(&FilmServiceExt{h}, opts...))
}

type filmServiceExtHandler struct {
	FilmServiceExtHandler
}

func (h *filmServiceExtHandler) HotPlayMovies(ctx context.Context, in *HotPlayMoviesReq, out *HotPlayMoviesRep) error {
	return h.FilmServiceExtHandler.HotPlayMovies(ctx, in, out)
}

func (h *filmServiceExtHandler) MovieDetail(ctx context.Context, in *MovieDetailReq, out *MovieDetailRep) error {
	return h.FilmServiceExtHandler.MovieDetail(ctx, in, out)
}

func (h *filmServiceExtHandler) MovieCreditsWithTypes(ctx context.Context, in *MovieCreditsWithTypesReq, out *MovieCreditsWithTypesRep) error {
	return h.FilmServiceExtHandler.MovieCreditsWithTypes(ctx, in, out)
}

func (h *filmServiceExtHandler) ImageAll(ctx context.Context, in *ImageAllReq, out *ImageAllRep) error {
	return h.FilmServiceExtHandler.ImageAll(ctx, in, out)
}

func (h *filmServiceExtHandler) LocationMovies(ctx context.Context, in *LocationMoviesReq, out *LocationMoviesRep) error {
	return h.FilmServiceExtHandler.LocationMovies(ctx, in, out)
}

func (h *filmServiceExtHandler) MovieComingNew(ctx context.Context, in *MovieComingNewReq, out *MovieComingNewRep) error {
	return h.FilmServiceExtHandler.MovieComingNew(ctx, in, out)
}

func (h *filmServiceExtHandler) Search(ctx context.Context, in *SearchReq, out *SearchRep) error {
	return h.FilmServiceExtHandler.Search(ctx, in, out)
}

func (h *filmServiceExtHandler) GetFilmsByCidADay(ctx context.Context, in *GetFilmsByCidADayReq, out *GetFilmsByCidADayRsp) error {
	return h.FilmServiceExtHandler.GetFilmsByCidADay(ctx, in, out)
}
