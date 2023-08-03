package db

import (
	"cinema/src/film-srv/entity"
	"database/sql"
)

// 获取影片剧照
func SelectFilimImages(movie_id int64) ([]*entity.Image, error) {
	images := []*entity.Image{}
	err := db.Select(&images, "SELECT `image_id`,`movie_id`,`image_url` FROM `image`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return images, err
}
