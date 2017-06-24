package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"net/http"
)

type SpyServer struct {
}

func (s *SpyServer) StartListen(spy Spy) {
	http.HandleFunc("/trackers/", func(w http.ResponseWriter,r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			data := spy.GetSpyList()
			str, _ := json.Marshal(data)
			w.Write(str)
			break
		case http.MethodPost:
			id, _ := strconv.Atoi(r.PostForm.Get("id"))
			tracker := Tracker{
				Name: r.PostForm.Get("name"),
				Source: r.PostForm.Get("source"),
				TrackerID: uint(id),
			}
			data := spy.AddToSpyList(tracker)
			str, _ := json.Marshal(data)
			w.Write(str)
			break
		case http.MethodDelete:
			id ,_ := strconv.Atoi(strings.Trim(r.URL.Path, "/trackers/"))
			spy.Delete(uint(id))
			break
		}
	})

	http.HandleFunc("/update/", func(w http.ResponseWriter,r *http.Request) {
		data := spy.UpdateAll()
		str, _ := json.Marshal(data)
		w.Write(str)
	})

	http.ListenAndServe(":80", nil)
}
