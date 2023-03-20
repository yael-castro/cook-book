package server

import (
	"encoding/json"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server/response"
	"mime"
	"net/http"
)

// JSON sends a http response encode
func JSON(w http.ResponseWriter, status int, a any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(a)
}

// Bind based on the content type, decodes the request body
//
// Content types supported:
//
// - application/json
func Bind(w http.ResponseWriter, r *http.Request, a any) (ok bool) {
	content := r.Header.Get("Content-Type")
	if content == "" {
		_ = JSON(w, http.StatusBadRequest, response.Common{
			Message: "missing content type",
		})
		return
	}

	media, _, _ := mime.ParseMediaType(content)
	switch media {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(a)
		if err != nil {
			_ = JSON(w, http.StatusUnprocessableEntity, response.Common{
				Message: "malformed body or unexpected error occurs during decoding process",
			})
			return
		}

		return true
	}

	w.Header().Set("Accept", "application/json")

	_ = JSON(w, http.StatusUnsupportedMediaType, response.Common{
		Message: fmt.Sprintf("unsupported media type: '%s' is not valid", media),
	})

	return
}
