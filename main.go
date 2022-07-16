package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Mem struct {
	Id       string
	Title    string
	MarkDown string
	Tags     []string
	Created  string
	Updated  string
}

func main() {

	content, err := ioutil.ReadFile("./mem.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var mems []Mem
	err = json.Unmarshal(content, &mems)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	dirName := "Notes"
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(mems); i++ {
		mem := mems[i]
		file, err := os.Create(path.Join(dirName, mem.Title+".md"))
		tagString := strings.Join(mem.Tags, ",") + "\n"
		_, ferr := file.WriteString(tagString)

		if ferr != nil {
			log.Fatal(err)
		}

		file.WriteString(mem.MarkDown)
		if err != nil {
			log.Fatal(err)
		}
	}
}
