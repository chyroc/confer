package loader

import (
	"fmt"

	"github.com/chyroc/go-loader/internal"
	"github.com/chyroc/go-loader/internal/extractor"
	"github.com/chyroc/go-loader/internal/load"
)

// Load get data from everywhere
//
// Load add some build-in extractor and transformer
//   - extractor: env
//   - extractor: file
func Load(source interface{}, options ...Option) error {
	cli, err := New(
		append([]Option{
			WithExtractor(extractor.NewEnv()),  // extractor: env
			WithExtractor(extractor.NewFile()), // extractor: file
		}, options...)...,
	)
	if err != nil {
		return err
	}

	return load.Load(source, cli.internalOption())
}

// Loader impl this package
type Loader struct {
	tagName      string
	extractors   map[string]Extractor
	transformers map[string]Transformer
}

// Option config how Loader work
type Option func(r *Loader) error

type (
	KeyVal         = internal.KeyVal
	ExtractorReq   = internal.ExtractorReq
	TransformerReq = internal.TransformerReq
	Extractor      = internal.Extractor
	Transformer    = internal.Transformer
)

// WithExtractor add extractor to Loader
func WithExtractor(extractors ...Extractor) Option {
	return func(r *Loader) error {
		for _, v := range extractors {
			if _, ok := r.extractors[v.Name()]; ok {
				return fmt.Errorf("extractor(%q) registed", v.Name())
			}
			r.extractors[v.Name()] = v
		}
		return nil
	}
}

// WithTransform add transform to Loader
func WithTransform(transfers ...Transformer) Option {
	return func(r *Loader) error {
		for _, v := range transfers {
			if _, ok := r.transformers[v.Name()]; ok {
				return fmt.Errorf("transformer(%q) registed", v.Name())
			}
			r.transformers[v.Name()] = v
		}
		return nil
	}
}

// New create new Loader instance
//
// Generally, you don’t need to call this function, just use Load directly.
func New(options ...Option) (*Loader, error) {
	r := &Loader{
		tagName:      "loader",
		extractors:   map[string]Extractor{},
		transformers: map[string]Transformer{},
	}
	for _, v := range options {
		if err := v(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (r *Loader) internalOption() *load.Option {
	resp := &load.Option{
		TagName:      r.tagName,
		Extractors:   map[string]internal.Extractor{},
		Transformers: map[string]internal.Transformer{},
	}

	for k, v := range r.extractors {
		resp.Extractors[k] = v
	}
	for k, v := range r.transformers {
		resp.Transformers[k] = v
	}
	return resp
}
