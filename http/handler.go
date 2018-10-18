package http

import (
	"github.com/mdanzinger/mywhatsapp/chat"
	"github.com/mdanzinger/mywhatsapp/report"
	"html/template"
	"net/http"
)

type handler struct {
	ChatService   chat.ChatService
	ReportService report.ReportService
	templates     *template.Template
}

// Compile templates
func (h *handler) CompileTemplates() {
	h.templates = template.Must(template.ParseGlob("../../web/template/*"))
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
	err := h.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) newReport(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
