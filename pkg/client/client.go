package client

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) []byte {
	client := &http.Client{}

	response, error := client.Get(url)

	if error != nil {
		panic(error)
	}

	fmt.Println(response.Status)

	data, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	return data
}
