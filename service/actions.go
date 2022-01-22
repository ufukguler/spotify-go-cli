package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
)

const BASEURL = "https://api.spotify.com/v1/me/player"

var response ErrorResponse

func Current() error {
	req, err := http.NewRequest(http.MethodGet, BASEURL+"/currently-playing", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	query := req.URL.Query()
	req.URL.RawQuery = query.Encode()
	initHeaders(req)
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	var current CurrentSong
	decoder := json.NewDecoder(res.Body)
	if err2 := decoder.Decode(&current); err2 != nil {
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
	req, err := http.NewRequest(http.MethodPut, BASEURL, &buffer)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	return sendRequest(req)
}

func Repeat(arg string) error {
	url := BASEURL + "/repeat?state=" + arg
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	return sendRequest(req)
}

func Next() error {
	req, err := http.NewRequest(http.MethodPost, BASEURL+"/next", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	return sendRequest(req)
}

func Previous() error {
	req, err := http.NewRequest(http.MethodPost, BASEURL+"/previous", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	return sendRequest(req)
}

func PauseResume() error {
	req, err := http.NewRequest(http.MethodPut, BASEURL+"/pause", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	return sendRequest(req)
}

func Start() error {
	req, err := http.NewRequest(http.MethodPut, BASEURL+"/play", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	return sendRequest(req)
}

func Volume(arg string) error {
	req, err := http.NewRequest(http.MethodPut, BASEURL+"/volume", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	query := req.URL.Query()
	query.Set("volume_percent", arg)
	req.URL.RawQuery = query.Encode()
	return sendRequest(req)
}

func Devices() error {
	req, err := http.NewRequest(http.MethodGet, BASEURL+"/devices", nil)
	if err != nil {
		fmt.Printf("Http request error:%v", err)
		return err
	}
	query := req.URL.Query()
	req.URL.RawQuery = query.Encode()
	initHeaders(req)
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	var devices Device
	decoder := json.NewDecoder(res.Body)
	if err2 := decoder.Decode(&devices); err2 != nil {
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

func sendRequest(req *http.Request) error {
	initHeaders(req)
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("Response Code %d\n", res.StatusCode)
	if res.StatusCode >= 400 {
		decoder := json.NewDecoder(res.Body)
		err2 := decoder.Decode(&response)
		if err2 != nil {
			fmt.Printf("ErrorResponse Message %v", response.Error.Message)
			return err
		}
		fmt.Printf("Error Message: %s", response.Error.Message)
	}
	return err
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
