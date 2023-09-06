package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	reader "github.com/toumorokoshi/aep-sandbox/aepc/readers"
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
			reader.ReadResourceFromProto(inputFiles)
			fmt.Printf("output: %s", outputFile)
		},
	}
	c.Flags().StringSliceVarP(&inputFiles, "input", "i", []string{}, "input files with resource")
	c.Flags().StringVarP(&outputFile, "output", "o", "", "output file to use")
	return c
}
