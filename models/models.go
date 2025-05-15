package models

type ShortPostPaginated struct {
	TotalItems int         `json:"totalItems"`
	Page       int         `json:"page"`
	Items      []ShortPost `json:"items"`
}

type ShortPost struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Author      struct {
		Name string `json:"name"`
	} `json:"author"`
}
