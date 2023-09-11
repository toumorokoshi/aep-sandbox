package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/toumorokoshi/aep-sandbox/aepc/reader"
	"github.com/toumorokoshi/aep-sandbox/aepc/writer/proto"
)

func NewCommand() *cobra.Command {
	var inputFiles []string
	var outputFile string

	c := &cobra.Command{
		Use:   "aepc",
		Short: "aepc compiles resource representations to full proto rpcs",
		Long:  "aepc compiles resource representations to full proto rpcs",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Printf("input: %s\n", inputFiles)
			// TODO: error handling
			service, _ := reader.ReadServiceFromProto(inputFiles)
			proto, _ := proto.WriteServiceToProto(*service)
			f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatal(err)
			}
			bytesWritten, err := f.Write(proto)
			log.Printf("%v", bytesWritten)
			if err != nil {
				log.Fatal(err)
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("output file: %s\n", outputFile)
			fmt.Printf("output proto: %s\n", proto)
		},
	}
	c.Flags().StringSliceVarP(&inputFiles, "input", "i", []string{}, "input files with resource")
	c.Flags().StringVarP(&outputFile, "output", "o", "", "output file to use")
	return c
}
