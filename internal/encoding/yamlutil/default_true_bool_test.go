package yamlutil

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDefaultTrueBool_UnmarshalJSON_AbsentNullAndValues(t *testing.T) {
	var b DefaultTrueBool

	// null
	if err := json.Unmarshal([]byte("null"), &b); err != nil {
		t.Fatalf("unmarshal null json: %v", err)
	}
	if b.Valid {
		t.Fatalf("expected Valid=false for null json")
	}
	if got := b.OrDefault(); got != true {
		t.Fatalf("expected OrDefault true for null, got %v", got)
	}

	// explicit false
	var bf DefaultTrueBool
	if err := json.Unmarshal([]byte("false"), &bf); err != nil {
		t.Fatalf("unmarshal false json: %v", err)
	}
	if !bf.Valid || bf.Value != false {
		t.Fatalf("expected Valid=true Value=false, got Valid=%v Value=%v", bf.Valid, bf.Value)
	}

	// explicit true
	var bt DefaultTrueBool
	if err := json.Unmarshal([]byte("true"), &bt); err != nil {
		t.Fatalf("unmarshal true json: %v", err)
	}
	if !bt.Valid || bt.Value != true {
		t.Fatalf("expected Valid=true Value=true, got Valid=%v Value=%v", bt.Valid, bt.Value)
	}
}

func TestDefaultTrueBool_UnmarshalYAML_AbsentAndValues(t *testing.T) {
	var b DefaultTrueBool
	// absent -> UnmarshalYAML not called; simulate nil node
	if err := b.UnmarshalYAML(nil); err != nil {
		t.Fatalf("unmarshal nil yaml node: %v", err)
	}
	if b.Valid {
		t.Fatalf("expected Valid=false for nil yaml node")
	}
	if b.OrDefault() != true {
		t.Fatalf("expected OrDefault true for nil yaml node")
	}

	// explicit false
	var n yaml.Node
	n.Kind = yaml.ScalarNode
	n.Value = "false"
	var bf DefaultTrueBool
	if err := bf.UnmarshalYAML(&n); err != nil {
		t.Fatalf("unmarshal false yaml: %v", err)
	}
	if !bf.Valid || bf.Value != false {
		t.Fatalf("expected Valid=true Value=false, got Valid=%v Value=%v", bf.Valid, bf.Value)
	}
}

func TestDefaultTrueBool_MarshalJSON(t *testing.T) {
	b := NewDefaultTrueBool(true)
	data, err := json.Marshal(b)
	if err != nil {
		t.Fatalf("marshal json: %v", err)
	}
	if string(data) != "true" {
		t.Fatalf("expected 'true', got %s", string(data))
	}

	var empty DefaultTrueBool
	data2, err := json.Marshal(empty)
	if err != nil {
		t.Fatalf("marshal empty json: %v", err)
	}
	if string(data2) != "true" {
		t.Fatalf("expected 'true' for empty, got %s", string(data2))
	}
}
