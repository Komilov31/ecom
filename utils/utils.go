package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// type JSONTime time.Time

var Validate = validator.New()

func ParseJson(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(&v)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

// func (t JSONTime) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(time.DateOnly))), nil
// }

// func (t *JSONTime) UnmarshalJSON(b []byte) error {
// 	parsedTime, err := time.Parse(fmt.Sprintf("\"%s\"", time.DateOnly), string(b))
// 	if err != nil {
// 		return err
// 	}
// 	*t = JSONTime(parsedTime)
// 	return nil
// }
