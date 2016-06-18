package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"github.com/jmcvetta/randutil"
)

const (
	GPX2HTML_COMMAND = "./gpx2html.sh"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/index.html")
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("gpx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	rndString, err := randutil.String(8, randutil.Alphabet)
	if err != nil {
		log.Println(err)
		return
	}

	gpxFileName := "./www/tmp/" + rndString + handler.Filename

	gpxFile, err := createFile(gpxFileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer gpxFile.Close()

	io.Copy(gpxFile, file)

	cmd := exec.Command(GPX2HTML_COMMAND, gpxFileName)

	finalHtmlFileName := "/rendered/"+rndString+"stub.html"
	finalHtml, err := createFile("./www" + finalHtmlFileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer finalHtml.Close()

	cmd.Stdout = finalHtml

	err = cmd.Start()
	if err != nil {
		log.Println(err)
		return
	}
	cmd.Wait()

	http.Redirect(w, r, finalHtmlFileName, 301)
}

func createFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
    http.Handle("/rendered/", http.FileServer(http.Dir("./www")))
	addr := "127.0.0.1:3000"

	print("http://")
	println(addr)
	http.ListenAndServe(addr, nil)
}
