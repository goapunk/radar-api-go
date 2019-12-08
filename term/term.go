package term

import "encoding/json"

type Term struct {
	// Term Fields
	Id          int64  `json:"tid,string"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Weight      int    `json:"weight,string"`
	NodeCount   int    `json:"node_count"`
	Url         string `json:"url"`
	// Parents seems to always be empty
	Parents    []*Term `json:"parent,omitempty"`
	ParentsAll []*Term `json:"parents_all"`
	FeedNid    string  `json:"feed_nid"`
	Type       string  `json:"type"`
	UUID       string  `json:"uuid"`
	// Reference fields (when a term is referenced in another entity),
	// see ResolveField().
	ReferenceUri      string `json:"uri,omitempty"`
	ReferenceId       string `json:"id,omitempty"`
	ReferenceResource string `json:"resource,omitempty"`
}

func (e *Term) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	type tmp *Term
	err := json.Unmarshal(data, tmp(e))
	return err
}
