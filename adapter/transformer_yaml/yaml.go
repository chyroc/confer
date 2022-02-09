package transformer_yaml

import (
	"github.com/chyroc/go-loader/iface"
)

type Yaml struct{}

func (r *Yaml) Name() string {
	return "yaml"
}

func (r *Yaml) Transform(data string, args *iface.TransformerReq) (interface{}, error) {
	// if len(args) != 1 {
	// 	return "", fmt.Errorf("yaml transformer expect one args")
	// }
	//
	// val, err := yaml_query.QueryYaml([]byte(data), args[0])
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func New() *Yaml {
	return &Yaml{}
}
