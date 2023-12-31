package repository

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Id            int64  `json:"id,omitempty"`
	Author        *User  `json:"author,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

func (Video) TableName() string {
	return "videos"
}
