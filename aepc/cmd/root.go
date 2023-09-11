package cmd

import (
	"fmt"

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
			fmt.Printf("output file: %s\n", outputFile)
			fmt.Printf("output proto: %s\n", proto)
		},
	}
	c.Flags().StringSliceVarP(&inputFiles, "input", "i", []string{}, "input files with resource")
	c.Flags().StringVarP(&outputFile, "output", "o", "", "output file to use")
	return c
}
