package helper

import (
	"fmt"
	"strings"
)

type aequalbParser struct {
	data []rune
	idx  int
}

func ParseAEqualB(s string) (string, string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", "", nil
	}
	parser := &aequalbParser{data: []rune(s), idx: 0}
	return parser.parse()
}

// ` a = b `
// ` a = ' b ' `
func (r *aequalbParser) parse() (a string, b string, err error) {
	r.removeSpace()

	a, err = r.parseString()
	if err != nil {
		return a, b, err
	}
	// fmt.Println(a, string(r.data))

	if err := r.findRune(true, '='); err != nil {
		return a, b, err
	}

	b, err = r.parseString()
	if err != nil {
		return a, b, err
	}

	return a, b, nil
}

func (r *aequalbParser) parseString() (string, error) {
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
			return string(res), nil
		case !quoteFound && (d == ' ' || d == '='):
			return string(res), nil
		default:
			res = append(res, r.data[r.idx])
			r.idx++
		}
	}
	if quoteFound {
		return "", fmt.Errorf("expect end with quota(%s)", string([]rune{quoteRune}))
	}
	return string(res), nil
}

func (r *aequalbParser) findRune(isKey bool, rs ...rune) error {
	if r.idx >= len(r.data) {
		return fmt.Errorf("reach end of data, `%s` cannot found", string(rs))
	}
	if isKey {
		r.removeSpace()
	}
	c := 0
	for i := r.idx; i < len(r.data) && i-r.idx >= 0 && i-r.idx < len(rs); i++ {
		if r.data[i] == rs[i-r.idx] {
			c++
			continue
		}
		return fmt.Errorf("expect: `%s`, but got: `%s`", string(rs), string(r.data[r.idx:Min(r.idx+len(rs), len(r.data))]))
	}
	r.idx += c
	if isKey {
		r.removeSpace()
	}
	return nil
}

func (r *aequalbParser) removeSpace() (n int) {
	for i := r.idx; i < len(r.data); i++ {
		if r.data[i] != ' ' {
			return
		}
		r.idx++
		n++
	}
	return
}
