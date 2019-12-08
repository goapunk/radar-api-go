package group

import (
	"0xacab.org/radarapi/file"
	"0xacab.org/radarapi/location"
	"0xacab.org/radarapi/term"
	"encoding/json"
)

type Group struct {
	// Location Fields
	Body          *Body                `json:"body"`
	Category      []term.Term          `json:"category"`
	IsGroup       bool                 `json:"group_group"`
	Logo          *Logo                `json:"group_logo"`
	Image         []*file.File         `json:"image"`
	Email         string               `json:"email"`
	Link          []*Link              `json:"link"`
	Offline       []*location.Location `json:"offline"`
	OpeningTimes  *OpeningTimes        `json:"opening_times"`
	Phone         string               `json:"phone"`
	Topic         []*term.Term         `json:"topic"`
	Notifications []string             `json:"notifications"`
	Active        bool                 `json:"active"`
	// Seems unused?
	Flyer                      []interface{} `json:"flyer"`
	Members                    []*User       `json:"members"`
	Members1                   []*User       `json:"members__1"`
	Members2                   []*User       `json:"members__2"`
	Members3                   []*User       `json:"members__3"`
	RadarGroupListedBy         []*Group      `json:"radar_group_listed_by"`
	NId                        int64         `json:"nid,string"`
	VId                        int64         `json:"vid,string"`
	IsNew                      bool          `json:"is_new"`
	Type                       string        `json:"type"`
	Title                      string        `json:"title"`
	Language                   string        `json:"language"`
	Url                        string        `json:"url"`
	EditUrl                    string        `json:"edit_url"`
	Status                     int           `json:"status,string"`
	Promote                    int           `json:"promote,string"`
	Sticky                     int           `json:"sticky,string"`
	Created                    int64         `json:"created,string"`
	Changed                    int64         `json:"changed,string"`
	FeedNid                    string        `json:"feed_nid"`
	FlagAbuseNodeUser          []*User       `json:"flag_abuse_node_user"`
	FlagAbuseWhitelistNodeUser []*User       `json:"flag_abuse_whitelist_node_user"`
	UUID                       string        `json:"uuid"`
	VUUID                      string        `json:"vuuid"`
	// Reference fields (when a term is referenced in another entity),
	// see ResolveField().
	ReferenceUri string `json:"uri"`
	// interface because there is an inconsistency, usually it's a string, in radar_group_listed_by it's an int.
	ReferenceId       interface{} `json:"id"`
	ReferenceResource string      `json:"resource"`
}

type Body struct {
	Value   string `json:"value"`
	Summary string `json:"summary"`
	Format  string `json:"format"`
}

type Logo struct {
	File *file.File
}

type Link struct {
	Url        string   `json:"url"`
	Attributes []string `json:"attributes"`
	DisplayUrl string   `json:"display_url"`
}

type OpeningTimes struct {
	Value  string `json:"value"`
	Format string `json:"format"`
}

type User struct {
	Uri      string `json:"uri"`
	Id       int    `json:"id,string"`
	Resource string `json:"resource"`
	Label    string `json:"label"`
}

func (e *Group) UnmarshalJSON(data []byte) error {
	type tmp *Group
	return unmarshalJSON(tmp(e), data)
}

func (e *Body) UnmarshalJSON(data []byte) error {
	type tmp *Body
	return unmarshalJSON(tmp(e), data)
}

func (e *Logo) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	var buf map[string]*file.File
	err := json.Unmarshal(data, &buf)
	if err != nil {
		return err
	}
	e.File = buf["file"]
	return nil
}

func (e *Link) UnmarshalJSON(data []byte) error {
	type tmp *Link
	return unmarshalJSON(tmp(e), data)
}

func (e *OpeningTimes) UnmarshalJSON(data []byte) error {
	type tmp *OpeningTimes
	return unmarshalJSON(tmp(e), data)
}

func (e *User) UnmarshalJSON(data []byte) error {
	type tmp *User
	return unmarshalJSON(tmp(e), data)
}

func unmarshalJSON(e interface{}, data []byte) error {
	if len(data) < 3 {
		return nil
	}
	err := json.Unmarshal(data, e)
	return err
}
