package syobocal

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/syobocal/api"
	"github.com/otiai10/syobocal/factory"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestScenario_001(t *testing.T) {
	srv := mockserver()
	b := api.TitleLookup()
	res, err := http.Get(srv.URL + "?" + b.Build().Encode())
	Require(t, err).ToBe(nil)
	defer res.Body.Close()
	lookup := api.TitleLookupResponse{}
	err = xml.NewDecoder(res.Body).Decode(&lookup)
	Require(t, err).ToBe(nil)
	animes, err := factory.ToAnimeListFromTitleLookup(lookup)
	Expect(t, err).ToBe(nil)

	Expect(t, animes).TypeOf("[]models.Anime")
}

func mockserver() *httptest.Server {
	r := marmoset.NewRouter()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.URL.Query().Get("Command") {
		case "TitleLookup":
			f, err := os.Open(filepath.Join("testdata", "titlelookup_20220817.xml"))
			if err != nil {
				panic(err)
			}
			defer f.Close()
			io.Copy(w, f)
		}
	})
	return httptest.NewServer(r)
}
