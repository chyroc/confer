package extractor

import (
	"fmt"
	"os"
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
	if len(args) != 1 {
		return "", fmt.Errorf("env extractor expect one args")
	}

	val := os.Getenv(args[0])

	return val, nil
}
