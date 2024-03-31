package utils

import (
	"io"
	"net/http"
)

func MakeRequest(url string) ([]byte, error) {
	request, error := http.NewRequest("GET", url, nil)

	if error != nil {
		return []byte(""), error
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		return []byte(""), error
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		return []byte(""), error
	}

	defer response.Body.Close()
	return responseBody, nil
}
