package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zackb/hello-k8s/db"
	"github.com/zackb/hello-k8s/handlers"
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

	// lewl
	handlers.Init(data)

	handler := http.NewServeMux()

	handler.HandleFunc("/record/", handlers.HandleRecord)

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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

	server := &http.Server{Addr: ":" + PORT, Handler: handler}

	go func() {
		fmt.Println("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("ListenAndServe err: ", err)
		}
	}()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, os.Kill) //syscall.SIGINT, syscall.SIGTERM)

	<-sig

	fmt.Println("Shutting Down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Failed shutting down", err)
	}

}

func getEnv(name string, defaul string) string {
	var val string
	if val = os.Getenv(name); val == "" {
		val = defaul
	}
	return val
}
