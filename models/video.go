package models

type Video struct {
	Id            int64  `gorm:"column:id;type:int" json:"id,omitempty"`
	AuthorId      int64  `gorm:"column:author_id" json:"author_id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(255)" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:play_url;type:varchar(255)" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count;type:int" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count;type:int" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"column:is_favorite;type:tinyint(1)" json:"is_favorite,omitempty"`
}

func (Video) TableName() string {
	return "video"
}
