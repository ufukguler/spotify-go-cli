package service

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"net/http"
)

func Search(q string) error {
	fmt.Printf("Searching for: %v\n", q)
	req, err := http.NewRequest(http.MethodGet, SearchBaseUrl, nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	query := req.URL.Query()
	query.Add("q", q)
	query.Add("type", "track")
	query.Add("offset", "0")
	req.URL.RawQuery = query.Encode()

	res, err := sendRequest(req)
	if err != nil {
		return err
	}
	var search SearchTrack
	if err2 := decode(res.Body, &search); err2 != nil {
		fmt.Printf("Decoding error: %v", err2)
	}

	rows := make([]table.Row, 0)
	for _, item := range search.Tracks.Items {
		rows = append(rows, table.Row{
			item.Name,
			item.Album.Name,
			item.Explicit,
			item.Uri,
		})
	}
	header := table.Row{"Name", "Album Name", "Explicit", "Uri"}
	prettyPrintTable(header, rows)
	return err
}
