package load

import (
	"fmt"
	"strings"

	"github.com/chyroc/go-loader/internal/helper"
)

type Tag struct {
	extractorName string
	extractorArgs []string
	// ;
	transformerName string
	transformerArgs []string
	// ;
	Required bool
	Default  string
}

func ParseTag(tag string) (*Tag, error) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return nil, fmt.Errorf("tag conf can not be empty")
	}
	parser := &tagParser{data: []rune(tag), idx: 0}
	return parser.parse()
}

type tagParser struct {
	data []rune
	idx  int
}

// name,arg1,arg2;name,arg3,arg4;required,default=x
func (r *tagParser) parse() (resp *Tag, err error) {
	parseKeyArgs := func() (key string, args []string, err error) {
		key, err = r.parseString()
		if err != nil {
			return "", nil, err
		}
		for {
			if err := r.findRune(true, ','); err != nil {
				break
			}
			arg, err := r.parseString()
			if err != nil {
				return "", nil, err
			}
			args = append(args, arg)
		}
		return key, args, nil
	}
	resp = new(Tag)
	r.removeSpace()

	// extractor
	resp.extractorName, resp.extractorArgs, err = parseKeyArgs()
	if err != nil {
		return nil, err
	} else if resp.extractorName == "" {
		return nil, fmt.Errorf("expect get extractor name")
	}

	// transformer
	if err := r.findRune(true, ';'); err == nil {
		resp.transformerName, resp.transformerArgs, err = parseKeyArgs()
		if err != nil {
			return nil, err
		}
	}

	// keyword
	if err := r.findRune(true, ';'); err == nil {
		for r.idx < len(r.data) {
			// lookup keyword
			switch r.data[r.idx] {
			case 'r': // required
				if err := r.findRune(true, []rune("required")...); err != nil {
					return nil, err
				}
				resp.Required = true
			case 'd': // default=
				if err := r.findRune(true, []rune("default")...); err != nil {
					return nil, err
				}
				if err := r.findRune(true, '='); err != nil {
					return nil, err
				}
				val, err := r.parseString()
				if err != nil {
					return nil, err
				}
				resp.Default = val
			default:
				return nil, fmt.Errorf("unsupport keyword")
			}

			if err := r.findRune(true, ','); err != nil {
				break
			}
		}
	}

	r.removeSpace()

	// expect end of data
	if r.idx < len(r.data) {
		if r.idx == len(r.data)-1 && r.data[r.idx] == ';' {
		} else {
			return nil, fmt.Errorf("unwanted chars: %s", string(r.data[r.idx:len(r.data)]))
		}
	}

	return resp, nil
}

func (r *tagParser) parseString() (string, error) {
	var quoteRune rune = 0
	quoteFound := false
	if r.findRune(false, '"') == nil {
		quoteFound = true
		quoteRune = '"'
	}
	if !quoteFound && r.findRune(false, '\'') == nil {
		quoteFound = true
		quoteRune = '\''
	}

	res := []rune{}
	for r.idx < len(r.data) {
		d := r.data[r.idx]
		switch {
		case d == '\\':
			r.idx++
			if r.idx >= len(r.data) {
				return "", fmt.Errorf("no char found after the escape char")
			}
			res = append(res, r.data[r.idx])
			r.idx++
		case quoteFound && d == quoteRune:
			r.idx++
			return string(res), nil // 不能 trim-space
		// case !quoteFound && d == ' ':
		// 	return string(res), nil
		case !quoteFound && (d == ',' || d == ';'):
			return strings.TrimSpace(string(res)), nil
		default:
			res = append(res, r.data[r.idx])
			r.idx++
		}
	}
	if quoteFound {
		return "", fmt.Errorf("expect end with quota(%s)", string([]rune{quoteRune}))
	}
	return strings.TrimSpace(string(res)), nil
}

func (r *tagParser) findRune(isKey bool, rs ...rune) error {
	if isKey {
		r.removeSpace()
	}
	if r.idx >= len(r.data) {
		return fmt.Errorf("`%s` cannot found", string(rs))
	} else if r.idx+len(rs) > len(r.data) {
		return fmt.Errorf("`%s` cannot found", string(rs))
	}
	c := 0
	for i := r.idx; i < len(r.data) && i-r.idx >= 0 && i-r.idx < len(rs); i++ {
		if r.data[i] == rs[i-r.idx] {
			c++
			continue
		}
		return fmt.Errorf("expect: %s, but got: %s", string(rs), string(r.data[r.idx:helper.Min(r.idx+len(rs), len(r.data))]))
	}
	r.idx += c
	if isKey {
		r.removeSpace()
	}
	return nil
}

func (r *tagParser) removeSpace() (n int) {
	for i := r.idx; i < len(r.data); i++ {
		if r.data[i] != ' ' {
			return
		}
		r.idx++
		n++
	}
	return
}
