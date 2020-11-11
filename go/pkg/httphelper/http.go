package httphelper

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	ForwardedPrefix = "x-forwarded-prefix"
)
// Redirect to a fixed URL
type redirectHandler struct {
	url  string
	code int
}

func (rh *redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Redirect(w, r, rh.url, rh.code)
}

// Redirect generates a redirect header take into account the prefix ForwardedPrefix
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	prefix := r.Header.Get(ForwardedPrefix)

	// Normalize the "/" remove trailing slash from prefix and leading "/" from url.
	//
	if len(prefix) > 1 && strings.HasSuffix(prefix, "/")  {
		log.Debugf("Removing trailing / from prefix %v", prefix)
		prefix = prefix[:len(prefix) -  1]
	}

	if len(url) > 1 && strings.HasPrefix(url, "/") {
		log.Debugf("Removing leading / from url %v", url)
		url = url[1:len(url)]
	}
	log.Debugf("Header %v ; Value: %v", ForwardedPrefix, prefix)
	url = prefix + "/" + url

	log.Debugf("Redirect URL: %v", url)
	http.Redirect(w, r, url,code)
}

// RelativeRedirectHandler returns a request handler that redirects
// each request it receives to the given url using the given
// status code.
//
// It is identical to the built in http.RedirectHandler except it uses the header
// x-forwarded-prefix to add a forwarding prefix to URls
//
// This only makes sense if url is a relative path.
//
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func RelativeRedirectHandler(url string, code int) http.Handler {
	return &redirectHandler{url, code}
}