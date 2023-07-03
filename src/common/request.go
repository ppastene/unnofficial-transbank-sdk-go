package common

import (
	"bytes"
	"io"
	"net/http"
)

const CONTENT_TYPE string = "application/json"

func HTTPRequest(method string, endpoint string, payload *[]byte, options Options) []byte {
	commerceCode, apiKey := options.CommerceCode, options.ApiKey
	url := options.GetApiBaseUrl() + endpoint
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(*payload))
	if err != nil {
		panic("Error al preparar la request")
	}
	req.Header.Set("Tbk-Api-Key-Id", commerceCode)
	req.Header.Set("Tbk-Api-Key-Secret", apiKey)
	req.Header.Set("Content-Type", CONTENT_TYPE)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error al comunicar con el servidor")
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error al leer la respuesta")
	}
	return bodyText
}
