package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jo-nike/vikunja-cli/internal/config"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

type RequestOption func(req *http.Request)

func WithQuery(key, value string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set(key, value)
		req.URL.RawQuery = q.Encode()
	}
}

func WithPage(page int) RequestOption {
	return WithQuery("page", strconv.Itoa(page))
}

func WithPerPage(perPage int) RequestOption {
	return WithQuery("per_page", strconv.Itoa(perPage))
}

func WithSearch(s string) RequestOption {
	return WithQuery("s", s)
}

func WithSort(sort string) RequestOption {
	return WithQuery("sort_by", sort)
}

func WithOrderBy(order string) RequestOption {
	return WithQuery("order_by", order)
}

func WithFilter(filter string) RequestOption {
	return WithQuery("filter", filter)
}

func WithFilterBy(filterBy string) RequestOption {
	return WithQuery("filter_by", filterBy)
}

func WithFilterValue(filterValue string) RequestOption {
	return WithQuery("filter_value", filterValue)
}

func WithFilterComparator(filterComparator string) RequestOption {
	return WithQuery("filter_comparator", filterComparator)
}

func New(cfg *config.Config) (*Client, error) {
	if cfg.URL == "" {
		return nil, fmt.Errorf("vikunja URL not configured (set url in config.toml or VIKUNJA_URL env var)")
	}

	baseURL := strings.TrimRight(cfg.URL, "/")

	return &Client{
		baseURL:    baseURL + "/api/v1",
		token:      cfg.Token,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) do(method, path string, body interface{}, opts ...RequestOption) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	reqURL := c.baseURL + path
	req, err := http.NewRequest(method, reqURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, parseErrorResponse(resp)
	}

	return resp, nil
}

// Get performs a GET request and decodes the response into result.
func (c *Client) Get(path string, result interface{}, opts ...RequestOption) error {
	resp, err := c.do(http.MethodGet, path, nil, opts...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}

// GetList performs a GET request for list endpoints and returns paginated results.
func (c *Client) GetList(path string, result interface{}, opts ...RequestOption) (*PaginationInfo, error) {
	resp, err := c.do(http.MethodGet, path, nil, opts...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return parsePaginationHeaders(resp), nil
}

// Create performs a PUT request (Vikunja uses PUT for creation).
func (c *Client) Create(path string, body interface{}, result interface{}, opts ...RequestOption) error {
	resp, err := c.do(http.MethodPut, path, body, opts...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

// Update performs a POST request (Vikunja uses POST for updates).
func (c *Client) Update(path string, body interface{}, result interface{}, opts ...RequestOption) error {
	resp, err := c.do(http.MethodPost, path, body, opts...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

// Delete performs a DELETE request.
func (c *Client) Delete(path string, opts ...RequestOption) error {
	resp, err := c.do(http.MethodDelete, path, nil, opts...)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// DeleteWithBody performs a DELETE request with a JSON body.
func (c *Client) DeleteWithBody(path string, body interface{}, opts ...RequestOption) error {
	resp, err := c.do(http.MethodDelete, path, body, opts...)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// Post performs a raw POST request (for endpoints that actually use POST for creation).
func (c *Client) Post(path string, body interface{}, result interface{}, opts ...RequestOption) error {
	resp, err := c.do(http.MethodPost, path, body, opts...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

// GetRaw performs a GET request and returns the raw response body.
func (c *Client) GetRaw(path string, opts ...RequestOption) ([]byte, error) {
	resp, err := c.do(http.MethodGet, path, nil, opts...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// DoUpload performs a multipart file upload.
func (c *Client) DoUpload(method, path, fieldName, filePath string, result interface{}, opts ...RequestOption) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("creating form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("copying file data: %w", err)
	}

	writer.Close()

	reqURL := c.baseURL + path
	req, err := http.NewRequest(method, reqURL, &buf)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseErrorResponse(resp)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

// DownloadFile performs a GET request and saves the response body to a file.
func (c *Client) DownloadFile(path, destPath string, opts ...RequestOption) error {
	reqURL := c.baseURL + path
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseErrorResponse(resp)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// PostDownloadFile performs a POST request with a JSON body and saves the response body to a file.
func (c *Client) PostDownloadFile(path, destPath string, body interface{}, opts ...RequestOption) error {
	resp, err := c.do(http.MethodPost, path, body, opts...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// BaseURL returns the configured base URL (without /api/v1).
func (c *Client) BaseURL() string {
	return strings.TrimSuffix(c.baseURL, "/api/v1")
}

// PostForm performs a POST with URL-encoded form data.
func (c *Client) PostForm(path string, values url.Values, result interface{}, opts ...RequestOption) error {
	reqURL := c.baseURL + path
	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseErrorResponse(resp)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}
