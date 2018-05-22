package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Payload - what is sent to the frontend..
type Payload struct {
	Status    int    `json:"status"`
	Data      []Info `json:"data"`
	Msg       string `json:"msg,omitempty"`
	Processes int    `json:"processes"`
}

// Info - the info we are sending back
type Info struct {
	Count     int    `json:"count"`
	EocStatus string `json:"eocStatus"`
	Time      string `json:"time"`
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	resp, err := json.MarshalIndent("Hello, World!", "", "  ")
	check(err)

	fmt.Fprintf(w, string(resp))
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	db, on := getDB()
	var ret []Info
	var err error

	if on {
		rows, err := db.Query("")

		defer rows.Close()

		if err != nil {
			fmt.Println("Error while retrieving info...")
		} else {
			for rows.Next() {
				var info Info
				rows.Scan(&info.Count, &info.EocStatus)
				info.Time = getTime()
				ret = append(ret, info)
			}
		}
	}

	var payload Payload
	payload.Processes, err = GetNumOfProcesses()

	if err == nil {
		payload.Data = ret
		payload.Status = 200
	} else {
		payload.Msg = err.Error()
		payload.Status = 500
	}

	resp, err := json.MarshalIndent(payload, "", "  ")
	check(err)
	fmt.Fprintf(w, string(resp))
}

func killProcesses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	err := KillAllProcesses()

	payload := Payload{
		200,
		nil,
		"Processes were killed",
		0,
	}

	if err != nil {
		payload.Status = 500
		payload.Msg = err.Error()
	}

	resp, err := json.MarshalIndent(payload, "", "  ")
	check(err)
	fmt.Fprintf(w, string(resp))
}

func resetTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	db, on := getDB()
	var err error

	if on {
		_, err = db.Exec("")
	}

	payload := Payload{
		200,
		nil,
		"DB Reset",
		-1,
	}

	if err != nil {
		payload.Msg = err.Error()
		payload.Status = 500
	}

	resp, err := json.MarshalIndent(payload, "", "  ")
	check(err)
	fmt.Fprintf(w, string(resp))
}

func resetFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	err := ResetFiles()

	payload := Payload{
		200,
		nil,
		"DB Reset",
		-1,
	}

	if err != nil {
		payload.Msg = err.Error()
		payload.Status = 500
	}

	resp, err := json.MarshalIndent(payload, "", "  ")
	check(err)
	fmt.Fprintf(w, string(resp))
}
