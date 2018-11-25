package yousign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	stagingApiURL    = "https://staging-api.yousign.com"
	productionApiURL = "https://api.yousign.com"

	stagingAppURL    = "https://staging-app.yousign.com"
	productionAppURL = "https://api.yousign.com"
)

type Client struct {
	apiKey string
	apiURL *url.URL
	appURL *url.URL
	client *http.Client

	User        *UserService
	UserGroup   *UserGroupService
	Member      *MemberService
	File        *FileService
	FileObject  *FileObjectService
	Procedure   *ProceduresService
	ServerStamp *ServerStampService
}

// NewClient returns a new YouSign API client. You must provide a valid
// apiKey retrieved from your YouSign account.
func NewClient(apiKey string) *Client {
	return newClient(apiKey, false)
}

func NewClientStaging(apiKey string) *Client {
	return newClient(apiKey, true)
}

func newClient(apiKey string, staging bool) *Client {
	c := &Client{
		apiKey: apiKey,
		client: http.DefaultClient,
	}

	var err error
	apiURL := productionApiURL
	appURL := productionAppURL
	if staging {
		apiURL = stagingApiURL
		appURL = stagingAppURL
	}
	c.apiURL, err = url.Parse(apiURL)
	if err != nil {
		panic(err)
	}
	c.appURL, err = url.Parse(appURL)
	if err != nil {
		panic(err)
	}

	c.User = &UserService{c}
	c.UserGroup = &UserGroupService{c}
	c.Member = &MemberService{c}
	c.File = &FileService{c}
	c.FileObject = &FileObjectService{c}
	c.Procedure = &ProceduresService{c}
	c.ServerStamp = &ServerStampService{c}
	return c
}

func (c *Client) SignURL(memberID string) string {
	return fmt.Sprintf("%s/procedure/sign?members=%s", c.appURL.String(), memberID)
}

// NewRequest creates an API request. This method can be used to performs
// API request not implemented in this library. Otherwise it should not be
// be used directly.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method, urlStr string, opt interface{}, body interface{}) (*http.Request, error) {
	rel, err := addOptions(urlStr, opt)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// Do performs the request, the json received in the response is decoded
// and stored in the value pointed by v.
// Do can be used to perform the request created with NewRequest, which
// should be used only for API requests not implemented in this library.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c := resp.StatusCode; c < 200 || c > 299 {
		return resp, fmt.Errorf("Server returns status %d", c)
	}

	if v != nil {
		defer resp.Body.Close()
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (*url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	if opt == nil {
		return u, nil
	}

	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		// No query string to add
		return u, nil
	}

	qs, err := query.Values(opt)
	if err != nil {
		return nil, err
	}

	u.RawQuery = qs.Encode()
	return u, nil
}

func String(v string) *string { return &v }

func Bool(v bool) *bool { return &v }

func Int(v int) *int { return &v }
