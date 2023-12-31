package handler

import (
	"cinema/src/place-srv/db"
	"cinema/src/share/pb"
	"cinema/src/share/utils/log"
	"context"
	"go.uber.org/zap"
)

type PlaceServiceExtHandler struct {
	logger *zap.Logger
}

func NewPlaceServiceExtHandler() *PlaceServiceExtHandler {
	return &PlaceServiceExtHandler{
		logger: log.Instance(),
	}
}

func (p *PlaceServiceExtHandler) HotCitiesByCinema(ctx context.Context, req *pb.HotCitiesByCinemaReq, rsp *pb.HotCitiesByCinemaRep) error {

	places, err := db.SelectPlaces()
	if err != nil {
		p.logger.Error("err", zap.Any("places", err))
		return err
	}
	placesPB := []*pb.Place{}
	for _, place := range places {
		placePB := place.ToProtoDBHotPlayMovies()
		placesPB = append(placesPB, placePB)
	}
	rsp.P = placesPB
	return nil
}
