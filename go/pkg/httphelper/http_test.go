package httphelper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test different URL formats and schemes
func TestRedirect(t *testing.T) {


	var tests = []struct {
		forwardedPrefix string
		in   string
		want string
	}{
		{"/prefix", "/foobar.com/baz", "/prefix/foobar.com/baz"},
		{"","/foobar.com/baz", "/foobar.com/baz"},
	}

	for _, tt := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "http://example.com/qux/", nil)
		req.Header.Add(ForwardedPrefix, tt.forwardedPrefix)
		Redirect(rec, req, tt.in, 302)
		if got, want := rec.Code, 302; got != want {
			t.Errorf("Redirect(%q) generated status code %v; want %v", tt.in, got, want)
		}
		if got := rec.Header().Get("Location"); got != tt.want {
			t.Errorf("Redirect(%q) generated Location header %q; want %q", tt.in, got, tt.want)
		}
	}
}
