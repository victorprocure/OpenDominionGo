package helpers

import (
	"encoding/json"
	"errors"
	"io/fs"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type BoolDefaultTrue struct {
	v *bool
}

func NewBoolDefaultTrue(v bool) BoolDefaultTrue {
    return BoolDefaultTrue{v: &v}
}

func (b BoolDefaultTrue) Bool() bool {
	if b.v == nil {
		return true
	}

	return *b.v;
}

func (b BoolDefaultTrue) OrDefault() bool {
	if b.v == nil {
		return true
	}

	return *b.v
}

func (b BoolDefaultTrue) MarshalJSON() ([]byte, error) {
    return json.Marshal(b.OrDefault())
}
func (b *BoolDefaultTrue) UnmarshalJSON(data []byte) error {
    // null => default true
    if string(data) == "null" {
        b.v = nil
        return nil
    }
    var v bool
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    b.v = &v
    return nil
}

func (b BoolDefaultTrue) MarshalYAML() (any, error) {
	return b.OrDefault(), nil
}

func (b *BoolDefaultTrue) UnmarshalYAML(n *yaml.Node) error {
	// if key is absent, this method won't be called and v stays nil (defaulting to true now)
	if n == nil || n.Kind == 0{
		b.v = nil
		return nil
	}

	var v bool
	if err := n.Decode(&v); err != nil {
		return err
	}

	b.v = &v
	return nil
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