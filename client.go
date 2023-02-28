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
	// HTTPClient is used to perform all HTTP requests. You can specify your own
	// to set a custom timeout, proxy, etc.
	HTTPClient = http.Client{
		Timeout: 30 * time.Second,
	}

	// UserAgent will be used in each request's user agent header field.
	UserAgent = "github.com/prophittcorey/matrix"

	// BaseURL is the location to base all API requests from.
	BaseURL = "https://matrix.org/_matrix/client/r0"
)

type Client struct {
	username string
	password string
	token    string
}

func (c *Client) Authenticate() error {
	if len(c.token) != 0 {
		return nil /* we already have an auth token */
	}

	var data = []byte(fmt.Sprintf(`{"user": "%s", "password": "%s", "type": "m.login.password"}`, Username, Password))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/login", BaseURL), bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	res, err := HTTPClient.Do(req)

	if err != nil {
		return err
	}

	var auth struct {
		Token string `json:"access_token"`
		Error string `json:"error"`
	}

	if bs, err := io.ReadAll(res.Body); err == nil {
		if err := json.Unmarshal(bs, &auth); err == nil {
			if len(auth.Error) != 0 {
				return fmt.Errorf(`error: %s`, auth.Error)
			}

			c.token = auth.Token

			return nil
		}
	}

	return fmt.Errorf(`error: failed to unmarshal auth object`)
}

func (c *Client) Send(roomID, message string) error {
	if err := c.Authenticate(); err != nil {
		return err
	}

	var data = []byte(fmt.Sprintf(`{ "body": "%s", "msgtype": "m.text" }`, message))

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/rooms/%s/send/m.room.message", BaseURL, url.QueryEscape(roomID)),
		bytes.NewBuffer(data),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, c.token))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	if _, err := HTTPClient.Do(req); err != nil {
		return err
	}

	return nil
}

// New creates a new Matrix client for a given user.
func New(username, password string) *Client {
	return Client{
		username: username,
		password: password,
	}
}
