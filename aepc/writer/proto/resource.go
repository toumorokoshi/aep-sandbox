package proto

import (
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/toumorokoshi/aep-sandbox/aepc/schema"
)

// AddResource adds a resource's protos and RPCs to a file and service.
func AddResource(r schema.Resource, fb *builder.FileBuilder, sb *builder.ServiceBuilder) error {
	err := AddCreate(r, fb, sb)
	if err != nil {
		return err
	}
	return nil
}

func AddCreate(r schema.Resource, fb *builder.FileBuilder, sb *builder.ServiceBuilder) error {
	// create request messages
	mb := builder.NewMessage("Create" + r.Kind + "Request")
	mb.AddField(
		builder.NewField("id", builder.FieldTypeString()).SetNumber(1),
	)
	fb.AddMessage(mb)
	return nil
}
