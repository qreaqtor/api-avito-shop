package web

import (
	"encoding/json"
	"io"
	"net/http"
)

// Читает тело запроса, используя json.Unmarshal
func ReadRequestBody(r *http.Request, v any) error {
	if r.Header.Get(ContentType) != ContentTypeJSON {
		return errUnknownPayload
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = r.Body.Close()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
