package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/report"
	"net/http"
)

type Server struct {
	*mux.Router
	ReportServer report.ReportServer
}

// NewServer creates a http instance and starts it
func NewServer(rs report.ReportServer) *Server {
	// Create router
	s := &Server{
		Router:       mux.NewRouter(),
		ReportServer: rs,
	}
	s.HandleFunc("/report", s.handlePostReport).Methods(http.MethodPost)
	s.HandleFunc("/report", s.handleReport).Methods(http.MethodGet)

	return s

}

func (s *Server) handleReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<form action="/report" method="post" enctype="multipart/form-data">
    Select image to upload:
    <input type="file" name="file" id="file">
    <input type="text" name="email" id="email">
    <input type="submit" value="Upload Image" name="submit">
</form>`
	fmt.Fprint(w, html)
}

func (s *Server) handlePostReport(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	rpt := report.NewReport(file)
	rpt.Email = r.FormValue("email")

	err = s.ReportServer.Upload(r.Context(), rpt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	err = s.ReportServer.Notify(r.Context(), rpt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	w.Write([]byte("Success"))

}

// Start starts the http with the supplied port
func (s *Server) ListenAndServe(port string) {
	http.ListenAndServe(port, s.Router)
}
