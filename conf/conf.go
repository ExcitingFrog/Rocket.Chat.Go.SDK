package conf

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func ReadConf() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	data, err := ioutil.ReadFile("application.yml")
	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(data))
}

func GetEnv(env string) string {
	return fmt.Sprintf("%v", viper.Get(env))
}
