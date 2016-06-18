package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/pik4ez/apisal/apisal"
)

func main() {
	var inputObjOne = []byte("{\"point\": {\"lat\": 10.0, \"lng\": 10.0, \"time\": \"2016-06-17T20:30:25+0000\"}, \"title\": \"obj_one_title\", \"description\": \"obj_one_description\", \"images\": [{\"w\": 100, \"h\": 50, \"title\": \"img_one_title\", \"url\": \"http://example.org/image.png\"}]}")
	var objs []apisal.Object
	var objOne apisal.Object
	json.Unmarshal(inputObjOne, &objOne)
	objs = append(objs, objOne)
	fmt.Println(objOne.Images[0])
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(os.Stdout, objs)
}
