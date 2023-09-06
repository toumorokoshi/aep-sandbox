package writer

import (
	"bytes"
	// "github.com/jhump/protoreflect/desc"
	// "github.com/jhump/protoreflect/desc/protoparse"
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
