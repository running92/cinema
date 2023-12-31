package db

import (
	"cinema/src/place-srv/entity"
	"database/sql"
)

func SelectPlaces() ([]*entity.Place, error) {

	places := []*entity.Place{}
	err := db.Select(&places, "SELECT `id`,`count`,`name`,`pinyin_full`,`pinyin_short` FROM `place`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return places, err
}
