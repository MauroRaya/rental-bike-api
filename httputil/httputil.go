package httputil

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ParsePathID(r *http.Request) (int32, error) {
	value := r.PathValue("id")
	id, err := strconv.ParseInt(value, 10, 32)
	return int32(id), err
}

func DecodeJSON(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func EncodeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
