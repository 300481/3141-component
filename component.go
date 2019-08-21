package component

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Component struct {
	Name       string
	Enabled    bool
	Repository string
	Path       string
	Reference  string
	Values     string
}

type ComponentList struct {
	Items []Component
}

func NewFromPath(path string) (compList *ComponentList, err error) {
	glob := path + "/*yaml"
	files, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}

	var componentList ComponentList

	for _, filename := range files {
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var component Component
		err = yaml.Unmarshal(yamlFile, &component)
		if err != nil {
			return nil, err
		}

		componentList.Items = append(componentList.Items, component)
	}

	return &componentList, nil
}
