package entity

type Like struct {
	BlogID uint64 `gorm:"primaryKey;foreignKey" json:"blog_id"`
	UserID uint64 `gorm:"primaryKey;foreignKey" json:"user_id"`
	Blog   *Blog  `json:"blog,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   *User  `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
