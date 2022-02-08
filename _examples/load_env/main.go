package main

import (
	"fmt"

	"github.com/chyroc/go-loader"
)

type Conf struct {
	GitHubToken string `loader:"env,key=GITHUB_TOKEN;;required"`
}

func main() {
	conf := new(Conf)
	err := loader.Load(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("loader:", conf.GitHubToken)
}
