package tg

import (
	"github.com/naoina/toml"

	"io/ioutil"
	"os"
)

type Config struct {
	TgPubKey string
}

func LoadConfig(filename string, config interface{}) (err error) {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	if err := toml.Unmarshal(buf, config); err != nil {
		panic(err)
	}

	return nil
}
