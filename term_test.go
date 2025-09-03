package radarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTerm(t *testing.T) {
	const result = `{"tid":"3","name":"course/workshop","description":"","weight":"5","node_count":10,"url":"https://radar.squat.net/en/category/course-workshop","parent":[],"parents_all":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/3","id":"3","resource":"taxonomy_term"}],"feed_nid":null,"type":"category","uuid":"2a7f6975-4c01-4777-8611-dffe0306c06f"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/taxonomy_term/2a7f6975-4c01-4777-8611-dffe0306c06f.json" {
			t.Errorf("Expected to request '/api/1.2/taxonomy_term/2a7f6975-4c01-4777-8611-dffe0306c06f.json', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()

	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")

	value, err := radar.Term("2a7f6975-4c01-4777-8611-dffe0306c06f", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if value.UUID != "2a7f6975-4c01-4777-8611-dffe0306c06f" {
		t.Errorf("Expected UUID 2a7f6975-4c01-4777-8611-dffe0306c06f, got: %s", value.UUID)
	}
	if value.Name != "course/workshop" {
		t.Errorf("Expected name 'course/workshop', got: %s", value.Name)
	}

}
