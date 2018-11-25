package yousign

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

var (
	client = NewClientStaging("STAGING-KEY")
)

func fatal(t *testing.T, err error, resp *http.Response) {
	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	t.Fatalf("error %v : %s", err, b.String())
}
