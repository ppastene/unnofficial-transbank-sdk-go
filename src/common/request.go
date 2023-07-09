package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ppastene/unnofficial-transbank-sdk-go/src/exceptions"
)

const CONTENT_TYPE string = "application/json"

func HTTPRequest(method string, endpoint string, payload *[]byte, options Options) ([]byte, error) {
	commerceCode, apiKey := options.CommerceCode, options.ApiKey
	url := options.GetApiBaseUrl() + endpoint
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(*payload))
	if err != nil {
		return nil, &exceptions.TransbankExpcetion{ErrorMessage: "Error preparing the request", StatusCode: 0}
	}
	req.Header.Set("Tbk-Api-Key-Id", commerceCode)
	req.Header.Set("Tbk-Api-Key-Secret", apiKey)
	req.Header.Set("Content-Type", CONTENT_TYPE)
	resp, err := client.Do(req)
	if err != nil {
		return nil, &exceptions.TransbankExpcetion{ErrorMessage: "Error communicating with Transbank server", StatusCode: 0}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &exceptions.TransbankExpcetion{ErrorMessage: "Error reading the response", StatusCode: 0}
	}
	if !isStatusOK(resp.StatusCode) {
		var exception exceptions.TransbankExpcetion
		err = json.Unmarshal(body, &exception)
		exception.StatusCode = resp.StatusCode
		return nil, &exception
	}
	return body, nil
}

func isStatusOK(status int) bool {
	if status >= 200 && status < 300 {
		return true
	}
	return false
}
