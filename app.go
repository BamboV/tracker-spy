package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"os"
	"fmt"
	"errors"
)

func main() {

	torrentURL := os.Getenv("TORRENT_URL")
	if torrentURL == "" {
		panic(errors.New("Variable TORRENT_URL is not set"))
	}

	db, err := gorm.Open("postgres", "host=postgres user=tracker_spy dbname=tracker_spy sslmode=disable password=tracker_spy")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	switch os.Args[1] {
	case "install":
		if !db.HasTable(&Tracker{}) {
			db.CreateTable(&Tracker{})
		}
		break
	case "start":
		repository := DBRepository{DB: db}
		spy := Spy{repository:repository, torrent: TorrentClient{client:http.Client{}, torrentURL:torrentURL}}
		server := SpyServer{}
		server.StartListen(spy)
		break
	case "update":
		repository := DBRepository{DB: db}
		spy := Spy{repository:repository, torrent: TorrentClient{client:http.Client{}, torrentURL:torrentURL}}
		data := spy.UpdateAll()
		if len(data) > 0 {
			fmt.Println("Updated: ")
			for _, value := range data {
				fmt.Println(value.Name)
			}
		}
		break
	}
}
