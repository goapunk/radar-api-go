package radarapi

import (
	"fmt"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseUrl = "https://radar.squat.net/api/1.2/"

type RadarClient struct {
	web *http.Client
}

// Returns an instance of RadarClient which can be used to interact with the
// radar server.
func NewRadarClient() *RadarClient {
	return &RadarClient{
		&http.Client{Timeout: time.Second * 10},
	}
}

// Search the radar database.
//
// The returned string is the raw json response. See the examples on how to umarshal it.
func (radar *RadarClient) Search(search *SearchBuilder) (string, error) {
	u, err := prepareSearchUrl(search)
	if err != nil {
		return "", err
	}
	return radar.runQuery(u)
}

func (radar *RadarClient) prepareAndRunEntityQuery(rawUrl string, language *language.Tag, fields []string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	query := u.Query()
	if len(fields) != 0 {
		query.Set("fields", fieldsToCommaString(fields))
	}
	if language != nil {
		base, _ := language.Base()
		query.Set("language", base.String())
	}
	u.RawQuery = query.Encode()
	return radar.runQuery(u)
}

func (radar *RadarClient) runQuery(u *url.URL) (string, error) {
	resp, err := radar.web.Get(u.String())
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	return string(body), nil
}
