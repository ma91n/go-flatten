package main

import (
	"github.com/spf13/cobra"
	"fmt"
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

			result := make(map[string]interface{})

			for k, v := range input {
				switch tv := v.(type) {
				case map[string]interface{}:
					for inKey, inValue := range tv {
						key := k + "." + inKey
						result[key] = inValue
					}
				default:
					result[k] = v
				}
			}

			output, err := json.Marshal(result)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		},
	}
	root.Flags().StringVarP(&format, "format", "f", "JSON", "times to echo the input")

	root.Execute()
}
