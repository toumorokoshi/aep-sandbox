package proto

import (
	"bytes"
	"fmt"

	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/toumorokoshi/aep-sandbox/aepc/schema"
)

func WriteServiceToProto(s schema.Service) ([]byte, error) {
	fileBuilder := builder.NewFile("test.proto")
	serviceBuilder := builder.NewService(s.Name)
	fileBuilder.AddService(serviceBuilder)
	for _, r := range s.Resources {
		err := AddResource(r, fileBuilder, serviceBuilder)
		if err != nil {
			return []byte{}, fmt.Errorf("adding resource %v failed: %w", r.Kind, err)
		}
	}
	fd, err := fileBuilder.Build()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to build service file %v: %w", fileBuilder.GetName(), err)
	}
	printer := protoprint.Printer{}
	var output bytes.Buffer
	err = printer.PrintProtoFile(fd, &output)
	if err != nil {
		return []byte{}, err
	}
	return output.Bytes(), nil
}
