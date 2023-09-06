package reader

import (
	"fmt"
	"log"

	"github.com/jhump/protoreflect/desc/protoparse"
)

func ReadResourceFromProto(resourceFiles []string) {
	// Create a new proto parser.
	parser := protoparse.Parser{}

	// Parse the proto file.
	files, err := parser.ParseFiles(resourceFiles...)
	if err != nil {
		log.Fatal(err)
	}

	for _, fd := range files {
		// find all messages
		resources := fd.GetMessageTypes()
		for _, r := range resources {
			fmt.Printf("%s", r.GetName())
		}
	}
}
