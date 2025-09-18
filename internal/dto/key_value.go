package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type KeyValue struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

type KeyValues []KeyValue

func (pl *KeyValues) UnmarshalYAML(n *yaml.Node) error {
	switch n.Kind {
	case yaml.MappingNode:
		out := make([]KeyValue, 0, len(n.Content)/2)
		for i := 0; i < len(n.Content); i += 2 {
			k := n.Content[i].Value
			// try as string, then as number -> string
			var sv string
			if err := n.Content[i+1].Decode(&sv); err != nil {
				var f float64
				if err2 := n.Content[i+1].Decode(&f); err2 != nil {
					return fmt.Errorf("perks[%s]: decode: %w", k, err)
				}
				// trim .0 for ints
				sv = strings.TrimRight(strings.TrimRight(strconv.FormatFloat(f, 'f', -1, 64), "0"), ".")
			}
			out = append(out, KeyValue{Key: k, Value: sv})
		}
		*pl = out
		return nil
	case yaml.SequenceNode:
		var seq []KeyValue
		if err := n.Decode(&seq); err != nil {
			return err
		}
		*pl = seq
		return nil
	case 0:
		*pl = nil
		return nil
	default:
		return fmt.Errorf("perks: unsupported YAML node kind %v", n.Kind)
	}
}

// Optional: emit back as a mapping to keep YAML tidy.
func (pl KeyValues) MarshalYAML() (any, error) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	for _, kv := range pl {
		node.Content = append(node.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: kv.Key},
			&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: kv.Value},
		)
	}
	return node, nil
}

func (pl *KeyValues) UnmarshalJSON(b []byte) error {
	// Trim to detect empty/null
	bt := bytes.TrimSpace(b)
	if len(bt) == 0 || string(bt) == "null" {
		*pl = nil
		return nil
	}

	switch bt[0] {
	case '{':
		// Preserve key order using a Decoder
		dec := json.NewDecoder(bytes.NewReader(bt))
		t, err := dec.Token()
		if err != nil || t != json.Delim('{') {
			return fmt.Errorf("perks: invalid JSON object")
		}
		out := make([]KeyValue, 0, 8)
		for dec.More() {
			// Next token must be a key (string)
			tk, err := dec.Token()
			if err != nil {
				return fmt.Errorf("perks: read key: %w", err)
			}
			key, ok := tk.(string)
			if !ok {
				return fmt.Errorf("perks: object key must be string")
			}
			// Read the value into RawMessage
			var rm json.RawMessage
			if err := dec.Decode(&rm); err != nil {
				return fmt.Errorf("perks[%s]: decode value: %w", key, err)
			}
			val, err := jsonValueToString(rm)
			if err != nil {
				return fmt.Errorf("perks[%s]: %w", key, err)
			}
			out = append(out, KeyValue{Key: key, Value: val})
		}
		// Consume closing '}'
		if t, err = dec.Token(); err != nil || t != json.Delim('}') {
			return fmt.Errorf("perks: invalid JSON object end")
		}
		*pl = out
		return nil

	case '[':
		var seq []KeyValue
		if err := json.Unmarshal(bt, &seq); err != nil {
			return fmt.Errorf("perks: decode array: %w", err)
		}
		*pl = seq
		return nil

	default:
		return fmt.Errorf("perks: unsupported JSON (must be object or array)")
	}
}

// MarshalJSON emits an object mapping, preserving the slice order.
func (pl KeyValues) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, kv := range pl {
		if i > 0 {
			buf.WriteByte(',')
		}
		kb, err := json.Marshal(kv.Key)
		if err != nil {
			return nil, fmt.Errorf("perks: marshal key: %w", err)
		}
		vb, err := json.Marshal(kv.Value)
		if err != nil {
			return nil, fmt.Errorf("perks: marshal value: %w", err)
		}
		buf.Write(kb)
		buf.WriteByte(':')
		buf.Write(vb)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// jsonValueToString converts a JSON scalar (string/number/bool/null) to string.
func jsonValueToString(rm json.RawMessage) (string, error) {
	// string
	var s string
	if err := json.Unmarshal(rm, &s); err == nil {
		return s, nil
	}
	// number
	var f float64
	if err := json.Unmarshal(rm, &f); err == nil {
		// trim trailing .0
		return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(f, 'f', -1, 64), "0"), "."), nil
	}
	// bool
	var b bool
	if err := json.Unmarshal(rm, &b); err == nil {
		if b {
			return "true", nil
		}
		return "false", nil
	}
	// null
	if bytes.Equal(bytes.TrimSpace(rm), []byte("null")) {
		return "", nil
	}
	return "", fmt.Errorf("unsupported JSON value")
}
