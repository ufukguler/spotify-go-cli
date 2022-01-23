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
	"time"
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
		"&scope=" + scope +
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
	return sendTokenRequest(data)
}

func RefreshToken() error {
	t, err := time.Parse(time.RFC3339, viper.GetString("expires_at"))
	if err != nil {
		return err
	}
	if time.Now().After(t) {
		data := url.Values{
			"refresh_token": []string{viper.GetString("refresh_token")},
			"grant_type":    []string{"refresh_token"},
		}
		return sendTokenRequest(data)
	}
	return nil
}

func sendTokenRequest(data url.Values) error {
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodPost, tokenUrl, body)
	if err != nil {
		return err
	}

	basicHeader := base64.StdEncoding.EncodeToString([]byte(viper.GetString("CLIENT_ID") + ":" + viper.GetString("SECRET")))
	req.Header.Set("Authorization", "Basic "+basicHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	now := time.Now()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var resp map[string]interface{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return err
	}
	if resp["error"] != nil {
		return errors.New("\nLogin error: " + fmt.Sprintf("%v", resp["error_description"]))
	}
	return writeConfigs(now, resp)
}

//writeConfigs https://developer.spotify.com/documentation/general/guides/authorization/code-flow/
func writeConfigs(now time.Time, resp map[string]interface{}) error {
	expiresIn := resp["expires_in"].(float64) / 60
	expiresAt := now.Add(time.Minute * time.Duration(expiresIn-1))
	viper.Set("access_token", resp["access_token"])
	viper.Set("expires_in", resp["expires_in"])
	viper.Set("refresh_token", resp["refresh_token"])
	viper.Set("scope", resp["scope"])
	viper.Set("token_type", resp["token_type"])
	viper.Set("created_at", now)
	viper.Set("expires_at", expiresAt)
	return viper.WriteConfig()
}
