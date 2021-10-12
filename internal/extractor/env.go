package extractor

import (
	"fmt"
	"os"

	"github.com/chyroc/go-loader/internal"
)

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (r *Env) Name() string {
	return "env"
}

// Extract impl Extract for `env`
func (r *Env) Extract(args *internal.ExtractorReq) (string, error) {
	key, _ := args.Get("key")
	if key == "" {
		return "", fmt.Errorf("env extractor expect `key` args")
	}

	val := os.Getenv(key)

	return val, nil
}
