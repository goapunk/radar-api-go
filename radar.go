// radarapi - provides bindings for the radar.squat.net event API.
// Copyright (C) 2019  Julian Dehm
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package radarapi

import (
	"fmt"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const defautlBaseUrl = "https://radar.squat.net/api/1.2"

type RadarClient struct {
	web     *http.Client
	baseUrl string
}

// Returns an instance of RadarClient which can be used to interact with the
// radar server.
func NewRadarClient() *RadarClient {
	return &RadarClient{
		web:     &http.Client{Timeout: time.Second * 10},
		baseUrl: defautlBaseUrl,
	}
}

func (radar *RadarClient) SetBaseUrl(baseUrl string) {
	if baseUrl != "" {
		radar.baseUrl = baseUrl
	}
}

func (radar *RadarClient) prepareAndRunEntityQuery(rawUrl string, language *language.Tag, fields []string) (string, error) {
	u, err := url.Parse(radar.baseUrl + "/" + rawUrl)
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
