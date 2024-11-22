package main

import (
	"fmt"

	"github.com/CHlluanma/go-pkg/config"
)

type testConfig struct {
	Env       string     `mapstructure:"env"`
	PkgConfig *PkgConfig `mapstructure:"pkgconfig"`
}

type PkgConfig struct {
	Name    string   `mapstructure:"name"`
	Version string   `mapstructure:"version"`
	Use     []string `mapstructure:"use"`
}

func main() {
	var conf testConfig
	c := config.New(config.WithUnmarshalStruct(&conf)) // config.WithEnvFilePath("./config/dev"),

	c.LoadConfig()

	fmt.Printf("%v\n", c.GetEnv("title_dotenv"))
}
