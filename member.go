package yousign

import (
	"net/http"
	"time"
)

type MemberService struct {
	client *Client
}

type Member struct {
	ID                   *string               `json:"id,omitempty"`
	User                 *string               `json:"user,omitempty"`
	Type                 *string               `json:"type,omitempty"`
	Firstname            *string               `json:"firstname,omitempty"`
	Lastname             *string               `json:"lastname,omitempty"`
	Email                *string               `json:"email,omitempty"`
	Phone                *string               `json:"phone,omitempty"`
	Position             *int                  `json:"position,omitempty"`
	CreatedAt            *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time            `json:"updatedAt,omitempty"`
	Status               *string               `json:"status,omitempty"`
	FileObjects          []MemberFileObject    `json:"fileObjects,omitempty"`
	Comment              *string               `json:"comment,omitempty"`
	Procedure            *string               `json:"procedure,omitempty"`
	OperationLevel       *string               `json:"operationLevel,omitempty"`
	OperationCustomModes []string              `json:"operationCustomModes,omitempty"`
	ModeSmsConfiguration *ModeSmsConfiguration `json:"modeSmsConfiguration,omitempty"`
}

type MemberFileObject struct {
	ID               *string    `json:"id,omitempty"`
	File             *File      `json:"file,omitempty"`
	Page             *int       `json:"page,omitempty"`
	Position         *string    `json:"position,omitempty"`
	FieldName        *string    `json:"fieldName,omitempty"`
	Mention          *string    `json:"mention,omitempty"`
	Mention2Internal *string    `json:"mention2 (internal),omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
	ExecutedAt       *time.Time `json:"executedAt,omitempty"`
}

type MemberRequest struct {
	User                 *string               `json:"user,omitempty"`
	Type                 *string               `json:"type,omitempty"`
	Firstname            *string               `json:"firstname,omitempty"`
	Lastname             *string               `json:"lastname,omitempty"`
	Email                *string               `json:"email,omitempty"`
	Phone                *string               `json:"phone,omitempty"`
	Position             *int                  `json:"position,omitempty"`
	FileObjects          []FileObjectRequest   `json:"fileObjects,omitempty"`
	Procedure            *string               `json:"procedure,omitempty"`
	OperationLevel       *string               `json:"operationLevel,omitempty"`
	OperationCustomModes []string              `json:"operationCustomModes,omitempty"`
	ModeSmsConfiguration *ModeSmsConfiguration `json:"modeSmsConfiguration,omitempty"`
}

type ModeSmsConfiguration struct {
	Content string `json:"content,omitempty"`
}

func (s *MemberService) Create(r *MemberRequest) (*Member, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "/members", nil, r)
	if err != nil {
		return nil, nil, err
	}

	var m Member
	resp, err := s.client.Do(req, &m)
	return &m, resp, err
}

func (s *MemberService) Get(id string) (*Member, *http.Response, error) {
	req, err := s.client.NewRequest("GET", id, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var m Member
	resp, err := s.client.Do(req, &m)
	return &m, resp, err
}
