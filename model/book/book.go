package book

type Book struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
