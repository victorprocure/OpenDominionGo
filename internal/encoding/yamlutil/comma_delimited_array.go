package yamlutil

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type CommaDelimitedArray struct {
	v []string
}

func NewCommaDelimitedArray(v []string) CommaDelimitedArray {
	return CommaDelimitedArray{v: v}
}

func (c *CommaDelimitedArray) MarshalYAML() (any, error) {
	if len(c.v) == 0 {
		return []string{}, nil
	}

	return c.v, nil
}

func (c *CommaDelimitedArray) UnmarshalYAML(n *yaml.Node) error {
	if n == nil || n.Kind == 0 {
		c.v = []string{}
		return nil
	}

	switch n.Kind {
	case yaml.ScalarNode:
		var s string
		if err := n.Decode(&s); err != nil {
			return err
		}

		s = strings.TrimSpace(s)
		if s == "" {
			c.v = []string{}
			return nil
		}

		parts := strings.FieldsFunc(s, func(r rune) bool {
			return r == ',' || r == ' ' || r == '\t' || r == '\n' || r == '\r'
		})

		c.v = parts
		return nil

	case yaml.SequenceNode:
		out := make([]string, 0, len(n.Content))
		for _, child := range n.Content {
			var s string
			if err := child.Decode(&s); err != nil {
				return fmt.Errorf("requires: expected string items: %w", err)
			}
			s = strings.TrimSpace(s)
			if s != "" {
				out = append(out, s)
			}
		}
		c.v = out
		return nil

	default:
		return fmt.Errorf("requires: unsupported YAML node kind %v", n.Kind)
	}

}

func (c *CommaDelimitedArray) ToString() string {
	return strings.Join(c.v, ", ")
}
