package entity

type Comment struct {
	ID     uint64 `json:"id" gorm:"primary_key"`
	Text   string `json:"text"`
	BlogID uint64 `json:"blog_id"`
	UserID uint64 `json:"user_id"`
	Blog   *Blog  `json:"blog,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   *User  `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
