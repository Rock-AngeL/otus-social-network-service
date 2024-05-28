package model

// Пост пользователя
type Post struct {
	Id string `json:"id,omitempty"`

	Text string `json:"text,omitempty"`

	AuthorUserId string `json:"author_user_id,omitempty"`
}
