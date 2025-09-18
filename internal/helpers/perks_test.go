package helpers

import (
	"encoding/json"
	"testing"

	"github.com/victorprocure/opendominiongo/internal/dto"
)

type kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestMarshalPerksAsJSONArrayFromMap_DeterministicOrder(t *testing.T) {
	m := map[string]string{
		"zeta":  "3",
		"alpha": "1",
		"beta":  "2",
	}

	var prev string
	for i := range 10 {
		b, err := MarshalPerksAsJSONArrayFromMap(m)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		if b == nil {
			t.Fatalf("unexpected nil bytes for non-empty map")
		}
		var arr []kv
		if err := json.Unmarshal(b, &arr); err != nil {
			t.Fatalf("unmarshal output: %v", err)
		}
		if len(arr) != 3 {
			t.Fatalf("expected 3 items, got %d", len(arr))
		}
		expected := []string{"alpha", "beta", "zeta"}
		for i, k := range expected {
			if arr[i].Key != k {
				t.Fatalf("expected key[%d]=%s, got %s", i, k, arr[i].Key)
			}
		}
		cur := string(b)
		if i > 0 && cur != prev {
			t.Fatalf("output not deterministic across runs: prev=%s cur=%s", prev, cur)
		}
		prev = cur
	}
}

func TestMarshalPerksAsJSONArrayFromMap_Empty(t *testing.T) {
	b, err := MarshalPerksAsJSONArrayFromMap(nil)
	if err != nil {
		t.Fatalf("unexpected error for nil map: %v", err)
	}
	if b != nil {
		t.Fatalf("expected nil bytes for nil map")
	}
	b, err = MarshalPerksAsJSONArrayFromMap(map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error for empty map: %v", err)
	}
	if b != nil {
		t.Fatalf("expected nil bytes for empty map")
	}
}

func TestPerksToMap_Basic(t *testing.T) {
	kv := dto.KeyValues{
		{Key: "a", Value: "1"},
		{Key: "b", Value: "2"},
	}
	m := PerksToMap(kv)
	if len(m) != 2 || m["a"] != "1" || m["b"] != "2" {
		t.Fatalf("unexpected map: %#v", m)
	}
}

func TestPerksToMap_Empty(t *testing.T) {
	var kvNil dto.KeyValues
	if m := PerksToMap(kvNil); m != nil {
		t.Fatalf("expected nil for nil KeyValues, got %#v", m)
	}
	kvEmpty := dto.KeyValues{}
	if m := PerksToMap(kvEmpty); m != nil {
		t.Fatalf("expected nil for empty KeyValues, got %#v", m)
	}
}

func TestMarshalPerksAsJSONArrayFromKeyValues_Deterministic(t *testing.T) {
	kvs := dto.KeyValues{
		{Key: "zeta", Value: "3"},
		{Key: "alpha", Value: "1"},
		{Key: "beta", Value: "2"},
	}
	var prev string
	for i := range 10 {
		b, err := MarshalPerksAsJSONArrayFromKeyValues(kvs)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		var arr []kv
		if err := json.Unmarshal(b, &arr); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		expected := []string{"alpha", "beta", "zeta"}
		for i, k := range expected {
			if arr[i].Key != k {
				t.Fatalf("expected key[%d]=%s, got %s", i, k, arr[i].Key)
			}
		}
		cur := string(b)
		if i > 0 && cur != prev {
			t.Fatalf("nondeterministic output: prev=%s cur=%s", prev, cur)
		}
		prev = cur
	}
}
