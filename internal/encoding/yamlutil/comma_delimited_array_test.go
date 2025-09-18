package yamlutil

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestCommaDelimitedArray_UnmarshalYAML_Scalar(t *testing.T) {
	var n yaml.Node
	n.Kind = yaml.ScalarNode
	n.Value = "a, b, c"
	var c CommaDelimitedArray
	if err := c.UnmarshalYAML(&n); err != nil {
		t.Fatalf("unmarshal scalar: %v", err)
	}
	if got := c.ToString(); got != "a, b, c" {
		t.Fatalf("expected 'a, b, c', got %s", got)
	}
}

func TestCommaDelimitedArray_UnmarshalYAML_Sequence(t *testing.T) {
	var n yaml.Node
	n.Kind = yaml.SequenceNode
	// Build sequence nodes with content
	child1 := yaml.Node{Kind: yaml.ScalarNode, Value: "one"}
	child2 := yaml.Node{Kind: yaml.ScalarNode, Value: "two"}
	n.Content = []*yaml.Node{&child1, &child2}
	var c CommaDelimitedArray
	if err := c.UnmarshalYAML(&n); err != nil {
		t.Fatalf("unmarshal sequence: %v", err)
	}
	if got := c.ToString(); got != "one, two" {
		t.Fatalf("expected 'one, two', got %s", got)
	}
}
