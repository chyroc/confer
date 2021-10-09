package confer

import (
	"fmt"
	"os"

	"github.com/chyroc/confer/internal"
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
	// 1 or 2
	if len(args) != 1 && len(args) != 2 {
		return "", fmt.Errorf("env loader expect one or two args")
	}

	val := os.Getenv(args[0])

	if len(args) == 2 {
		a, b, err := internal.ParseAEqualB(args[1])
		if a != "default" {
			return "", fmt.Errorf("env loader second args expect default=<val>")
		}
		if err != nil {
			return "", err
		}
		if val == "" {
			val = b
		}
	}
	return val, nil
}
