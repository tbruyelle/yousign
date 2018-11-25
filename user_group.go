package yousign

type UserGroupService struct {
	client *Client
}

type UserGroup struct {
	ID          *string  `json:"id"`
	Name        *string  `json:"name"`
	Permissions []string `json:"permissions"`
}
