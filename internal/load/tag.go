package load

import (
	"github.com/chyroc/go-loader/iface"
	"github.com/chyroc/go-loader/internal/tag_parser"
)

type Tag struct {
	extractorName string
	extractorArgs *iface.ExtractorReq
	// ;
	transformerName string
	transformerArgs *iface.TransformerReq
	// ;
	Required bool
	Default  string
}

func ParseTag(tag string) (*Tag, error) {
	res, err := tag_parser.Parse(tag)
	if err != nil {
		return nil, err
	}
	resp := new(Tag)
	if res.Extractor != nil {
		resp.extractorName = res.Extractor.Name
		resp.extractorArgs = new(iface.ExtractorReq)
		for _, v := range res.Extractor.Args {
			resp.extractorArgs.KeyVal = append(resp.extractorArgs.KeyVal, iface.KeyVal{Key: v.Key, Val: v.Val})
		}
	}
	if res.Transformer != nil {
		resp.transformerName = res.Transformer.Name
		resp.transformerArgs = new(iface.TransformerReq)
		for _, v := range res.Transformer.Args {
			resp.transformerArgs.KeyVal = append(resp.transformerArgs.KeyVal, iface.KeyVal{Key: v.Key, Val: v.Val})
		}
	}

	resp.Required = res.Required
	resp.Default = res.Default

	return resp, nil
}
