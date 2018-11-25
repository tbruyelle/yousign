package yousign

import (
	"encoding/json"
	"net/http"
	"time"
)

type ServerStampService struct {
	client *Client
}

type ServerStamp struct {
	ID          *string         `json:"id,omitempty"`
	File        *string         `json:"file,omitempty"`
	Certificate *string         `json:"certificate,omitempty"`
	FileObjects []FileObject    `json:"fileObjects,omitempty"`
	Config      json.RawMessage `json:"config,omitempty"`
	CreatedAt   *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time      `json:"updatedAt,omitempty"`
	FinishedAt  *time.Time      `json:"finishedAt,omitempty"`
	Status      *string         `json:"status,omitempty"`
	Company     *string         `json:"company,omitempty"`
}

type ServerStampRequest struct {
	File        *string             `json:"file,omitempty"`
	Certificate *string             `json:"certificate,omitempty"`
	Config      *ServerStampConfig  `json:"config,omitempty"`
	FileObjects []FileObjectRequest `json:"fileObjects,omitempty"`
	SignImage   *string             `json:"signImage,omitempty"`
}

type ServerStampConfig struct {
	Webhook struct {
		ServerStampFinished []Webhook `json:"server_stamp.finished,omitempty"`
	}
}

func (s *ServerStampService) Create(r *ServerStampRequest) (*ServerStamp, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "server_stamps", nil, r)
	if err != nil {
		return nil, nil, err
	}

	var v ServerStamp
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *ServerStampService) Get(id string) (*ServerStamp, *http.Response, error) {
	req, err := s.client.NewRequest("GET", id, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v ServerStamp
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
