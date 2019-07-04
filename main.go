package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zackb/hello-k8s/db"
	"github.com/zackb/hello-k8s/handler"
)

func main() {
	fmt.Println("Startup")

	var PORT string
	var MSG string
	var ENGINE string

	PORT = getEnv("PORT", "8080")
	MSG = getEnv("MESSAGE", "NONE")
	ENGINE = getEnv("ENGINE", "mem")

	var data db.Db
	if ENGINE == "redis" {
		data = db.NewRedisDb()
	} else {
		data = db.NewMemDb()
	}

	http.HandleFunc("/record", handler.HandleRecord)

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

func getEnv(name string, defaul string) string {
	var val string
	if val = os.Getenv(name); val == "" {
		val = defaul
	}
	return val
}
