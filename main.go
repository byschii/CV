package main

import (
	"cv_gen/model/fsutils"
	"cv_gen/model/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	fmt.Println("Lancio compilaizone CV")

	// 0 delete "oputput.html" and "output.pdf"
	os.RemoveAll("output.html")
	os.RemoveAll("output.pdf")

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
	err = tmpl.Execute(destHtml, cv)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	// 4 salvare un html
	fmt.Println("FINE CV")

	// 5 launch browser with "python to_html.py"
	err = exec.Command("python", "to_html.py").Run()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

}
