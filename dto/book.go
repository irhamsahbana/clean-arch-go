package dto

type BookUpsertRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
}
