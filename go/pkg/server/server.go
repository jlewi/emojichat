package server

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/jlewi/emojichat/go/pkg/httphelper"
	"net/http"
	"path"
)

const (
	AccessControlHeader = "Access-Control-Allow-Origin"
)

type Server struct {
	Router *mux.Router
	log logr.Logger
}

// New creates a new server and register routes.
func New(staticFilesPath string, log logr.Logger) (*Server, error) {
	s := &Server{}

	s.Router = mux.NewRouter().StrictSlash(true)

	s.Router.HandleFunc("/healthz", s.Healthz)


	s.log = log

	if staticFilesPath == ""{
		return nil, fmt.Errorf("staticFilesPath can't be empty.")
	}

	log.Info("Configuring static files.", "path", staticFilesPath)
	fs := http.FileServer(http.Dir(path.Join(staticFilesPath)))
	// We need need to strip the prefix from the handler to correctly map the files.
	s.Router.PathPrefix("/ui/").Handler(http.StripPrefix("/ui/", fs))

	// Add an additional handler for js so we can continue to serve js off that path.
	jsFs := http.FileServer(http.Dir(path.Join(staticFilesPath, "js")))
	s.Router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsFs))

	// TODO(jlewi): The redirects are problematic behind a reverse proxy. We need to strip the prefix
	// when rewriting the URI otherwise the PathPrefix rules won't work. But in the redirect case
	// we need to know the base URI when issuing the redirect. one solution would be to use the header
	// x-forwarded-prefix to store the base path. We then need to modify the redirect handler to
	// get the http header and include it.
	// TODO(jlewi): I think ISTIO can setup the redirects directly in the virtual service so maybe we should
	// just rely on ISTIO in the reverse proxy case? And possibly in all cases.
	s.Router.Handle("/ui", httphelper.RelativeRedirectHandler("ui/", http.StatusMovedPermanently))
	// Redirect "/" to the UI
	s.Router.Handle("/", httphelper.RelativeRedirectHandler("ui/", http.StatusMovedPermanently))

	return s, nil
}

type responseError struct {
	Message string `json:"message"`
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	// TODO(jlewi): Is using "*" overly permissive
	w.Header().Set(AccessControlHeader, "*")
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}
