package dto

type BlogCreate struct {
	Title   string `json:"title" binding:"required"`
	URL     string `json:"url" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type BlogUpdate struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}
