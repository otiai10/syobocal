package api

import (
	m "github.com/otiai10/mint"

	"testing"
)

func TestTitleLookup(t *testing.T) {
	b := TitleLookup()
	m.Expect(t, b.Build().Encode()).ToBe("Command=TitleLookup&TID=%2A")
}
