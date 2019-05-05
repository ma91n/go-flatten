package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"bufio"
	"os"
	"encoding/json"
)

func main() {
	var format string

	root := &cobra.Command{
		Use: "flatten",
		Run: func(cmd *cobra.Command, args []string) {
			line, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			input := map[string]interface{}{}
			if err := json.Unmarshal(line, &input); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Printf("%v\n", input)

			fmt.Println(strings.Join(args, " "))
			fmt.Println("complete")
		},
	}
	root.Flags().StringVarP(&format, "format", "f", "JSON", "times to echo the input")

	root.Execute()
}
