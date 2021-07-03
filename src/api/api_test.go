package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type Recorder struct {
	w *httptest.ResponseRecorder
	b []byte
}

func testRequest(method string, u string, body io.Reader) *Recorder {
	e := createEngine()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(method, u, body)
	req.Header.Set("Content-Type", "application/json")
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
	r := testRequest("GET", "/proper/extract/"+text, nil)
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
	r := testRequest("GET", "/proper/extract/"+text, nil)
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
	r := testRequest("GET", "/proper/count/"+text, nil)

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

func TestCountProperPost(t *testing.T) {
	te := map[string]string{"text": "習近平国家主席「説教は決して受け入れない」アメリカをけん制　中国共産党創設100周年"}
	// text := url.QueryEscape(te)
	bytes, err := json.Marshal(te)
	r := testRequest("POST", "/proper/count", strings.NewReader(string(bytes)))
	body := CountProperResponse{}
	err = json.Unmarshal(r.b, &body)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	c, ok := body.ProperNounsCount["習近平"]
	if !ok || c < 1 {
		t.Fail()
	}
}
