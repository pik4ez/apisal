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
	http.ServeFile(w, r, "html/index.html")
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

	finalHTMLFileName := "/html/" + rndString + ".html"
	finalHTML, err := createFile("." + finalHTMLFileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer finalHTML.Close()

	cmd.Stdout = finalHTML

	err = cmd.Start()
	if err != nil {
		log.Println(err)
		return
	}
	cmd.Wait()

	http.Redirect(w, r, finalHTMLFileName, 301)
}

func createFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.Handle("/html/", http.FileServer(http.Dir("./")))
	addr := "127.0.0.1:3000"

	print("http://")
	println(addr)
	http.ListenAndServe(addr, nil)
}
