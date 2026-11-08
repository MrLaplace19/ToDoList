package webui

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmbeddedUI(t *testing.T) {
	for _, tc := range []struct {
		path, contentType, contains string
		status                      int
	}{
		{"/", "text/html", "Порядок", 200},
		{"/app.js", "javascript", "/api/v1", 200},
		{"/styles.css", "text/css", "@media", 200},
		{"/missing", "text/plain", "404", 404},
	} {
		t.Run(tc.path, func(t *testing.T) {
			rr := httptest.NewRecorder()
			Handler().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, tc.path, nil))
			if rr.Code != tc.status || !strings.Contains(rr.Header().Get("Content-Type"), tc.contentType) || !strings.Contains(rr.Body.String(), tc.contains) {
				t.Fatalf("unexpected response: %d %s %s", rr.Code, rr.Header().Get("Content-Type"), rr.Body.String())
			}
		})
	}
}
