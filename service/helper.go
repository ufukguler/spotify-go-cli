package service

import (
	"encoding/json"
	"errors"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

func sendRequest(req *http.Request) (*http.Response, error) {
	initHeaders(req)
	res, err := httpClient.Do(req)
	if err != nil {
		return res, err
	}

	if res.StatusCode >= 400 {
		decoder := json.NewDecoder(res.Body)
		if err2 := decoder.Decode(&response); err2 != nil {
			return res, err2
		}
		return res, errors.New(response.Error.Message)
	}
	return res, nil
}

func initHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}

func prettyPrintTable(header table.Row, rows []table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.Render()
}

func decode(body io.ReadCloser, dst interface{}) error {
	return json.NewDecoder(body).Decode(dst)
}
