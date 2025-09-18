package helpers

import (
	"encoding/json"
	"testing"
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
