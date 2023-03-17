package dto

type CommentCreate struct {
	BlogID uint64 `json:"blog_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
