package rls

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// setHeaders sets the headers for all HTTP requests
func (rls *RLSClient) setHeaders(req *http.Request) {
	for k, v := range rls.cfg.ExtraHeaders {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json; charset-utf-8")
	req.Header.Set("Accept", "application/json; charset-utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("basic %s", rls.Credential()))
}

// handleResponse handles HTTP responses and unmarshals JSON to the appropriate object
func handleResponse(res *http.Response, response interface{}) error {
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		// var errresp ErrorResponse
		var errmsg string
		body, err := io.ReadAll(res.Body)
		// err is not json
		if err != nil {
			// body is empty
			if len(body) == 0 {
				errmsg = "response body is empty"
			}
		} else {
			errmsg = string(body)
		}
		return fmt.Errorf("error code %d: %s", res.StatusCode, errmsg)
	}

	if response != nil {
		err := json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			msg, err := io.ReadAll(res.Body)
			return fmt.Errorf("failed to parse response : %s : %w", msg, err)
		}
	}
	return nil
}

// createCredential creates the basic auth credential used to authenticate requests to RLS API
func createCredential(apiKey string) string {
	key := fmt.Sprintf("%s:%s", apiKey, apiKey)
	return b64.StdEncoding.EncodeToString([]byte(key))
}

// sendRequest handles sending HTTP requests
func (rls *RLSClient) sendRequest(req *http.Request, response interface{}) error {
	req = req.WithContext(rls.Ctx)
	rls.setHeaders(req)

	res, err := rls.HTTPClient.Do(req)
	if err != nil {
		select {
		case <-rls.Ctx.Done():
			return rls.Ctx.Err()
		default:
			return err
		}
	}
	defer res.Body.Close()
	return handleResponse(res, response)
}
