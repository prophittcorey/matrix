package matrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	Username string
	Password string

	// HTTPClient is used to perform all HTTP requests. You can specify your own
	// to set a custom timeout, proxy, etc.
	HTTPClient = http.Client{
		Timeout: 30 * time.Second,
	}

	// UserAgent will be used in each request's user agent header field.
	UserAgent = "github.com/prophittcorey/matrix"
)

const (
	baseURL = "https://matrix.org/_matrix/client/r0"
)

func Send(roomID, message string) error {
	token, err := AuthToken()

	if err != nil {
		return err
	}

	var data = []byte(fmt.Sprintf(`{
		"body": "%s",
		"msgtype": "m.text"
	}`, message))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/rooms/%s/send/m.room.message?access_token=%s", baseURL, url.QueryEscape(roomID), url.QueryEscape(token)), bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	if _, err := HTTPClient.Do(req); err != nil {
		return err
	}

	return nil
}

func AuthToken() (string, error) {
	var data = []byte(fmt.Sprintf(`{"user": "%s", "password": "%s", "type": "m.login.password"}`, Username, Password))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/login", baseURL), bytes.NewBuffer(data))

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	res, err := HTTPClient.Do(req)

	if err != nil {
		return "", err
	}

	var auth struct {
		Token string `json:"access_token"`
	}

	if bs, err := io.ReadAll(res.Body); err == nil {
		if err := json.Unmarshal(bs, &auth); err == nil {
			return auth.Token, nil
		}
	}

	return "", fmt.Errorf(`error: failed to unmarshal auth object`)
}
