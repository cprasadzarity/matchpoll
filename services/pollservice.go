package services

import (
	"io"
	redis "matchpoll/config"
	"net/http"
)

func Poll() {
	endpoint := "https://reqres.in/api/users/2"
	resp, _ := http.Get(endpoint)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	redis.Upsert("score", string(body))
}
