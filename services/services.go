package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"ssrhtmx/models"
)

func PostContactForm(body []byte) error {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		return errors.New("API_BASE_URL not set")
	}

	resp, err := http.Post(baseURL+"/contact", "application/x-www-form-urlencoded", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status: " + resp.Status)
	}

	return nil
}

func FetchPostByID(id string) (*models.Post, error) {
	apiURL := os.Getenv("API_BASE_URL")
	if apiURL == "" {
		return nil, errors.New("API_BASE_URL is not set")
	}

	url := fmt.Sprintf("%s/posts/%s", apiURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API error: " + resp.Status)
	}

	var post models.Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func FetchPosts(page int) (*models.ShortPostPaginated, error) {
	apiURL := os.Getenv("API_BASE_URL")
	if apiURL == "" {
		return nil, errors.New("API_BASE_URL is not set")
	}

	url := fmt.Sprintf("%s/posts?page=%d", apiURL, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API error: " + resp.Status)
	}

	var paginated models.ShortPostPaginated
	if err := json.NewDecoder(resp.Body).Decode(&paginated); err != nil {
		return nil, err
	}

	return &paginated, nil
}
