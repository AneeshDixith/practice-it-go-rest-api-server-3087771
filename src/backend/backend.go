package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_"github.com/mattn/go-sqlite3"
)

type App struct {
	DB *sql.DB
	Port string
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func (a *App) Initialize() {
	var prev *sql.DB
	prev, err := sql.Open("sqlite3", "../../practiceit.db")
	
	if err != nil {
		log.Fatal(err.Error())
		prev.Close()
	}
}

func (a *App) Run() {
	http.HandleFunc("/", helloworld)
	fmt.Println("Server started and listening on port ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}
