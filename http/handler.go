package http

import "net/http"

type Handler struct {
	ChatHandler   *ChatHandler
	ReportHandler *ReportHandler
}

// index returns the site index
func (*Handler) index(w http.ResponseWriter, r *http.Request) {

}
