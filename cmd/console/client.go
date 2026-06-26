package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "http://localhost:8080"

func doRequest(
	method string,
	endpoint string,
	data any,
) ([]byte, error) {

	var body io.Reader

	if data != nil {

		jsonData, err := json.Marshal(data)

		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(
		method,
		baseURL+endpoint,
		body,
	)

	if err != nil {
		return nil, err
	}

	if data != nil {
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {

		return nil, fmt.Errorf(
			"server returned %s\n%s",
			resp.Status,
			string(responseBody),
		)
	}

	return responseBody, nil
}

// function to call create tasks
func post(endpoint string, data any) ([]byte, error) {

	return doRequest(
		http.MethodPost,
		endpoint,
		data,
	)
}

// functoin to get tasks
func get(endpoint string) ([]byte, error) {

	return doRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
}

// function to update tasks
func put(endpoint string, data any) ([]byte, error) {

	return doRequest(
		http.MethodPut,
		endpoint,
		data,
	)
}

// functio to delete tasks
func deleteRequest(endpoint string) error {

	_, err := doRequest(
		http.MethodDelete,
		endpoint,
		nil,
	)

	return err
}
