package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiKey string `yaml:"key"`
}

type application struct {
	riotClient *golio.Client
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Itemizer"))
}

func (app *application) showChampion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.NotFound(w, r)
		return
	}
	name = strings.Title(strings.ToLower(name))
	fmt.Println(name)
	champion, err := app.riotClient.DataDragon.GetChampion(name)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(champion.Lore))
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := golio.NewClient(
		config.ApiKey,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New()),
	)

	var app application
	app.riotClient = client

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/champion", app.showChampion)

	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func getConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("./private/config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
