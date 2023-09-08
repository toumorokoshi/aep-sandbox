package writer

import (
	"bytes"
	// "github.com/jhump/protoreflect/desc"
	// "github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/toumorokoshi/aep-sandbox/aepc/schema"

	// "golang.org/x/exp/maps"
	"text/template"
)

const serviceTemplate = `
syntax = "proto3";
package {{.Name}};

service {{.Name}} {
{{range .Resources}}
  message {{.Kind}} {
  }
{{end}}
}
`

func WriteServiceToProto(s schema.Service) (string, error) {
	// we write the
	// files := map[string]string{}
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return "", err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, s)

	if err != nil {
		return "", err
	}
	return output.String(), nil
	// p := protoparse.Parser{
	// 	Accessor:              protoparse.FileContentsFromMap(files),
	// 	IncludeSourceCodeInfo: true,
	// 	LookupImport:          desc.LoadFileDescriptor,
	// }
	// filenames := maps.Keys(files)
	// fds, err := p.ParseFiles(filenames...)
	// protodesc.NewFileDescriptor()
	// f := protodesc.NewFile()
	// m := protodesc.NewMessageDescriptor()
}

func WriteServiceToProto2(s schema.Service) (string, error) {
	fileBuilder := builder.NewFile("test.proto")
	serviceBuilder := builder.NewService(s.Name)
	fileBuilder.AddService(serviceBuilder)
	fd, err := fileBuilder.Build()
	if err != nil {
		return "", err
	}
	printer := protoprint.Printer{}
	var output bytes.Buffer
	err = printer.PrintProtoFile(fd, &output)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
