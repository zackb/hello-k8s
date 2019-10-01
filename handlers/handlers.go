package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/zackb/hello-k8s/db"
	"github.com/zackb/hello-k8s/model"
)

var data db.Db

func Init(dataStore db.Db) {
	data = dataStore
}

func HandleRecord(w http.ResponseWriter, r *http.Request) {
	id := ""
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) > 1 {
		id = parts[2]
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		rec := &model.Record{}
		if err := json.NewDecoder(r.Body).Decode(rec); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ERROR: %s", err)
			return
		}
		if err := set(rec.Id, rec); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ERROR: %s", err)
			return
		}
		fmt.Printf("Saving: %s\n", rec.Id)
		return
	case "PUT":
	case "GET":
		rec, _ := get(id)
		jsn, _ := json.Marshal(rec)
		fmt.Fprintf(w, "%s", string(jsn))
	case "DELETE":
	default:
		fmt.Fprintf(w, "I dont know about %s", r.Method)
	}
}

func get(id string) (*model.Record, error) {
	bytes, err := data.GetBytes(id)
	if err != nil {
		return nil, err
	}
	record := &model.Record{}
	err2 := json.Unmarshal(bytes, record)
	return record, err2
}

func set(id string, record *model.Record) error {
	bytes, err := json.Marshal(record)
	if err == nil {
		data.SetBytes(id, bytes)
	}
	return err
}
