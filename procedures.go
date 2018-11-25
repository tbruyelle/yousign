package yousign

import (
	"encoding/json"
	"net/http"
	"time"
)

type ProceduresService struct {
	client *Client
}

type Procedure struct {
	ID                 *string         `json:"id,omitempty"`
	Name               *string         `json:"name,omitempty"`
	Description        *string         `json:"description,omitempty"`
	CreatedAt          *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt          *time.Time      `json:"updatedAt,omitempty"`
	ExpiresAt          *time.Time      `json:"expiresAt,omitempty"`
	Status             *string         `json:"status,omitempty"`
	Creator            *string         `json:"creator,omitempty"`
	CreatorFirstName   *string         `json:"creatorFirstName,omitempty"`
	CreatorLastName    *string         `json:"creatorLastName,omitempty"`
	Company            *string         `json:"company,omitempty"`
	Template           *bool           `json:"template,omitempty"`
	Ordered            *bool           `json:"ordered,omitempty"`
	Parent             *string         `json:"parent,omitempty"`
	Metadata           json.RawMessage `json:"metadata,omitempty"`
	Config             json.RawMessage `json:"config,omitempty"`
	Members            []Member        `json:"members,omitempty"`
	Files              []File          `json:"files,omitempty"`
	RelatedFilesEnable *bool           `json:"relatedFilesEnable,omitempty"`
	Archive            *bool           `json:"archive,omitempty"`
}

type ProcedureConfig struct {
	Email struct {
		ProcedureStarted  []Msg `json:"procedure.started,omitempty"`
		ProcedureFinished []Msg `json:"procedure.finished,omitempty"`
		ProcedureRefused  []Msg `json:"procedure.refused,omitempty"`
		ProcedureExpired  []Msg `json:"procedure.expired,omitempty"`
		ProcedureDeleted  []Msg `json:"procedure.deleted,omitempty"`
		MemberStarted     []Msg `json:"member.started,omitempty"`
		MemberFinished    []Msg `json:"member.finished,omitempty"`
		CommentCreated    []Msg `json:"comment.created,omitempty"`
	} `json:"email,omitempty"`
	Reminders []struct {
		Interval *int `json:"interval,omitempty"`
		Limit    *int `json:"limit,omitempty"`
		Config   struct {
			Email struct {
				ReminderExecuted []Msg `json:"reminder.executed,omitempty"`
			} `json:"email,omitempty"`
		} `json:"config,omitempty"`
	} `json:"reminders,omitempty"`
	Webhook struct {
		ProcedureStarted  []Webhook `json:"procedure.started,omitempty"`
		ProcedureFinished []Webhook `json:"procedure.finished,omitempty"`
		ProcedureRefused  []Webhook `json:"procedure.refused,omitempty"`
		ProcedureExpired  []Webhook `json:"procedure.expired,omitempty"`
		ProcedureDeleted  []Webhook `json:"procedure.deleted,omitempty"`
		MemberStarted     []Webhook `json:"member.started,omitempty"`
		MemberFinished    []Webhook `json:"member.finished,omitempty"`
		CommentCreated    []Webhook `json:"comment.created,omitempty"`
	} `json:"webhook,omitempty"`
}

type Msg struct {
	To       []string `json:"to,omitempty"`
	Subject  *string  `json:"subject,omitempty"`
	Message  *string  `json:"message,omitempty"`
	FromName *string  `json:"fromName,omitempty"`
}

type Webhook struct {
	URL     *string        `json:"url,omitempty"`
	Method  *string        `json:"method,omitempty"`
	Headers *WebhookHeader `json:"headers,omitempty"`
}

type WebhookHeader struct {
	XYousignCustomHeader string `json:"X-Yousign-Custom-Header,omitempty"`
}

type ProcedureRequest struct {
	Name               *string           `json:"name,omitempty"`
	Description        *string           `json:"description,omitempty"`
	ExpiresAt          *string           `json:"expiresAt,omitempty"`
	Template           *bool             `json:"template,omitempty"`
	Ordered            *bool             `json:"ordered,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	Config             *ProcedureConfig  `json:"config,omitempty"`
	Members            []MemberRequest   `json:"members,omitempty"`
	Start              *bool             `json:"start,omitempty"`
	RelatedFilesEnable *bool             `json:"relatedFilesEnable,omitempty"`
	Archive            *bool             `json:"archive,omitempty"`
}

func (s *ProceduresService) Create(r *ProcedureRequest) (*Procedure, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "procedures", nil, r)
	if err != nil {
		return nil, nil, err
	}

	var v Procedure
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *ProceduresService) Get(id string) (*Procedure, *http.Response, error) {
	req, err := s.client.NewRequest("GET", id, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v Procedure
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *ProceduresService) Update(id string, r *ProcedureRequest) (*Procedure, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", id, nil, r)
	if err != nil {
		return nil, nil, err
	}

	var v Procedure
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
