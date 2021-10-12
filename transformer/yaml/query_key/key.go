package query_key

import (
	"fmt"
	"strconv"
)

type QueryKey struct {
	Type  string // key index
	Key   string
	Index int
	Next  *QueryKey
}

// key or index

func Parse(key string) (*QueryKey, error) {
	parser := &queryKeyParser{idx: 0, data: []rune(key), root: new(QueryKey)}

	return parser.parse()
}

type queryKeyParser struct {
	idx  int
	data []rune
	root *QueryKey
	next *QueryKey
}

func (r *queryKeyParser) parse() (*QueryKey, error) {
	if r.idx == len(r.data) {
		return nil, fmt.Errorf("empty query key")
	}

	r.removeSpace()

	// .
	if r.idx == len(r.data)-1 && r.data[r.idx] == '.' {
		return &QueryKey{Type: "key", Key: "."}, nil
	}

	for r.idx < len(r.data) {
		switch v := r.data[r.idx]; v {
		case '.':
			// key
			r.idx++

			str, err := r.parseString([]rune{'.', '['})
			if err != nil {
				return nil, err
			}
			r.appendKey("key", str, 0)
		case '[':
			// index
			r.idx++

			index, err := r.parseString([]rune{']'})
			if err != nil {
				return nil, err
			}
			idx, err := strconv.ParseInt(index, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid index %s", index)
			}
			if err := r.findRune(false, ']'); err != nil {
				return nil, err
			}
			r.appendKey("index", "", int(idx))
		default:
			return nil, fmt.Errorf("invalid query key: %s", string([]rune{v}))
		}
	}
	if r.idx == len(r.data) {
		return r.root, nil
	}
	return nil, fmt.Errorf("invalid query key")
}

func (r *queryKeyParser) findRune(isKey bool, rs ...rune) error {
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
		return fmt.Errorf("expect: `%s`, but got: `%s`", string(rs), string(r.data[r.idx:min(r.idx+len(rs), len(r.data))]))
	}
	r.idx += c
	if isKey {
		r.removeSpace()
	}
	return nil
}

func (r *queryKeyParser) removeSpace() (n int) {
	for i := r.idx; i < len(r.data); i++ {
		if r.data[i] != ' ' {
			return
		}
		r.idx++
		n++
	}
	return
}

func (r *queryKeyParser) appendKey(typ string, key string, index int) {
	if r.root.Type == "" {
		r.root.Type = typ
		r.root.Key = key
		r.root.Index = index
		r.next = r.root
	} else {
		current := new(QueryKey)
		current.Type = typ
		current.Key = key
		current.Index = index
		r.next.Next = current
		r.next = current
	}
}

func (r *queryKeyParser) parseString(endRunes []rune) (string, error) {
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
		case !quoteFound && (isInRune(d, endRunes)):
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isInRune(a rune, bs []rune) bool {
	for _, v := range bs {
		if a == v {
			return true
		}
	}
	return false
}
