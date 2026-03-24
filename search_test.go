package radarapi

import (
	"github.com/goapunk/radar-api-go/event"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchEventsCategoryDate(t *testing.T) {
	const result = `{"result":{"519308":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/1e1543ec-18eb-411f-9bc7-357e82dddc92","id":"1e1543ec-18eb-411f-9bc7-357e82dddc92","resource":"location","title":"Regenbogencafe Lausitzer Str. 22a  Berlin Deutschland"}],"title":"Mittwochscaf\u00e9"},"496473":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/fbb94a74-7bc0-4392-b708-b9a3f642fd66","id":"fbb94a74-7bc0-4392-b708-b9a3f642fd66","resource":"location","title":"JUP Florastr. 84  Berlin Deutschland"}],"title":"K\u00fcche f\u00fcr Alle (jeden Mittwoch)"},"538487":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/78d7792b-84f1-41ed-b2b5-fd3986324b2c","id":"78d7792b-84f1-41ed-b2b5-fd3986324b2c","resource":"location","title":"JUP Florastra\u00dfe 84  Berlin Germany"}],"title":"Kiezk\u00fcche f\u00fcr Alle - Kiezteam Pankow"},"542917":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/2ea82d85-d598-48b7-84f0-962de758e860","id":"2ea82d85-d598-48b7-84f0-962de758e860","resource":"location","title":"Wagenburg Lohm\u00fchle Lohm\u00fchlenstr. 17  Berlin Deutschland"}],"title":"Soli - Kinoabend: Not Just Your Picture"},"533537":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/a232175b-8c59-44be-82fa-1ad600100dd5","id":"a232175b-8c59-44be-82fa-1ad600100dd5","resource":"location","title":"Kadterschmiede Rigaer Str. 94  Berlin Deutschland"}],"title":"SOBER VOxK\u00dc"},"542374":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/a232175b-8c59-44be-82fa-1ad600100dd5","id":"a232175b-8c59-44be-82fa-1ad600100dd5","resource":"location","title":"Kadterschmiede Rigaer Str. 94  Berlin Deutschland"}],"title":"Infoevent on the Situation of Rigaer 94 + K\u00fcfa"},"501974":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/5696c22e-aba7-493c-950d-51d175449055","id":"5696c22e-aba7-493c-950d-51d175449055","resource":"location","title":"KuBiZ Bernkasteler Str. 78  Berlin Deutschland"}],"title":"K\u00fcfa (vegan)"},"543330":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/5696c22e-aba7-493c-950d-51d175449055","id":"5696c22e-aba7-493c-950d-51d175449055","resource":"location","title":"KuBiZ Bernkasteler Str. 78  Berlin Deutschland"}],"title":"VideoKino - Open Air"}},"count":8,"facets":{}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/search/events.json" {
			t.Errorf("Expected to request '/api/1.2/search/events.json, got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "facets%5Bcategory%5D%5B%5D=food&facets%5Bcity%5D%5B%5D=Berlin&facets%5Bdate%5D%5B%5D=2025-09-03&fields=title%2Coffline" {
			t.Errorf("Expected to request raw query 'facets%%5Bcategory%%5D%%5B%%5D=food&facets%%5Bcity%%5D%%5B%%5D=Berlin&facets%%5Bdate%%5D%%5B%%5D=2025-09-03&fields=title%%2Coffline', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	facets := make([]Facet, 3)
	facets[0] = Facet{event.FacetCity, "Berlin"}
	facets[1] = Facet{event.FacetCategory, "food"}
	facets[2] = Facet{event.FacetDate, "2025-09-03"}
	sb := radar.NewSearchBuilder()
	sb.Facets(facets...)
	sb.Fields(event.FieldTitle, event.FieldOffline)
	results, err := radar.SearchEvents(sb)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if results.Count != 8 {
		t.Errorf("Expected count to be 8, got: %d", results.Count)
	}
}

func TestSearchEventsNoResult(t *testing.T) {
	const result = `{"result":false,"count":0,"facets":null}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/search/events.json" {
			t.Errorf("Expected to request '/api/1.2/search/events.json', got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "facets%5Bcategory%5D%5B%5D=nonexisting" {
			t.Errorf("Expected to request raw query 'facets%%5Bcategory%%5D%%5B%%5D=nonexisting', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	facets := make([]Facet, 1)
	facets[0] = Facet{event.FacetCategory, "nonexisting"}
	sb := radar.NewSearchBuilder()
	sb.Facets(facets...)
	results, err := radar.SearchEvents(sb)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if results != nil {
		t.Errorf("Expected no results, got: %d", results.Count)
	}
}

func TestSearchEmptyEventsGroup(t *testing.T) {
	const result = `{"result":false,"count":0,"facets":{"group":[{"filter":"1599","count":1265,"formatted":" Stressfaktor"},{"filter":"547541","count":361,"formatted":"Rotterdam Alternative Events"},{"filter":"10866","count":322,"formatted":"Centro Popolare Autogestito Firenze Sud"},{"filter":"186126","count":256,"formatted":"Espace autog\u00e9r\u00e9 des Tanneries"},{"filter":"1976","count":220,"formatted":"NieuwLand"},{"filter":"109276","count":214,"formatted":"Weggeefwinkel Utrecht"},{"filter":"5018","count":210,"formatted":""},{"filter":"292860","count":180,"formatted":"L&#039;Ades&#039;if"},{"filter":"423454","count":172,"formatted":"Planlos Leipzig"},{"filter":"5075","count":166,"formatted":"Bollox"},{"filter":"65692","count":161,"formatted":"FightFamily"},{"filter":"405374","count":160,"formatted":"latscher.in"},{"filter":"41","count":158,"formatted":"Joe&#039;s Garage"},{"filter":"436012","count":157,"formatted":"KarlsUNRUH"},{"filter":"8188","count":132,"formatted":"Jugendclub Cafe K\u00f6penick"},{"filter":"378884","count":132,"formatted":"l\u2019\u00cele Egalit\u00e9"},{"filter":"4932","count":127,"formatted":"MKZ"},{"filter":"500309","count":104,"formatted":"kalender.mietenwahnsinn.info"},{"filter":"22747","count":95,"formatted":"De Klinker"},{"filter":"1352","count":93,"formatted":"Cowley Club"},{"filter":"194587","count":92,"formatted":"Plotter"},{"filter":"402346","count":85,"formatted":"Infoladen Scherer 8 "},{"filter":"16721","count":84,"formatted":"La Gryffe"},{"filter":"109236","count":82,"formatted":"Biblioteca Antiautoritaria Sacco y Vanzetti"},{"filter":"349727","count":82,"formatted":"Le Chat Noir"},{"filter":"426248","count":82,"formatted":"Espacio F\u00e9nix"},{"filter":"337427","count":81,"formatted":"Biblioth\u00e8que Anarcha-f\u00e9ministe"},{"filter":"337781","count":81,"formatted":"Biblioteca y Archivo Alberto Ghiraldo"},{"filter":"381725","count":81,"formatted":"Ath\u00e9n\u00e9e Libertaire"},{"filter":"190306","count":78,"formatted":"Le Silure"},{"filter":"456560","count":77,"formatted":"L\u2019impasse"},{"filter":"2436","count":76,"formatted":"KuZeB"},{"filter":"141066","count":74,"formatted":"backbord.tk"},{"filter":"525997","count":74,"formatted":"Vlaggen voor Palestina"},{"filter":"335919","count":73,"formatted":"Kiezhaus Agnes Reinhold"},{"filter":"34409","count":72,"formatted":"Quartier Libre des Lentill\u00e8res"},{"filter":"31746","count":69,"formatted":"Le Kiosk"},{"filter":"512988","count":69,"formatted":"Snackbar Frieda"},{"filter":"137721","count":63,"formatted":"ROMP Info &amp; Vinyl"},{"filter":"301670","count":61,"formatted":"CICP"},{"filter":"373407","count":60,"formatted":"La Kunda"},{"filter":"1686","count":59,"formatted":"KuBiZ"},{"filter":"277813","count":59,"formatted":"Plattenladen"},{"filter":"1960","count":58,"formatted":"Regenbogenfabrik"},{"filter":"17026","count":58,"formatted":"Caf\u00e9 librairie Mich\u00e8le Firk"},{"filter":"543238","count":58,"formatted":"Caf\u00e9 Gilde"},{"filter":"451645","count":56,"formatted":"L\u2019Amicale du Futur"},{"filter":"150491","count":55,"formatted":"Amicale du Combat Libre"},{"filter":"462199","count":54,"formatted":"Le Steki"},{"filter":"347081","count":52,"formatted":"BASTA! Erwerbsloseninitiative"}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/search/events.json" {
			t.Errorf("Expected to request 'https://radar.squat.net/api/1.2/search/events.json', got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "facets%5Bgroup%5D%5B%5D=1773" {
			t.Errorf("Expected to request raw query 'facets%%5Bgroup%%5D%%5B%%5D=1773', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	facets := make([]Facet, 1)
	facets[0] = Facet{event.FacetGroup, "1773"}
	sb := radar.NewSearchBuilder()
	sb.Facets(facets...)
	results, err := radar.SearchEvents(sb)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if results != nil {
		t.Errorf("Expected no results, got: %d", results.Count)
	}
}

func TestSearchEventsWithFilters(t *testing.T) {
	const result = `{"result":{"376690":{"body":{"value":"<p><strong>Kritische Orientierungswochen an der HU</strong></p>\n<p>Ein Vortrag von Bafta Sarbo.</p>\n<p>Marxismus und Antirassismus werden in deutschen akademischen und aktivistischen Kontexten selten zusammengedacht. Oft gibt es die (zum Teil legitime) Unterstellung Marxist*innen w\u00fcrden im Rassismus lediglich einen Nebenwiderspruch sehen und sich deshalb nicht w\u00fcrdig damit besch\u00e4ftigen. Eine historisch materialistische Auseinandersetzung k\u00f6nnte aber Antworten bieten auf die Fragen, was ist Rassismus und wie ist er entstanden? K\u00f6nnen wir \u00fcberhaupt von dem einen Rassismus, statt von Rassismen, sprechen? Und was hat kapitalistische Produktionsweise mit rassistischer Gewalt zu tun? Warum Rassismus mehr ist als nur eine Ideologie oder ein falsches Bewusstsein und was das Problem mit b\u00fcrgerlichem Antirassismus ist, soll in diesem Vortrag erl\u00e4utert werden.</p>\n<p>Die Zugangsdaten zur Online-Veranstaltung werden am Tag der Veranstaltung auf unserer Website und auf Facebook bekannt gegeben.</p>\n<p>--</p>\n<p>Die Kritischen Orientierungswochen sind eine von linken Studierenden organisierte Veranstaltungsreihe, die sich zum Ziel gesetzt hat, den Raum Universit\u00e4t von links zu (re)politisieren. Gleichzeitig richtet sie sich an j\u00fcngere Semester und will diesen den Einstieg ins Studi-Leben erleichtern, sowie die M\u00f6glichkeit er\u00f6ffnen, an der Uni politisch aktiv zu werden.</p>\n","summary":"","format":"rich_text_editor"},"category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/8463bb01-e974-4785-9c2d-b95d87c9ee2d","id":"8463bb01-e974-4785-9c2d-b95d87c9ee2d","resource":"taxonomy_term","name":"discussion/presentation"}],"date_time":[{"value":"1604415600","value2":"1604415600","duration":0,"time_start":"2020-11-03T16:00:00+01:00","time_end":"2020-11-03T16:00:00+01:00","rrule":null}],"image":[],"price":null,"link":[],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/6bf1f990-a54c-4b5e-a729-67ea5b55d27f","id":"6bf1f990-a54c-4b5e-a729-67ea5b55d27f","resource":"location","title":"online    Berlin Deutschland"}],"phone":null,"topic":[],"title":"Online-Vortrag: Einf\u00fchrung in die materialistische Rassismuskritik ","language":"de","url":"https://radar.squat.net/en/node/376690","created":"1602849274","uuid":"34617caf-cc09-47db-aa80-b35fdc690ee7"},"376691":{"body":{"value":"<p><strong>Kritische Orientierungswochen an der HU</strong></p>\n<p>Im Fr\u00fchling 2020 kam es zu einer landesweiten Solidarisierungswelle mit deutschen Landwirt*innen. Jobvermittlungsportale schossen schneller aus dem Boden als der Spargel. In den Medien wurden inl\u00e4ndische Erntehelfer*innen ausf\u00fchrlich dokumentiert wie sie \u00fcber R\u00fcckenschmerzen klagten. Doch schnell zeigte sich, wie abh\u00e4ngig die deutsche Landwirtschaft von der Besch\u00e4ftigung osteurop\u00e4ischer Saisonarbeiter*innen ist - darunter immer mehr ukrainische Studierende, die auf Praktikabasis Spargel stechen und Erdbeeren ernten.</p>\n<p>In unserer Veranstaltung berichtet eine ehemalige Saisonarbeiterin \u00fcber ihre Erfahrungen auf einem mecklenburger Erdbeerhof, \u00fcber dubiose Vermittlungsstrukturen und dar\u00fcber, was der Aufenthalt f\u00fcr viele ukrainische Studierende bedeutet. Danach spricht die gewerkschaftsnahe Organisation Arbeit und Leben e. V. \u00fcber g\u00e4ngige Tricks, mit denen Betriebsleiter*innen Arbeitsrechte unterwandern. Abschlie\u00dfend diskutieren wir, welche M\u00f6glichkeiten es gibt, um Saisonarbeiter*innen in Deutschland zu unterst\u00fctzen. Denn internationale Ausbeutungsverh\u00e4ltnisse ben\u00f6tigen internationale solidarische Praxen.</p>\n<p>Die Veranstaltung wird organisiert vom PECO-Institut f\u00fcr nachhaltige Entwicklung e.V.\u00a0Es wird eine Simultan\u00fcbersetzung deutsch-russisch geben. Die Zugangsdaten zur Online-Veranstaltung werden am Tag der Veranstaltung auf unserer Website und auf Facebook bekannt gegeben.</p>\n<p>--</p>\n<p>Die Kritischen Orientierungswochen sind eine von linken Studierenden organisierte Veranstaltungsreihe, die sich zum Ziel gesetzt hat, den Raum Universit\u00e4t von links zu (re)politisieren. Gleichzeitig richtet sie sich an j\u00fcngere Semester und will diesen den Einstieg ins Studi-Leben erleichtern, sowie die M\u00f6glichkeit er\u00f6ffnen, an der Uni politisch aktiv zu werden.</p>\n","summary":"","format":"rich_text_editor"},"category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/8463bb01-e974-4785-9c2d-b95d87c9ee2d","id":"8463bb01-e974-4785-9c2d-b95d87c9ee2d","resource":"taxonomy_term","name":"discussion/presentation"}],"date_time":[{"value":"1604419200","value2":"1604419200","duration":0,"time_start":"2020-11-03T17:00:00+01:00","time_end":"2020-11-03T17:00:00+01:00","rrule":null}],"image":[],"price":null,"link":[],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/6bf1f990-a54c-4b5e-a729-67ea5b55d27f","id":"6bf1f990-a54c-4b5e-a729-67ea5b55d27f","resource":"location","title":"online    Berlin Deutschland"}],"phone":null,"topic":[],"title":"Online-Veranstaltung: Solidarisch ernten? Arbeitsausbeutung auf deutschen \u00c4ckern ","language":"de","url":"https://radar.squat.net/en/node/376691","created":"1602849521","uuid":"c7889ae2-72a7-4db6-bf71-638d899b7589"}},"count":2,"facets":{"city":[{"filter":"Berlin","count":2,"formatted":"Berlin"}],"country":[{"filter":"DE","count":2,"formatted":"DE"}],"date":[{"filter":"1604415600","count":1,"formatted":"1604415600"},{"filter":"1604419200","count":1,"formatted":"1604419200"}],"price":[{"filter":"free-121","count":2,"formatted":"free"}],"group":[{"filter":"1599","count":2,"formatted":" Stressfaktor"}],"category":[{"filter":"discussion-presentation","count":2,"formatted":"discussion/presentation"}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/search/events.json" {
			t.Errorf("Expected to request '/api/1.2/search/events.json', got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "facets%5Bcity%5D%5B%5D=Berlin&filter%5B~and%5D%5Bsearch_api_aggregation_1%5D%5B~gte%5D=1604358000&filter%5B~or%5D%5Bsearch_api_aggregation_1%5D%5B~lte%5D=1604419995" {
			t.Errorf("Expected to request raw query 'facets%%5Bcity%%5D%%5B%%5D=Berlin&filter%%5B~and%%5D%%5Bsearch_api_aggregation_1%%5D%%5B~gte%%5D=1604358000&filter%%5B~or%%5D%%5Bsearch_api_aggregation_1%%5D%%5B~lte%%5D=1604419995', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Filters(CreaterRangeFilter(FilterEventStartDateTime, "1604358000", "1604419995"))
	results, err := radar.SearchEvents(sb)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if results.Count != 2 {
		t.Errorf("Expected count to be 8, got: %d", results.Count)
	}
	if results.Results["376690"].Title != "Online-Vortrag: Einführung in die materialistische Rassismuskritik " {
		t.Errorf("Expected title to be 'Online-Vortrag: Einführung in die materialistische Rassismuskritik ', got: %s", results.Results["376690"].Title)
	}
}
