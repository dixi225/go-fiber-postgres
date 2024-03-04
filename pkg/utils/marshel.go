package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		return
	}

}
