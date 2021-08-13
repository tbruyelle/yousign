package yousign

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

var (
	client = NewClientStaging("STAGING-KEY")
)

func fatal(t *testing.T, err error, resp *http.Response) {
	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	t.Fatalf("error %v : %s", err, b.String())
}

func TestCorrectStagingSigningURL(t *testing.T) {
	assert := assert.New(t)

	apiKey := faker.UUIDHyphenated()
	client := NewClientStaging(apiKey)
	memberID := faker.UUIDHyphenated()

	signURL := client.SignURL(memberID)

	expectedURL := "https://staging-app.yousign.com/procedure/sign?members=" + memberID
	assert.Equal(expectedURL, signURL)
}

func TestCorrectProductionSigningURL(t *testing.T) {
	assert := assert.New(t)

	apiKey := faker.UUIDHyphenated()
	client := NewClient(apiKey)
	memberID := faker.UUIDHyphenated()

	signURL := client.SignURL(memberID)

	expectedURL := "https://webapp.yousign.com/procedure/sign?members=" + memberID
	assert.Equal(expectedURL, signURL)
}
