package transformer_yaml

import (
	"fmt"
	"reflect"

	"github.com/chyroc/go-loaders/transformer/transformer_yaml/yaml_query"
)

type Yaml struct{}

func (r *Yaml) Name() string {
	return "yaml"
}

func (r *Yaml) Transform(data string, args []string, typ reflect.Type) (interface{}, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("yaml transformer expect one args")
	}

	val, err := yaml_query.QueryYaml([]byte(data), args[0])
	if err != nil {
		return nil, err
	}

	return val, nil
}

func New() *Yaml {
	return &Yaml{}
}
