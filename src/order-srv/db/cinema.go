package db

import (
	"cinema/src/order-srv/entity"
	"database/sql"
)

func SelectCinemaByCid(cinemaId int64) (*entity.Cinema, error) {

	cinema := entity.Cinema{}
	err := db.Get(&cinema, "SELECT `cinema_name`,`cinema_add`,`cinema_phone` FROM `cinema` WHERE `cinema_id` = ?", cinemaId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &cinema, err
}
