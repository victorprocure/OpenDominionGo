package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type DefaultTrueBool struct {
	Value bool
	Valid bool
}

func NewDefaultTrueBool(v bool) DefaultTrueBool {
	return DefaultTrueBool{Value: v, Valid: true}
}

func (b DefaultTrueBool) OrDefault() bool {
	if !b.Valid {
		return true
	}
	return b.Value
}

func (b DefaultTrueBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.OrDefault())
}
func (b *DefaultTrueBool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Value = false
		b.Valid = false
		return nil
	}
	var v bool
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.Value = v
	b.Valid = true
	return nil
}

func (b DefaultTrueBool) MarshalYAML() (any, error) {
	return b.OrDefault(), nil
}

func (b *DefaultTrueBool) UnmarshalYAML(n *yaml.Node) error {
	if n == nil || n.Kind == 0 {
		b.Valid = false
		b.Value = false
		return nil
	}

	var v bool
	if err := n.Decode(&v); err != nil {
		return err
	}

	b.Value = v
	b.Valid = true
	return nil
}

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

func IsValidYamlFileName(e fs.DirEntry) (fileName string, valid bool, err error) {
	if e.IsDir() {
		return e.Name(), false, errors.New("directory found where file expected")
	}

	name := e.Name()
	ext := strings.ToLower(filepath.Ext(name))
	if ext != ".yml" && ext != ".yaml" {
		return name, false, errors.New("not a yaml file")
	}

	return name, true, nil
}
