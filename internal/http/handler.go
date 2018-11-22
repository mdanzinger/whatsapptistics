package http

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/mdanzinger/whatsapptistics/internal/job"

	"github.com/mdanzinger/whatsapptistics/internal/chat"
	"github.com/mdanzinger/whatsapptistics/internal/report"
)

type handler struct {
	ChatService   chat.ChatService
	ReportService report.ReportService
	Jobsource     job.Source
	templates     *template.Template
}

// Compile templates
func (h *handler) CompileTemplates() {
	h.templates = template.Must(template.ParseGlob("../../internal/web/template/*"))
}

// index returns the site index
func (h *handler) serveIndex(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) serveReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	report, err := h.ReportService.Get(r.Context(), vars["id"])
	if err != nil {
		http.Error(w, "Report failed :(", http.StatusBadRequest)
		return
	}
	reportjson, err := json.Marshal(report)
	if err != nil {
		http.Error(w, "Report failed :(", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s", reportjson)

	//err := h.templates.ExecuteTemplate(w, "index", nil)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}

func (h *handler) newReport(w http.ResponseWriter, r *http.Request) {
	c, _, err := r.FormFile("chat")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ensureFileIsTxt(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprint(w)
		fmt.Println(err)
		return
	}

	e := r.FormValue("email")
	if err = h.ChatService.New(r.Context(), c, e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func ensureFileIsTxt(f io.ReadSeeker) error {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return err
	}

	contentType := http.DetectContentType(buffer)
	if !strings.Contains(contentType, "text/plain") {
		return fmt.Errorf("file not text type: %s", contentType)
	}

	f.Seek(0, 0)
	return nil
}
