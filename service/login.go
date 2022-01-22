package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	redirectURI = "http://localhost:8080/callback"
	state       = "abc123"
	tokenUrl    = "https://accounts.spotify.com/api/token"
)

var (
	ch         = make(chan string)
	httpClient = &http.Client{}
)

//Login https://github.com/zmb3/spotify/blob/master/examples/authenticate/authcode/authenticate.go
func Login() {
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		_ = http.ListenAndServe(":8080", nil)
	}()

	loginUrl := "https://accounts.spotify.com/authorize?response_type=code&client_id=" + viper.GetString("client_id") +
		"&redirect_uri=" + redirectURI +
		"&scope=user-read-private%20user-modify-playback-state%20user-read-playback-state" +
		"&state=" + state

	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", loginUrl)

	// wait for auth to complete
	wait := <-ch
	fmt.Println(wait)

}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Get("error") != "" {
		fmt.Println("error: ", query.Get("error"))
		return
	}

	if query.Get("code") != "" {
		if err := getBearerToken(query.Get("code")); err != nil {
			fmt.Println(err)
			return
		}
	}
	ch <- "Successfully logged in"
}

func getBearerToken(code string) error {
	data := url.Values{
		"code":         []string{code},
		"redirect_uri": []string{redirectURI},
		"grant_type":   []string{"authorization_code"},
	}
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodPost, tokenUrl, body)
	if err != nil {
		return err
	}

	headerVal := base64.StdEncoding.EncodeToString([]byte(viper.GetString("CLIENT_ID") + ":" + viper.GetString("SECRET")))
	req.Header.Set("Authorization", "Basic "+headerVal)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	bodyBytes, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return err
	}

	var resp map[string]interface{}
	if err = json.Unmarshal(bodyBytes, &resp); err != nil {
		return err
	}
	if resp["error"] != nil {
		return errors.New("\nLogin error: " + fmt.Sprintf("%v", resp["error_description"]))
	}

	viper.Set("access_token", resp["access_token"])

	if err = viper.WriteConfig(); err != nil {
		return err
	}
	return err
}
