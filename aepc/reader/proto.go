package reader

import (
	"log"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/toumorokoshi/aep-sandbox/aepc/schema"
)

func ReadServiceFromProto(serviceFiles []string) (*schema.Service, error) {
	// Create a new proto parser.
	parser := protoparse.Parser{}

	// Parse the proto file.
	files, err := parser.ParseFiles(serviceFiles...)
	if err != nil {
		log.Fatal(err)
	}

	resources := []schema.Resource{}
	service := ""

	for _, fd := range files {
		// find all services
		services := fd.GetServices()
		for _, s := range services {
			service = s.GetName()
		}
		// find all messages
		messages := fd.GetMessageTypes()
		for _, m := range messages {
			r, err := MessageToResource(m)
			if err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
	}
	return &schema.Service{
		Name:      service,
		Resources: resources,
	}, nil
}

func MessageToResource(m *desc.MessageDescriptor) (schema.Resource, error) {
	r := schema.Resource{
		Kind: m.GetName(),
	}
	return r, nil
}
