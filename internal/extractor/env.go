package extractor

import (
	"fmt"
	"os"

	"github.com/chyroc/go-loader/internal/helper"
)

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (r *Env) Name() string {
	return "env"
}

// Extract impl Extract for `env`
func (r *Env) Extract(args []string) (string, error) {
	// 1 or 2
	if len(args) != 1 && len(args) != 2 {
		return "", fmt.Errorf("env extractor expect one or two args")
	}

	val := os.Getenv(args[0])

	if len(args) == 2 {
		a, b, err := helper.ParseAEqualB(args[1])
		if a != "default" {
			return "", fmt.Errorf("env extractor second args expect default=<val>")
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
