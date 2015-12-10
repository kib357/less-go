package main

import (
	"github.com/kib357/less-go"
	"github.com/spf13/cobra"
	"fmt"
)

func main() {
	var input, output *string
	rootCmd := &cobra.Command{
		Use:   "lessc",
		Short: "LESS CSS compiler",
		Long:  "Crossplatform LESS CSS compiler with no dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			render(*input, *output)
		},
	}
	input = rootCmd.PersistentFlags().StringP("input", "i", "styles.less", "Input less file")
	output = rootCmd.PersistentFlags().StringP("output", "o", "styles.css", "Output css file")

	rootCmd.Execute()
}

func render(input, output string) {
	fmt.Println(input, output)
	less.RenderFile(input, output)
}