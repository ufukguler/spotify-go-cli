package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"net/http"
	"strings"
)

const PlayerBaseUrl = "https://api.spotify.com/v1/me/player"

var response ErrorResponse

func Current() error {
	req, err := http.NewRequest(http.MethodGet, PlayerBaseUrl+"/currently-playing", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	res, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Http client error:%v", err)
		return err
	}

	var current CurrentSong
	if err2 := decode(res.Body, &current); err2 != nil {
		fmt.Printf("Decoding error: %v", err2)
	}
	artists := make([]string, 0)
	for _, artist := range current.Item.Artists {
		artists = append(artists, artist.Name)
	}
	img := ""
	if len(current.Item.Album.Images) > 0 {
		img = current.Item.Album.Images[0].Url
	}
	row := table.Row{current.Item.Name,
		strings.Join(artists, ", "),
		current.Item.Album.Name,
		img,
		current.IsPlaying}
	header := table.Row{"Song", "Artists", "Album Name", "Album Cover", "Is Playing"}
	prettyPrintTable(header, append(make([]table.Row, 0), row))
	return err
}

func Transfer(arg string) error {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	playback := TransferPlayback{
		DeviceIds: []string{arg},
		Play:      true,
	}
	if err := encoder.Encode(&playback); err != nil {
		fmt.Printf("Encoding error:%v", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPut, PlayerBaseUrl, &buffer)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	_, err = sendRequest(req)
	return err
}

func Repeat(arg string) error {
	url := PlayerBaseUrl + "/repeat?state=" + arg
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	_, err = sendRequest(req)
	return err
}

func Next() error {
	req, err := http.NewRequest(http.MethodPost, PlayerBaseUrl+"/next", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	_, err = sendRequest(req)
	return err
}

func Previous() error {
	req, err := http.NewRequest(http.MethodPost, PlayerBaseUrl+"/previous", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	_, err = sendRequest(req)
	return err
}

func PauseResume() error {
	req, err := http.NewRequest(http.MethodPut, PlayerBaseUrl+"/pause", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	_, err = sendRequest(req)
	return err
}

func Start() error {
	req, err := http.NewRequest(http.MethodPut, PlayerBaseUrl+"/play", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	_, err = sendRequest(req)
	return err
}

func Volume(arg string) error {
	req, err := http.NewRequest(http.MethodPut, PlayerBaseUrl+"/volume", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	query := req.URL.Query()
	query.Set("volume_percent", arg)
	req.URL.RawQuery = query.Encode()
	_, err = sendRequest(req)
	return err
}

func Devices() error {
	req, err := http.NewRequest(http.MethodGet, PlayerBaseUrl+"/devices", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	res, err := sendRequest(req)
	if err != nil {
		return err
	}
	var devices Device
	if err2 := decode(res.Body, &devices); err2 != nil {
		fmt.Printf("Decoding error: %v", err2)
	}
	rows := make([]table.Row, 0)
	for _, d := range devices.Devices {
		rows = append(rows, table.Row{
			d.Id, d.Name, d.IsActive, d.Type, d.VolumePercent, d.IsPrivateSession, d.IsRestricted,
		})
	}
	header := table.Row{"ID", "Name", "IsActive", "Type", "Volume", "Private Session", "Restricted"}
	prettyPrintTable(header, rows)
	return err
}
