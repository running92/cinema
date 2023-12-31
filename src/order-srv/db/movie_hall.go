package db

import (
	"cinema/src/order-srv/entity"
	"database/sql"
)

func SelectMHNameMHId(mhId int64) (*entity.MovieHall, error) {

	hall := entity.MovieHall{}
	err := db.Get(&hall, "SELECT `mh_name`,`cinema_id` FROM `movie_hall` WHERE `mh_id` = ?", mhId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &hall, err
}
