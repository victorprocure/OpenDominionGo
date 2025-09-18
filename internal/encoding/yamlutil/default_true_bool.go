package yamlutil

import (
	"encoding/json"

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
