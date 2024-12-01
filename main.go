package main

import (
	"cv_gen/model/fsutils"
	"cv_gen/model/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Print("Lancio compilaizone CV")

	// 1 leggere il template
	tmpl, err := template.ParseFiles("templates/cv_template_no_photo.html")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	// 2 leggere i dati
	input_data, err := os.ReadFile("cv_data_gen.json")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	var cv model.CV
	err = json.Unmarshal(input_data, &cv)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3 compilare il template
	destHtml := fsutils.CreateOrClearFile("output.html")
	tmpl.Execute(destHtml, cv)

	// 4 salvare un html
	fmt.Print("FINE CV")

}
