package main

import (
	"net/http"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
)

type TorrentClient struct {
	client http.Client
	torrentURL string
}

func (t *TorrentClient) GetTorrent(tracker string, id uint) Tracker {
	resp, err := t.client.Get(t.torrentURL + "/" + tracker + "/" + strconv.Itoa(int(id)))

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	tr := Tracker{}
	err = json.Unmarshal(body, &tr)

	if err != nil {
		panic(err)
	}

	tr.TrackerID = tr.ID
	tr.ID = id
	tr.Source = tracker

	return tr
}

func (t *TorrentClient) Download(tracker Tracker) bool {
	resp, err := t.client.Post("http://torrent/" + tracker.Source + "/" + strconv.Itoa(int(tracker.TrackerID)),"application/json", strings.NewReader(fmt.Sprintf("{id:%v}", tracker.TrackerID)))

	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}
