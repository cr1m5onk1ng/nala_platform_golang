package domain

type MappedUserPost struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	ResourceID      int64  `json:"resource_id"`
	PostTitle       string `json:"post_title"`
	PostDescription string `json:"post_description"`
}
