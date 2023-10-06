package helpers

import (
	"errors"
	"path/filepath"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"

	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/constants"
)

// ParseYAMLToStruct parses the given YAML file from a path to struct *T using
// "knadh/koanf" package.
func ParseYAMLToStruct[T any](path string, model *T) (*T, error) {
	// Check, if path is not empty.
	if path == "" {
		return nil, errors.New(constants.ErrorGoweblyConfigPathIsEmpty)
	}

	// Create a new koanf instance and parse the given path.
	k, err := newKoanfByPath(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	// Unmarshal structured data to the given struct.
	if err = k.Unmarshal("", &model); err != nil {
		return nil, errors.New(constants.ErrorGoweblyUnmarshalConfigFileToStructNotComplete)
	}

	return model, nil
}

// newKoanfByPath helps to parse the given path to ParseYAMLToStruct function.
func newKoanfByPath(path string) (*koanf.Koanf, error) {
	// Create a new koanf provider variable.
	var provider koanf.Provider

	// Create a new koanf instance.
	k := koanf.New(".")

	// Check, if config file ('.cgapp.yml') is existing.
	if IsExistInFolder(path, false) {
		// If exists, set provider to the file.Provider with the given path.
		provider = file.Provider(path)
	} else {
		// Else, read the default config file from the embed.FS.
		defaultConfig, err := attachments.ConfigsFiles.ReadFile("configs/default.yml")
		if err != nil {
			return nil, errors.New(constants.ErrorGoweblyDefaultYAMLConfigFileNotFound)
		}

		// Set provider to the rawbytes.Provider with a default config file.
		provider = rawbytes.Provider(defaultConfig)
	}

	// Load structured file from the given provider with the koanf's yaml.Parser.
	if err := k.Load(provider, yaml.Parser()); err != nil {
		return nil, errors.New(constants.ErrorGoweblyYAMLConfigFileNotValid)
	}

	return k, nil
}
