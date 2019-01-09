package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type questEmployee struct {
	First string
	Last  string
	Dept  string
}

var employees []questEmployee

func main() {
	csvImport("employees.csv")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func csvImport(in string) {
	f, err := os.Open(in)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	emps, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, v := range emps {
		e := questEmployee{
			First: v[0],
			Last:  v[1],
			Dept:  v[2],
		}

		employees = append(employees, e)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		panic(err)
	}
	t.Execute(res, employees)
	fmt.Println("Served")
}
