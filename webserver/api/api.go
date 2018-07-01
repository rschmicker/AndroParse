package api

import(
	"net/http"
	"io"
	"github.com/AndroParse/webserver/utils"
	"log"
	"time"
	"strings"
	"github.com/AndroParse/webserver/query"
)

func All(w http.ResponseWriter, req *http.Request) {
	Query(w, req)
}

func Query(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	args := req.URL.Query()
        toArg, _ := utils.GetArg("to", args)
        fromArg, _ := utils.GetArg("from", args)
	flds, _ := utils.GetArg("fields", args)
        filename := time.Now().Format(time.RFC850) + ".zip"
	filename = strings.Replace(filename, " ", "", -1)
	//http.Error(w, InfoMsg, http.StatusBadRequest)

	log.Println("===============================")
	log.Println("From: " + fromArg)
	log.Println("To: " + toArg)
	log.Println("Fields: ")
	log.Println(flds)
	log.Println("File name: " + filename)

	go query.Query(filename, fromArg, toArg, flds)

	log.Println("===============================")

	io.WriteString(w, "You queried for:\n")
	io.WriteString(w, "From: " + fromArg + "\n")
	io.WriteString(w, "To: " + toArg + "\n")
	io.WriteString(w, "Fields: " + flds + "\n\n")

	io.WriteString(w, "Your query is being processed\n")
	io.WriteString(w, "Your filename is " + filename + "\n")
	io.WriteString(w, "Please check back in an hour\n")
	io.WriteString(w, "This file will be purged in 30 days\n\n")

	io.WriteString(w, "To access your file, use an FTP client and connect to:\n")
	io.WriteString(w, "ftp://64.251.61.74\n")
	io.WriteString(w, "Anonymous read only access is provided as:\n")
	io.WriteString(w, "User: anonymous\n")
	io.WriteString(w, "Password: <empty>\n")
}
