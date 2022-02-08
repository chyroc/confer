package load

import (
	"fmt"
	"strings"

	"github.com/chyroc/go-loader/internal"
	"github.com/chyroc/go-loader/internal/helper"
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
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return nil, fmt.Errorf("tag conf can not be empty")
	}
	parser := &tagParser{StringParser: helper.NewStringParser(tag)}
	return parser.parse()
}

type tagParser struct {
	*helper.StringParser
}

// name,k:v,k2:v2;name,k:v;required,default:x
func (r *tagParser) parse() (resp *Tag, err error) {
	parseKeyArgs := func() (key string, args []internal.KeyVal, err error) {
		key, err = r.PopString(true, []rune{',', ';'}) // `name,`, `name;`
		if err != nil {
			return "", nil, err
		}
		for {
			if err := r.ExpectRune(true, ';'); err == nil {
				r.GoBack(';')
				break
			}
			if err := r.ExpectRune(true, ','); err != nil {
				break
			}
			k, err := r.PopString(true, []rune{',', ';', ':'})
			if err != nil {
				return "", nil, err
			}
			if err := r.ExpectRune(true, ':'); err != nil {
				return "", nil, err
			}
			v, err := r.PopString(true, []rune{',', ';', ':'})
			if err != nil {
				return "", nil, err
			}
			args = append(args, internal.KeyVal{Key: k, Val: v})
		}
		return key, args, nil
	}
	resp = &Tag{extractorArgs: new(internal.ExtractorReq), transformerArgs: new(internal.TransformerReq)}
	r.RemoveSpace()

	// extractor
	resp.extractorName, resp.extractorArgs.KeyVal, err = parseKeyArgs()
	if err != nil {
		return nil, err
	} else if resp.extractorName == "" {
		return nil, fmt.Errorf("expect get extractor name")
	}

	// transformer
	if err := r.ExpectRune(true, ';'); err == nil {
		resp.transformerName, resp.transformerArgs.KeyVal, err = parseKeyArgs()
		if err != nil {
			return nil, err
		}
	}

	// keyword
	if err := r.ExpectRune(true, ';'); err == nil {
		for !r.IsEnd() {
			char, _ := r.GetChar()
			switch char {
			case 'r': // required
				if err := r.ExpectRune(true, []rune("required")...); err != nil {
					return nil, err
				}
				resp.Required = true
			case 'd': // default:
				if err := r.ExpectRune(true, []rune("default")...); err != nil {
					return nil, err
				}
				if err := r.ExpectRune(true, ':'); err != nil {
					return nil, err
				}
				val, err := r.PopString(true, []rune{',', ';'})
				if err != nil {
					return nil, err
				}
				resp.Default = val
			default:
				return nil, fmt.Errorf("unsupport keyword")
			}

			if err := r.ExpectRune(true, ','); err != nil {
				break
			}
		}
	}

	r.RemoveSpace()

	if !r.IsEnd() {
		if r.IsAt(r.Length()-1) && r.GetLeft() == ";" {
			return resp, nil
		}
		return nil, fmt.Errorf("unwanted chars: %s", r.GetLeft())
	}

	return resp, nil
}
