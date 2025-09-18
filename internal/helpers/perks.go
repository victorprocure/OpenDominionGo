package helpers

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/victorprocure/opendominiongo/internal/dto"
)

// PerksToMap converts dto.KeyValues to a map[string]string.
// Returns nil for empty input.
func PerksToMap(kv dto.KeyValues) map[string]string {
	if len(kv) == 0 {
		return nil
	}
	m := make(map[string]string, len(kv))
	for _, p := range kv {
		m[p.Key] = p.Value
	}
	return m
}

// MarshalPerksAsJSONArrayFromMap marshals a map of perks into a JSON
// array of {key,value} objects suitable for jsonb_to_recordset in SQL.
// Returns nil, nil for empty input.
func MarshalPerksAsJSONArrayFromMap(m map[string]string) ([]byte, error) {
	if len(m) == 0 {
		return nil, nil
	}
	type kv struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	// deterministic order: sort keys before building the array
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	arr := make([]kv, 0, len(keys))
	for _, k := range keys {
		arr = append(arr, kv{Key: k, Value: m[k]})
	}
	b, err := json.Marshal(arr)
	if err != nil {
		return nil, fmt.Errorf("marshal perks: %w", err)
	}
	return b, nil
}

// MarshalPerksAsJSONArrayFromKeyValues converts KeyValues to map then marshals.
func MarshalPerksAsJSONArrayFromKeyValues(kv dto.KeyValues) ([]byte, error) {
	return MarshalPerksAsJSONArrayFromMap(PerksToMap(kv))
}
