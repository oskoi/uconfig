package uconfig

import (
	"github.com/oskoi/uconfig/plugins"
	"github.com/oskoi/uconfig/plugins/defaults"
	"github.com/oskoi/uconfig/plugins/env"
	"github.com/oskoi/uconfig/plugins/file"
)

// UnmarshalOptions represents a set of file paths and the appropriate unmarshaller function.
type UnmarshalOptions = file.UnmarshalOptions

// Load creates a uconfig manager with defaults,environment variables,
// and optionally file loaders based on the provided
// Files map.
func Load[C any](files Files, userPlugins ...plugins.Plugin) Config[C] {
	ps := make([]plugins.Plugin, 0, len(files)+2+len(userPlugins))

	// first defaults
	ps = append(ps, defaults.New())
	// then files
	ps = append(ps, files.Plugins()...)
	// then any user pugins, often just _secret_.
	ps = append(ps, userPlugins...)

	// followed by envs
	ps = append(ps, env.New())

	return New[C](ps...)
}
