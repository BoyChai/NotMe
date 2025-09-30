package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

func init() {
	if _, err := toml.DecodeFile("config.toml", &Global); err != nil {
		log.Fatal(err)
		return
	}
}

var Global config

type Xlsx struct {
	Path    string `toml:"path"`
	Sheet   string `toml:"sheet"`
	Col     string `toml:"col"`
	DescCol string `toml:"desc"`
	PASS    string `toml:"pass"`
	CSV     bool   `toml:"csv"`
}

type Text struct {
	Path  string `toml:"path"`
	Delim string `toml:"delim"`
}

// Config 结构体
type config struct {
	XLSX []Xlsx `toml:"xlsx"`
	TEXT []Text `toml:"text"`
}
