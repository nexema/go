package generator

type PluginOpts struct {
	// EnforceBuilder tells the generator to generate unexported fields for structs and
	// generate instead getter and setter, enforcing to only use NewX method for creating X instances
	EnforceBuilder bool `json:"enforceBuilder"`

	// Go's module name of the current project
	ModuleName string `json:"moduleName"`
}
