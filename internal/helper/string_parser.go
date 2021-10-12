package helper

import (
	"fmt"
	"strings"
)

type StringParser struct {
	idx  int
	data []rune
}

func NewStringParser(data string) *StringParser {
	return &StringParser{data: []rune(data), idx: 0}
}

func (r *StringParser) IsEnd() bool {
	return r.idx >= len(r.data)
}

func (r *StringParser) IsAt(idx int) bool {
	return r.idx == idx
}

func (r *StringParser) Length() int {
	return len(r.data)
}

func (r *StringParser) GetChar() (rune, bool) {
	if r.idx < len(r.data) {
		return r.data[r.idx], true
	}
	return 0, false
}

// allowQuote 是否允许引号
// endChar 结束字符
func (r *StringParser) PopString(allowQuote bool, endChar []rune) (string, error) {
	var quoteRune rune = 0
	quoteFound := false
	if allowQuote {
		if r.ExpectRune(false, '"') == nil {
			quoteFound = true
			quoteRune = '"'
		}
		if !quoteFound && r.ExpectRune(false, '\'') == nil {
			quoteFound = true
			quoteRune = '\''
		}
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
		case !quoteFound && isInRune(d, endChar):
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

func (r *StringParser) ExpectRune(isKey bool, rs ...rune) error {
	if isKey {
		r.RemoveSpace()
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
		return fmt.Errorf("expect: `%s`, but got: `%s`", string(rs), string(r.data[r.idx:min(r.idx+len(rs), len(r.data))]))
	}
	r.idx += c
	if isKey {
		r.RemoveSpace()
	}
	return nil
}

func (r *StringParser) RemoveSpace(emptyChar ...rune) (n int) {
	if len(emptyChar) == 0 {
		emptyChar = []rune{' '}
	}
	for i := r.idx; i < len(r.data); i++ {
		if !isInRune(r.data[i], emptyChar) {
			return
		}
		r.idx++
		n++
	}
	return
}

func (r *StringParser) GoBack(char rune) {
	for {
		if r.idx == 0 {
			panic("0 cannot go back")
		}
		r.idx--
		if r.idx == 0 {
			if r.data[r.idx] == char {
				break
			}
		} else {
			if r.data[r.idx] == char && r.data[r.idx-1] != '\\' {
				break
			}
		}
	}
}

func (r *StringParser) GetLeft() string {
	return string(r.data[r.idx:len(r.data)])
}

func isInRune(a rune, b []rune) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
