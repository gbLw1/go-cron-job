package utils

import (
	"encoding/json"
  "fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](fullUrl, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, error) {
  client := http.Client{}
  u, err := url.Parse(fullUrl)
  if err != nil {
    return responseType, err
  }

  if httpMethod == "GET" {
    q := u.Query()

    for k,v := range queryParameters {
      q.Set(k, strings.Join(v, ","))
    }

    u.RawQuery = q.Encode()
  }

  req, err := http.NewRequest(httpMethod, u.String(), body)
  if err != nil {
    return responseType, err
  }

  for k,v := range headers {
    req.Header.Set(k, v)
  }

  res, err := client.Do(req)
  if err != nil {
    return responseType, err
  }

  if res == nil {
    return responseType, fmt.Errorf("error: calling %s returned empty response", u.String())
  }

  responseData, err := io.ReadAll(res.Body)
  if err != nil {
    return responseType, err
  }

  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return responseType, fmt.Errorf("error: calling %s:\nstatus: %s\nresponse data: %s", u.String(), res.Status, responseData)
  }

  var responseObj T
  err = json.Unmarshal(responseData, &responseObj)
  if err != nil {
    return responseType, err
  }

  return responseObj, nil
}

