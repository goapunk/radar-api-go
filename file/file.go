package file

import (
	"encoding/json"
)

// Available fields for a file.
const (
	// All will return all available fields.
	FieldAll       = "*"
	FieldFileId    = "fid"
	FieldName      = "name"
	FieldMime      = "mime"
	FieldSize      = "size"
	FieldURL       = "url"
	FieldTimestamp = "timestamp"
	FielDFeedNid   = "feed_nid"
	FieldUUID      = "uuid"
	FieldType      = "type"
)

type File struct {
	FileId    int64  `json:"fid,string"`
	Name      string `json:"name"`
	Mime      string `json:"mime"`
	Size      int64  `json:"size,string"`
	Url       string `json:"url"`
	Timestamp int64  `json:"timestamp,string"`
	FeedNid   string `json:"feed_nid"`
	UUID      string `json:"uuid"`
	Type      string `json:"type"`
	// Reference fields (when a file is referenced in another entity),
	// see ResolveField().
	ReferenceUri      string `json:"uri,omitempty"`
	ReferenceId       string `json:"id,omitempty"`
	ReferenceResource string `json:"resource,omitempty"`
	// Only seems to be displayed in search results, not in event details
	ReferenceFilename string `json:"filename,omitempty"`
}

func (e *File) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	type tmp *File
	err := json.Unmarshal(data, tmp(e))
	return err
}
