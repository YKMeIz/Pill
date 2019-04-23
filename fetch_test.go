package pill

import (
	"net/http"
	"testing"
)

func TestFetchNewRequest(t *testing.T) {
	b, err := fetch("://", nil)

	if b != nil {
		t.Error(`fetch() returned result while it shouldn't'.`)
	}

	if _, e := http.NewRequest("GET", "://", nil); e.Error() != err.Error() {
		t.Error("fetch() return wrong http.NewRequest error:\n expect:\n", e, "\n get error message:\n", err)
	}
}

func TestFetchDo(t *testing.T) {
	b, err := fetch("", nil)

	if b != nil {
		t.Error(`fetch() returned result while it shouldn't'.`)
	}

	if _, e := http.Get(""); e.Error() != err.Error() {
		t.Error("fetch() return wrong http.NewRequest error:\n expect:\n", e, "\n get error message:\n", err)
	}
}
