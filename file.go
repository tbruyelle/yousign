package yousign

import (
	"encoding/json"
	"net/http"
	"time"
)

type FileService struct {
	client *Client
}

type File struct {
	ID          *string         `json:"id,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Type        *string         `json:"type,omitempty"`
	ContentType *string         `json:"contentType,omitempty"`
	Description *string         `json:"description,omitempty"`
	CreatedAt   *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time      `json:"updatedAt,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
	Company     *string         `json:"company,omitempty"`
	Creator     *string         `json:"creator,omitempty"`
}

type FileRequest struct {
	Name        *string           `json:"name,omitempty"`
	Type        *string           `json:"type,omitempty"`
	Password    *string           `json:"password,omitempty"`
	Description *string           `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Content     *string           `json:"content,omitempty"`
	Procedure   *string           `json:"procedure,omitempty"`
}

func (s *FileService) Create(r *FileRequest) (*File, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "files", nil, r)
	if err != nil {
		return nil, nil, err
	}

	var f File
	resp, err := s.client.Do(req, &f)
	return &f, resp, err
}

func (s *FileService) Get(id string) (*File, *http.Response, error) {
	req, err := s.client.NewRequest("GET", id, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var f File
	resp, err := s.client.Do(req, &f)
	return &f, resp, err
}

func (s *FileService) Download(id string) (*string, *http.Response, error) {
	req, err := s.client.NewRequest("GET", id+"/download", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var data string
	resp, err := s.client.Do(req, &data)
	return &data, resp, err
}
