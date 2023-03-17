package entity

type Blog struct {
	ID       uint64    `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title"`
	URL      string    `json:"url"`
	Content  string    `json:"content"`
	UserID   uint64    `json:"user_id" gorm:"foreignKey"`
	User     *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Likes    []Like    `json:"likes,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}
