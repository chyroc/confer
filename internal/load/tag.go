package load

import (
	"github.com/chyroc/go-loader/internal"
	"github.com/chyroc/go-loader/internal/tag_parser"
)

type Tag struct {
	extractorName string
	extractorArgs *internal.ExtractorReq
	// ;
	transformerName string
	transformerArgs *internal.TransformerReq
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
		resp.extractorArgs = new(internal.ExtractorReq)
		for _, v := range res.Extractor.Args {
			resp.extractorArgs.KeyVal = append(resp.extractorArgs.KeyVal, internal.KeyVal{Key: v.Key, Val: v.Val})
		}
	}
	if res.Transformer != nil {
		resp.transformerName = res.Transformer.Name
		resp.transformerArgs = new(internal.TransformerReq)
		for _, v := range res.Transformer.Args {
			resp.transformerArgs.KeyVal = append(resp.transformerArgs.KeyVal, internal.KeyVal{Key: v.Key, Val: v.Val})
		}
	}

	resp.Required = res.Required
	resp.Default = res.Default

	return resp, nil
}
