package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zackb/hello-k8s/db"
)

func main() {
	fmt.Println("Startup")

	var PORT string
	var MSG string
    var ENGINE string

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	if MSG = os.Getenv("MESSAGE"); MSG == "" {
		MSG = "NONE"
	}

	if ENGINE = os.Getenv("ENGINE"); ENGINE == "" {
		ENGINE = "mem"
	}

    var data db.Db
    if ENGINE == "redis" {
	    data = db.NewRedisDb()
    } else {
	    data = db.NewMemDb()
    }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s: Engine: %s\n", MSG, data.Name())

		key := r.URL.Path

		var err error
		var value int
		if value, err = data.Get(key); err != nil {
			fmt.Fprintf(w, "%s: Error: %s\n", MSG, err.Error())
		}
		fmt.Fprintf(w, "Count: %d for %s \n", value, key)
		data.Set(key, value+1)
	})

	http.ListenAndServe(":"+PORT, nil)

	fmt.Println("Initialized")
}
