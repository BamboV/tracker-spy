package tracker_spy

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
			decoder := json.NewDecoder(r.Body)
			tracker := Tracker{}
			err := decoder.Decode(&tracker)

			if err != nil || tracker.Name == "" || tracker.Source == "" || tracker.TrackerID == 0 {
				w.WriteHeader(400)
				return
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
