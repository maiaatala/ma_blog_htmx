package services

import (
	"bytes"
	"errors"
	"net/http"
	"os"
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
