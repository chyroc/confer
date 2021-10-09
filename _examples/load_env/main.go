package main

import (
	"fmt"

	"github.com/chyroc/confer"
)

type Conf struct {
	GitHubToken string `conf:"env,GITHUB_TOKEN"`
}

func main() {
	conf := new(Conf)
	err := confer.Load(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("conf:", conf.GitHubToken)
}
