package reports

import (
	"fmt"
	"net/http"

	"os"

	"strings"

	"whatsapp/app/common"

	"github.com/gorilla/mux"
	//"github.com/mdanzinger/whatsappold/mywhatsapp/common"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/report", addReport).Methods("POST")
	router.HandleFunc("/report", uploadReport).Methods("GET")
	router.HandleFunc("/report/{report_id}", viewReport).Methods("GET")
	router.HandleFunc("/report/{report_id}", deleteReport).Methods("DEL")
}

func uploadReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<form action="/report" method="post" enctype="multipart/form-data">
    Select image to upload:
    <input type="file" name="file" id="file">
    <input type="text" name="email" id="email">
    <input type="submit" value="Upload Image" name="submit">
</form>`
	fmt.Fprint(w, html)
}

func addReport(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	email := r.FormValue("email")
	fmt.Println("Email: " + email)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Validations
	tempBuff := make([]byte, 512)

	_, err = file.Read(tempBuff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filetype := http.DetectContentType(tempBuff)

	if !strings.Contains(filetype, "text/plain") {
		http.Error(w, "Error: file format not .txt. Please re-export the chat and try again.", http.StatusNotAcceptable)
		return
	}

	key, err := common.AddFileToS3(common.AWS_SESS, file, header)
	if err != nil {
		fmt.Println(err)
		return
	}

	common.SendSNSMessage(common.AWS_SESS, key, email)
	defer file.Close()
}

func viewReport(w http.ResponseWriter, r *http.Request) {
}

func deleteReport(w http.ResponseWriter, r *http.Request) {
}
