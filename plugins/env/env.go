// Package env provides environment variables support for uconfig
package env

import (
	"os"
	"strings"

	"github.com/oskoi/uconfig/flat"
	"github.com/oskoi/uconfig/plugins"
)

const tag = "env"

func init() {
	plugins.RegisterTag(tag)
}

// New returns an EnvSet.
func New(prefix string) plugins.Plugin {
	return &visitor{
		prefix: prefix,
	}
}

type visitor struct {
	prefix string
	fields flat.Fields
}

func makeEnvName(prefix, name string) string {
	name = strings.ReplaceAll(name, ".", "_")
	name = strings.ToUpper(name)
	if prefix != "" {
		name = prefix + "_" + name
	}
	return name
}

func (v *visitor) Visit(f flat.Fields) error {
	v.fields = f

	for _, f := range v.fields {
		name, explicit := f.Name(tag)
		if !explicit {
			name = makeEnvName(v.prefix, name)
		}

		f.Meta()[tag] = name
	}

	return nil
}

func (v *visitor) Parse() error {
	for _, f := range v.fields {

		name := f.Meta()[tag]

		if name == "-" {
			continue
		}

		value, ok := os.LookupEnv(name)
		if !ok {
			continue
		}

		err := f.Set(value)
		if err != nil {
			return err
		}
	}

	return nil
}
