package tag_parser

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type parser struct {
	TagParser
	Tag
}

type KeyVal struct {
	Key string
	Val string
}

type Param struct {
	Name string
	Args []KeyVal
}

type Tag struct {
	Extractor *Param
	// ;
	Transformer *Param
	// ;
	Required bool
	Default  string
}

func Parse(content string) (*Tag, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, fmt.Errorf("tag conf can not be empty")
	}
	p := &parser{}
	p.Buffer = content

	if err := p.Init(); err != nil {
		return nil, err
	}
	if err := p.TagParser.Parse(); err != nil {
		return nil, err
	}
	if false {
		p.PrintSyntaxTree()
		os.Exit(0)
	}
	if false {
		p.PrintSyntaxTree()
	}
	if err := p.selfParse(); err != nil {
		return nil, err
	}
	if p.Extractor == nil || p.Extractor.Name == "" {
		return nil, errors.New("extractor can not be empty")
	}
	return &p.Tag, nil
}

// copy from https://github.com/cloudwego/thriftgo/blob/main/parser/parser.go
func (p *parser) selfParse() (err error) {
	root := p.AST()
	if root == nil || root.pegRule != ruleDocument {
		return errors.New("not document")
	}
	for n := root.up; n != nil; n = n.next {
		switch n.pegRule {
		case ruleSkip:
			continue
		case ruleExtractor:
			if err = p.parseExtractor(n); err != nil {
				return
			}
		case ruleTransformer:
			if err = p.parseTransformer(n); err != nil {
				return
			}
		case ruleOption:
			if err = p.parseOption(n); err != nil {
				return
			}
		default:
			return fmt.Errorf("unknown rule: " + rul3s[n.pegRule])
		}
	}
	return nil
}

func (p *parser) parseExtractor(node *node32) (err error) {
	p.Extractor, err = p.parseParam(node, ruleExtractor)
	return err
}

func (p *parser) parseTransformer(node *node32) (err error) {
	p.Transformer, err = p.parseParam(node, ruleTransformer)
	return err
}

func (p *parser) parseOption(node *node32) (err error) {
	node, err = checkrule(node, ruleOption)
	if err != nil {
		return err
	}
	if node.pegRule == ruleSkip {
		node = node.next
	}

	//	Option <- (Required / Default)? (COMMA (Required / Default))*
	for node != nil {
		switch node.pegRule {
		case ruleRequired:
			p.Tag.Required = true
			node = node.next
		case ruleDefault:
			// node = node.next // DEFAULT
			// node = node.next // EQUAL
			p.Tag.Default = p.pegText(node)
			node = node.next
		default:
			node = node.next
		}
	}
	return nil
}

func (p *parser) parseParam(node *node32, rule pegRule) (param *Param, err error) {
	param = new(Param)

	node, err = checkrule(node, rule)
	if err != nil {
		return nil, err
	}
	if node.pegRule == ruleSkip {
		node = node.next
	}

	switch node.pegRule {
	case ruleIdentifier:
		param.Name = p.pegText(node)
		node = node.next
	}

	for node != nil {
		node = node.next // COMMA

		key := p.pegText(node)
		node = node.next

		node = node.next // EQUAL

		val := p.pegText(node)
		node = node.next

		param.Args = append(param.Args, KeyVal{key, val})
	}
	return param, nil
}

func (p *parser) pegText(node *node32) string {
	for n := node; n != nil; n = n.next {
		if s := p.pegText(n.up); s != "" {
			return s
		}
		if n.pegRule != rulePegText {
			continue
		}
		var runes []rune
		for i := n.begin; i < n.end; i++ {
			if p.buffer[i] == '\\' && i+1 < n.end {
				if r, ok := escapes[p.buffer[i+1]]; ok {
					runes = append(runes, r)
					i++
					continue
				}
			}
			runes = append(runes, p.buffer[i])
		}
		if text := string(runes); text != "" {
			return text
		}
	}
	return ""
}

func checkrule(node *node32, rule pegRule) (*node32, error) {
	if node.pegRule != rule {
		return nil, fmt.Errorf("mismatch rule: " + rul3s[node.pegRule])
	}
	return node.up, nil
}

var escapes = map[rune]rune{
	'\\': '\\', '"': '"', '\'': '\'',
	'a': '\a', 'b': '\b', 't': '\t', 'n': '\n',
	'v': '\v', 'f': '\f', 'r': '\r',
}
