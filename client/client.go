package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseUrl       string
	Timeout       time.Duration
	Authorization string
	httpClient    *http.Client
}

type HttpClientParam struct {
	Method      string
	Url         string
	Data        io.Reader
	ContentType string
	Accept      string
}

func NewClient(baseUrl string, timeout time.Duration, authorization string) *Client {
	return &Client{
		BaseUrl:       baseUrl,
		Timeout:       timeout,
		Authorization: authorization,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

func (c *Client) OauthRequest(data url.Values) ([]byte, int, error) {
	resp, err := c.createHttpRequest(context.Background(), HttpClientParam{
		Method:      "POST",
		Url:         c.BaseUrl,
		Data:        bytes.NewReader([]byte(data.Encode())),
		ContentType: "application/x-www-form-urlencoded",
	})

	if err != nil {
		return nil, 0, fmt.Errorf("oauth request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, resp.StatusCode, nil
}

func (c *Client) SendGetRequestWithContext(ctx context.Context, path string) ([]byte, int, error) {
	resp, err := c.createHttpRequest(ctx, HttpClientParam{
		Method: "GET",
		Url:    c.BaseUrl + path,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("get request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, resp.StatusCode, nil
}

func (c *Client) SendGetRequest(path string) ([]byte, int, error) {
	resp, err := c.createHttpRequest(context.Background(), HttpClientParam{
		Method: "GET",
		Url:    c.BaseUrl + path,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("get request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, resp.StatusCode, nil
}

func (c *Client) createHttpRequest(ctx context.Context, param HttpClientParam) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, param.Method, param.Url, param.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	accept := param.Accept
	if accept == "" {
		accept = "application/json"
	}
	req.Header.Set("Accept", accept)

	contentType := param.ContentType
	if contentType == "" {
		contentType = "application/json"
	}
	req.Header.Set("Content-Type", contentType)

	if c.Authorization != "" {
		req.Header.Set("Authorization", "Bearer "+c.Authorization)
	}

	req.Header.Set("User-Agent", "osuapi-wrapper-go/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	return resp, nil
}

func (c *Client) HandleError(statusCode int, respBody string) error {
	maxLen := 500

	if len(respBody) > maxLen {
		respBody = respBody[:maxLen] + "... (truncated)"
	}

	return fmt.Errorf("HTTP %d: %s", statusCode, respBody)
}
