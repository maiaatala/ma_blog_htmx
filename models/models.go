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

type Author struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type Post struct {
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Content     string   `json:"content"`
	Author      Author   `json:"author"`
	Tags        []string `json:"tags"`
}

type Comment struct {
	ID              string  `json:"id"`
	PostID          string  `json:"postId"`
	ParentCommentID *string `json:"parentCommentId,omitempty"`
	CreatedAt       string  `json:"createdAt"`
	Content         string  `json:"content"`
	HasReplies      bool    `json:"hasReplies"`
	Author          struct {
		Name  string `json:"name"`
		Photo string `json:"photo"`
	} `json:"author"`
}
