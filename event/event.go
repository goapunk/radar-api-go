package event

import (
	"github.com/goapunk/radar-api-go/file"
	"github.com/goapunk/radar-api-go/group"
	"github.com/goapunk/radar-api-go/location"
	"github.com/goapunk/radar-api-go/term"
	"encoding/json"
)

type Event struct {
	// Location Fields
	Body                       *Body                `json:"body"`
	Category                   []term.Term          `json:"category"`
	GroupContentAccess         int                  `json:"group_content_access,string"`
	OgGroupRef                 []*group.Group       `json:"og_group_ref"`
	OgGroupRequest             []*group.Group       `json:"og_group_request"`
	DateTime                   []*DateTime          `json:"date_time"`
	Image                      *Image               `json:"image"`
	Price                      []string             `json:"price"`
	PriceCategory              []*term.Term         `json:"price_category"`
	EventStatus                string               `json:"event_status"`
	Email                      string               `json:"email"`
	Link                       []*Link              `json:"link"`
	Offline                    []*location.Location `json:"offline"`
	Phone                      string               `json:"phone"`
	Topic                      []*term.Term         `json:"topic"`
	TitleField                 string               `json:"title_field"`
	Flyer                      []Flyer              `json:"flyer"`
	Membership                 []*User              `json:"og_membership"`
	Membership1                []*User              `json:"og_membership__1"`
	Membership2                []*User              `json:"og_membership__2"`
	Membership3                []*User              `json:"og_membership__3"`
	GroupRefMembership         []*User              `json:"og_group_ref__og_membership"`
	GroupRefMembership1        []*User              `json:"og_group_ref__og_membership__1"`
	GroupRefMembership2        []*User              `json:"og_group_ref__og_membership__2"`
	GroupRefMembership3        []*User              `json:"og_group_ref__og_membership__3"`
	GroupRequestMembership     []*User              `json:"og_group_request__og_membership"`
	GroupRequestMembership1    []*User              `json:"og_group_request__og_membership__1"`
	GroupRequestMembership2    []*User              `json:"og_group_request__og_membership__2"`
	GroupRequestMembership3    []*User              `json:"og_group_request__og_membership__3"`
	NId                        int64                `json:"nid,string"`
	VId                        int64                `json:"vid,string"`
	IsNew                      bool                 `json:"is_new"`
	Type                       string               `json:"type"`
	Title                      string               `json:"title"`
	Language                   string               `json:"language"`
	Url                        string               `json:"url"`
	EditUrl                    string               `json:"edit_url"`
	Status                     int                  `json:"status,string"`
	Promote                    int                  `json:"promote,string"`
	Sticky                     int                  `json:"sticky,string"`
	Created                    int64                `json:"created,string"`
	Changed                    int64                `json:"changed,string"`
	FeedNid                    string               `json:"feed_nid"`
	FlagAbuseNodeUser          []*User              `json:"flag_abuse_node_user"`
	FlagAbuseWithelistNodeUser []*User              `json:"flag_abuse_withelist_node_user"`
	UUID                       string               `json:"uuid"`
	VUUID                      string               `json:"vuuid"`
	// Reference fields (when a term is referenced in another entity),
	// see ResolveField().
	ReferenceUri      string `json:"uri,omitempty"`
	ReferenceId       string `json:"id,omitempty"`
	ReferenceResource string `json:"resource,omitempty"`
}

type Body struct {
	Value   string `json:"value"`
	Summary string `json:"summary"`
	Format  string `json:"format"`
}

type DateTime struct {
	Start     int64  `json:"value,string"`
	End       int64  `json:"value2,string"`
	Duration  int64  `json:"duration"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	RRule     string `json:"rrule"`
}

type Image struct {
	File *file.File
}

type Link struct {
	Url        string   `json:"url"`
	Attributes []string `json:"attributes"`
	DisplayUrl string   `json:"display_url"`
}

type Flyer struct {
	File        *file.File `json:"file"`
	Description string     `json:"description"`
}

type User struct {
	Uri      string `json:"uri"`
	Id       int    `json:"id"`
	Resource string `json:"resource"`
	Label    string `json:"label"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type tmp *Event
	err := json.Unmarshal(data, tmp(e))
	return err
}

func (b *Body) UnmarshalJSON(data []byte) error {
	type tmp *Body
	return unmarshalJSON(tmp(b), data)
}

func (l *Link) UnmarshalJSON(data []byte) error {
	type tmp *Link
	return unmarshalJSON(tmp(l), data)
}

func (u *User) UnmarshalJSON(data []byte) error {
	type tmp *User
	return unmarshalJSON(tmp(u), data)
}

func (i *Image) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		return nil
	}

	type tmp *Image
	if err := unmarshalJSON(tmp(i), data); err != nil {
		return err
	}
	return nil
}

func unmarshalJSON(e interface{}, data []byte) error {
	if len(data) < 3 {
		return nil
	}
	err := json.Unmarshal(data, e)
	return err
}
