package schema

// Resource is the in-memory for a resource
// used by aepc.
type Resource struct {
	// Kind is the type of the resource
	Kind       string
	Properties map[string]Property
}

type Property struct {
}
