package helpers

import (
	"encoding/json"
	"net/http"
)

func Send(rw http.ResponseWriter, data []byte) {
	res := APIResponse{}
	json.Unmarshal(data, &res)

	rw.Header().Set("Content-Type", "application/json")
	if res.StatCode >= 400 {
		rw.WriteHeader(int(res.StatCode))
	}
	json.NewEncoder(rw).Encode(&res)
}
