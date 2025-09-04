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
	"io/ioutil"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/text/language"
)

const defautlBaseUrl = "https://radar.squat.net/api/1.2"

type RadarClient struct {
	web     *http.Client
	baseUrl string
	timeout int
	log     *slog.Logger
}

// WithTimeout configures a timeout for RadarClient.
func WithTimeout(timeout int) func(*RadarClient) {
	return func(r *RadarClient) {
		r.timeout = timeout
	}
}

// WithLogger configures a logger for RadarClient.
func WithLogger(logger *slog.Logger) func(*RadarClient) {
	return func(r *RadarClient) {
		r.log = logger
	}
}

// Returns an instance of RadarClient which can be used to interact with the
// radar server.
func NewRadarClient(opts ...func(*RadarClient)) *RadarClient {
	radarClient := &RadarClient{
		web:     &http.Client{Timeout: time.Second * 30},
		baseUrl: defautlBaseUrl,
		log:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	for _, optFunc := range opts {
		optFunc(radarClient)
	}

	return radarClient
}

func (radar *RadarClient) SetBaseUrl(baseUrl string) {
	if baseUrl != "" {
		radar.baseUrl = baseUrl
	}
}

func (radar *RadarClient) GetLogger() *slog.Logger {
	return radar.log
}

func (radar *RadarClient) GetTimeout() int {
	return radar.timeout
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
		radar.log.Error(err.Error())
		return "", err
	}
	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		radar.log.Error(err.Error())
		return "", err
	}
	return string(body), nil
}
