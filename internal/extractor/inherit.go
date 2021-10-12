package extractor

import (
	"fmt"

	"github.com/chyroc/go-loader/internal"
)

type Inherit struct{}

func NewInherit() *Inherit {
	return &Inherit{}
}

func (r *Inherit) Name() string {
	return "inherit"
}

func (r *Inherit) Extract(args *internal.ExtractorReq) (string, error) {
	format, _ := args.Get("format")
	if format == "" {
		return "", fmt.Errorf("inherit extractor expect `format` args")
	}
	switch format {
	}
	panic("implement me")
}
