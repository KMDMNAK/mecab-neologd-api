package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type Recorder struct {
	w *httptest.ResponseRecorder
	b []byte
}

func testRequest(u string) *Recorder {
	e := createEngine()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", u, nil)
	e.ServeHTTP(w, req)
	b := w.Body.Bytes()
	return &Recorder{
		b: b,
		w: w,
	}
}

type ExtractProperResponse struct {
	ProperNouns []string `json:"proper_nouns"`
}

func TestExtractProperSuccess(t *testing.T) {
	te := "習近平国家主席「説教は決して受け入れない」アメリカをけん制　中国共産党創設100周年"
	text := url.QueryEscape(te)
	r := testRequest("/proper/extract/" + text)
	body := ExtractProperResponse{}
	err := json.Unmarshal(r.b, &body)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if len(body.ProperNouns) == 0 {
		t.Fail()
	}
}

func TestExtractProperFailed(t *testing.T) {
	te := ""
	text := url.QueryEscape(te)
	r := testRequest("/proper/extract/" + text)
	if r.w.Result().StatusCode != 404 {
		t.Fail()
	}
}

type CountProperResponse struct {
	ProperNounsCount map[string]int `json:"proper_nouns_count"`
}

func TestCountProper(t *testing.T) {
	te := "習近平国家主席「説教は決して受け入れない」アメリカをけん制　中国共産党創設100周年"
	text := url.QueryEscape(te)
	r := testRequest("/proper/count/" + text)

	body := CountProperResponse{}
	err := json.Unmarshal(r.b, &body)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	c, ok := body.ProperNounsCount["習近平"]
	if !ok || c < 1 {
		t.Fail()
	}
}
