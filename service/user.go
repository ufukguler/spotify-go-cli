package service

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"net/http"
)

func Profile() error {
	req, err := http.NewRequest(http.MethodGet, UserBaseUrl, nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	res, err := sendRequest(req)
	if err != nil {
		return err
	}
	var user UserProfile
	if err2 := decode(res.Body, &user); err2 != nil {
		fmt.Printf("Decoding error: %v", err2)
	}
	profileUrl := "https://open.spotify.com/user/" + user.Id
	row := table.Row{user.Id,
		user.DisplayName,
		user.Email,
		user.ExplicitContent.FilterEnabled,
		user.Followers.Total,
		user.Type,
		profileUrl}
	header := table.Row{"ID", "Name", "Email", "Explicit Filter", "Followers", "Type", "Url"}
	prettyPrintTable(header, append(make([]table.Row, 0), row))
	return err
}
