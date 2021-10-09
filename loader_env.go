package confer

import (
	"fmt"
	"os"
)

type loaderEnv struct{}

func newLoaderEnv() Loader {
	return &loaderEnv{}
}

func (r *loaderEnv) Name() string {
	return "env"
}

// Load impl Loader for `env`
func (r *loaderEnv) Load(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("env loader expect one args")
	}
	return os.Getenv(args[0]), nil
}
