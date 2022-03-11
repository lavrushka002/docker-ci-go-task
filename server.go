package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Text struct {
	gorm.Model
	TS   int64
	Line string
}

func AddLine(text string) {
	temp := Text{
		TS:   time.Now().Unix(),
		Line: text,
	}
	db.Create(&temp)
}

func Input(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		var Lines []Text
		db.Find(&Lines)
		fmt.Println(Lines)
		t, _ := template.ParseFiles("input.html")
		t.Execute(w, Lines)
	} else {
		r.ParseForm()
		AddLine(r.Form["text"][0])
		fmt.Println("text:", r.Form["text"])
		redirect_url := r.FormValue("redir")
		http.Redirect(w, r, redirect_url, 302)
	}
}

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=password")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Text{})
	http.HandleFunc("/", Input)
	log.Fatal(http.ListenAndServe(":80", nil))
}
