package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func Basic() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Batman"
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Superman"

		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/vulncmd", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["key"][0]

		kunciaws := ${{ secrets.AWS_KEY }}
		rahasiaaws := ${{ secrets.AWS_SECRET }}
		fmt.Println(kunciaws, rahasiaaws)

		if !ok || len(keys[0]) < 1 {
			fmt.Println("Url Param 'key' is missing")
			return
		}

		key := keys[0]

		fmt.Println("Url Param 'key' is: " + string(key))
		cmd := exec.Command("/bin/sh", "-c", string(key))
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Fprintf(w, string(stdout))
	})
}
