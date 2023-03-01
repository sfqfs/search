package main

import (
	"encoding/json"
	"flag"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"log"
	"os"
)

type AlgoliaUploadConfig struct {
	AppID  string `toml:"appID"`
	APIKey string `toml:"apiKey"`
}

const searchFile = "./public/search.json"

var appID = flag.String("appID", "", "appid")
var apiKey = flag.String("apiKey", "", "appid")
var indexFile = flag.String("file", searchFile, "index file")

func main() {
	flag.Parse()

	client := search.NewClient(*appID, *apiKey)
	index := client.InitIndex("blog")
	file, err := os.Open(*indexFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	var records []interface{}
	if err = json.NewDecoder(file).Decode(&records); err != nil {
		log.Fatalln(err)
	}
	res, err := index.SaveObjects(records, opt.AutoGenerateObjectIDIfNotExist(true))
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(os.Stdout).Encode(res)
}
