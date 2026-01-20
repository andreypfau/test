package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-service/models"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *APIClient) GetStudent(id string) (*models.Student, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/v1/students/%s", c.baseURL, id))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, body)
	}

	// try parsing as wrapped response first
	var apiResp models.StudentResponse
	if err := json.Unmarshal(body, &apiResp); err == nil && apiResp.Success {
		return &apiResp.Data, nil
	}

	// fallback to direct student object
	var student models.Student
	if err := json.Unmarshal(body, &student); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &student, nil
}
