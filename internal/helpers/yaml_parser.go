package helpers

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"

	"github.com/gowebly/gowebly/internal/attachments"
)

// ParseYAMLToStruct parses the given YAML file from a path to struct *T using
// "knadh/koanf" package.
func ParseYAMLToStruct[T any](path string, model *T) (*T, error) {
	// Check, if path is not empty.
	if path == "" {
		return nil, errors.New("the given path of the YAML config file is empty")
	}

	// Create a new koanf instance and parse the given path.
	k, err := newKoanfByPath(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	// Unmarshal structured data to the given struct.
	if err = k.Unmarshal("", &model); err != nil {
		return nil, fmt.Errorf(
			"an error with unmarshalling data process from the YAML config file to struct, %w",
			err,
		)
	}

	return model, nil
}

// newKoanfByPath helps to parse the given path for ParseYAMLToStruct function.
func newKoanfByPath(path string) (*koanf.Koanf, error) {
	// Create a new koanf provider variable.
	var provider koanf.Provider

	// Create a new koanf instance.
	k := koanf.New(".")

	// Check, if config file ('.cgapp.yml') is existing.
	if IsFileExist(path) {
		// If exists, set provider to the file.Provider with the given path.
		provider = file.Provider(path)
	} else {
		// Else, read the default config file from the embed.FS.
		defaultConfig, err := attachments.ConfigsFiles.ReadFile("configs/default.yml")
		if err != nil {
			return nil, errors.New("a default YAML config file is not found")
		}
		// Set provider to the rawbytes.Provider with a default config file.
		provider = rawbytes.Provider(defaultConfig)
	}

	// Load structured file from the given provider with the koanf's yaml.Parser.
	if err := k.Load(provider, yaml.Parser()); err != nil {
		return nil, errors.New("not valid structure of the YAML config file from the given path")
	}

	return k, nil
}
