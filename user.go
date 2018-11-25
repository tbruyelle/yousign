package yousign

import (
	"net/http"
	"time"
)

type UserService struct {
	client *Client
}

type User struct {
	ID               *string     `json:"id,omitempty"`
	Firstname        *string     `json:"firstname,omitempty"`
	Lastname         *string     `json:"lastname,omitempty"`
	FullName         *string     `json:"fullName,omitempty"`
	Title            *string     `json:"title,omitempty"`
	Email            *string     `json:"email,omitempty"`
	Phone            *string     `json:"phone,omitempty"`
	Status           *string     `json:"status,omitempty"`
	Company          *string     `json:"company,omitempty"`
	CreatedAt        *time.Time  `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time  `json:"updatedAt,omitempty"`
	Config           interface{} `json:"config,omitempty"`
	SamlNameID       *string     `json:"samlNameId,omitempty"`
	DefaultSignImage *string     `json:"defaultSignImage,omitempty"`
	FastSign         *bool       `json:"fastSign,omitempty"`
	Group            *UserGroup  `json:"group,omitempty"`
	Notifications    interface{} `json:"notifications,omitempty"`
	Deleted          *bool       `json:"deleted,omitempty"`
	DeletedAt        *time.Time  `json:"deletedAt,omitempty"`
}

type UserRequest struct {
	Firstname        *string     `json:"firstname,omitempty"`
	Lastname         *string     `json:"lastname,omitempty"`
	Title            *string     `json:"title,omitempty"`
	Phone            *string     `json:"phone,omitempty"`
	Company          *string     `json:"company,omitempty"`
	Config           interface{} `json:"config,omitempty"`
	Group            *string     `json:"group,omitempty"`
	DefaultSignImage *string     `json:"defaultSignImage,omitempty"`
	Notifications    *string     `json:"notifications,omitempty"`
}

func (s *UserService) All() ([]User, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "users", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []User
	resp, err := s.client.Do(req, &users)
	return users, resp, err
}
