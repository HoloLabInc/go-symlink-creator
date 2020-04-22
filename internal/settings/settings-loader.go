package settings

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type StringArray []string

func (a *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}

type SymLinkSettings struct {
	SymLinkSettings []SymLinkSetting `yaml:"symlinks"`
}

type SymLinkSetting struct {
	BasePath         string
	IncludeMeta      bool        `yaml:"include-meta-file"`
	CreateDestFolder bool        `yaml:"create-dest-folder"`
	Src              string      `yaml:"src"`
	Dest             StringArray `yaml:"dest"`
	Target           StringArray `yaml:"target"`
}

func LoadSettings(path string) (s SymLinkSettings) {
	p, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	s = parseSettings(b)
	for i := 0; i < len(s.SymLinkSettings); i++ {
		s.SymLinkSettings[i].BasePath = p
	}
	return s
}

func parseSettings(data []byte) (s SymLinkSettings) {
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
	return s
}
