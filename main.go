package main

import (
    "html/template"
    "net/http"
    "strconv"
)

var game = NewGame()

func main() {
    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/play", handlePlay)
    http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, game)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        colStr := r.FormValue("column")
        col, _ := strconv.Atoi(colStr)
        game.Play(col)
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

 
