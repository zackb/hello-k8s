package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zackb/hello-k8s/model"
)

func HandleRecord(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		rec := &model.Record{}
		/*
			bytes, _ := ioutil.ReadAll(r.Body)
			fmt.Printf("Got: %s\n", string(bytes))
		*/
		if err := json.NewDecoder(r.Body).Decode(rec); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ERROR: %s", err)
			return
		}
		jsn, _ := json.Marshal(rec)
		fmt.Fprintf(w, "I saw: %s", string(jsn))
		return
	case "PUT":
	case "GET":
		rec := &model.Record{
			Id:   "Hello",
			Name: "NewName",
		}
		jsn, _ := json.Marshal(rec)
		fmt.Fprintf(w, "%s", string(jsn))
	case "DELETE":
	default:
		fmt.Fprintf(w, "I dont know about %s", r.Method)
	}
}
