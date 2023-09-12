package schema

// Service is an in-memory representation
// for a resource used by aepc.
type Service struct {
	Name      string
	Resources []Resource
}

// Resource is the in-memory for a resource
// used by aepc.
type Resource struct {
	// Kind is the type of the resource
	Kind       string
	Properties map[string]Property
}

type Property struct {
	// Number represents the field number, identical to the usage
	// of the number in protobuf.
	Number int
}

type FieldType int64

const (
	String FieldType = iota
	Number
	Object
)
