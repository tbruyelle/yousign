package yousign

import (
	"net/http"
	"time"
)

type FileObjectService struct {
	client *Client
}

type FileObject struct {
	ID               *string    `json:"id,omitempty"`
	File             *string    `json:"file,omitempty"`
	Member           *Member    `json:"member,omitempty"`
	Page             *int       `json:"page,omitempty"`
	Position         *string    `json:"position,omitempty"`
	FieldName        *string    `json:"fieldName,omitempty"`
	Mention          *string    `json:"mention,omitempty"`
	Mention2Internal *string    `json:"mention2 (internal),omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
	ExecutedAt       *time.Time `json:"executedAt,omitempty"`
}

type FileObjectRequest struct {
	File      *string `json:"file,omitempty"`
	Page      *int    `json:"page,omitempty"`
	Position  *string `json:"position,omitempty"`
	FieldName *string `json:"fieldName,omitempty"`
	Mention   *string `json:"mention,omitempty"`
	Mention2  *string `json:"mention2,omitempty"`
	Member    *string `json:"member,omitempty"`
}

func (s *FileObjectService) Create(r *FileObjectRequest) (*FileObject, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "file_objects", nil, r)
	if err != nil {
		return nil, nil, err
	}

	var v FileObject
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *FileObjectService) Get(id string) (*FileObject, *http.Response, error) {
	req, err := s.client.NewRequest("GET", id, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v FileObject
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
