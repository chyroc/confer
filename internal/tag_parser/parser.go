package tag_parser

import (
	"errors"
	"fmt"
)

type parser struct {
	TagParser
	Tag
}

type KeyVal struct {
	Key string
	Val string
}

type Tag struct {
	extractorName string
	extractorArgs []KeyVal
	// ;
	transformerName string
	transformerArgs []KeyVal
	// ;
	Required bool
	Default  string
}

func Parse(content string) (*Tag, error) {
	p := &parser{}
	p.Buffer = content

	if err := p.Init(); err != nil {
		return nil, err
	}
	if err := p.TagParser.Parse(); err != nil {
		return nil, err
	}
	if err := p.selfParse(); err != nil {
		return nil, err
	}
	return &p.Tag, nil
}

func (p *parser) selfParse() (err error) {
	root := p.AST()
	if root == nil || root.pegRule != ruleDocument {
		return errors.New("not document")
	}
	// Header* Definition* !.
	for n := root.up; n != nil; n = n.next {
		switch n.pegRule {
		case ruleSkip, ruleSkipLine:
			continue
		case ruleHeader:
			if err := p.parseHeader(n); err != nil {
				return err
			}
		case ruleDefinition:
			if err := p.parseDefinition(n); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown rule: " + rul3s[n.pegRule])
		}
	}
	return nil
}
