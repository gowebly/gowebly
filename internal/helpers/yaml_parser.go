package helpers

import (
	"errors"
	"fmt"
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// ParseYAMLToStruct parses the given YAML file from a path to struct *T using
// "knadh/koanf" package.
func ParseYAMLToStruct[T any](path string, model *T) (*T, error) {
	// Check, if path is not empty.
	if path == "" {
		return nil, errors.New("the given path of the YAML config file is empty")
	}

	// Create a new koanf instance and parse the given path.
	k, err := newKoanfByPath(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal structured data to the given struct.
	if err = k.Unmarshal("", &model); err != nil {
		return nil, fmt.Errorf("an error unmarshalling data from the YAML config file to struct, %w", err)
	}

	return model, nil
}

// newKoanfByPath helps to parse the given path for ParseYAMLToStruct function.
func newKoanfByPath(path string) (*koanf.Koanf, error) {
	// Create a new koanf instance.
	k := koanf.New(".")

	// Get the structured file from system path.
	fileInfo, err := os.Stat(path)

	// Check, if file exists.
	if err == nil || !os.IsNotExist(err) {
		// Check, if file is not dir.
		if fileInfo.IsDir() {
			return nil, fmt.Errorf("the given path of the YAML config file (%s) is dir", path)
		}

		// Load structured file from path (with parser of the file format).
		if err = k.Load(file.Provider(path), yaml.Parser()); err != nil {
			return nil, fmt.Errorf(
				"not valid structure of the YAML config file from the given path (%s)",
				path,
			)
		}
	} else {
		return nil, fmt.Errorf("YAML config file is not found in the given path (%s)", path)
	}

	return k, nil
}
