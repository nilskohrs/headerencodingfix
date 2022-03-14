package headerencodingfix_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nilskohrs/headerencodingfix"
)

func TestShouldEncodeHeader(t *testing.T) {
	cfg := headerencodingfix.CreateConfig()
	ctx := context.Background()

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := headerencodingfix.New(ctx, next, cfg, "headerEncodingFix")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/admin/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Test-Header", "PyrÉnÉes")

	handler.ServeHTTP(recorder, req)
	encoded := string([]byte{80, 121, 114, 201, 110, 201, 101, 115})
	if req.Header.Get("Test-Header") != encoded {
		t.Error("Header value has not been encoded correctly")
	}
}
