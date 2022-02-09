package loader

import (
	"fmt"

	"github.com/chyroc/go-loader/adapter/extractor_env"
	"github.com/chyroc/go-loader/adapter/extractor_file"
	"github.com/chyroc/go-loader/iface"
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
			WithExtractor(extractor_env.NewEnv()),   // extractor: env
			WithExtractor(extractor_file.NewFile()), // extractor: file
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
	extractors   map[string]iface.Extractor
	transformers map[string]iface.Transformer
}

// Option config how Loader work
type Option func(r *Loader) error

// WithExtractor add extractor to Loader
func WithExtractor(extractors ...iface.Extractor) Option {
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
func WithTransform(transfers ...iface.Transformer) Option {
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
		extractors:   map[string]iface.Extractor{},
		transformers: map[string]iface.Transformer{},
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
		Extractors:   map[string]iface.Extractor{},
		Transformers: map[string]iface.Transformer{},
	}

	for k, v := range r.extractors {
		resp.Extractors[k] = v
	}
	for k, v := range r.transformers {
		resp.Transformers[k] = v
	}
	return resp
}
