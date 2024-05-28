package dto

type LoginBody struct {
	// Id string `json:"id,omitempty"`
	Email string `json:"email" binding:"required,email"`

	Password string `json:"password,omitempty"`
}
