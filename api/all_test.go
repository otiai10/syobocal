package api

import (
	. "github.com/otiai10/mint"

	"testing"
)

func TestTitleLookup(t *testing.T) {
	b := TitleLookup()
	Expect(t, b.Build().Encode()).ToBe("Command=TitleLookup&TID=%2A")
}
