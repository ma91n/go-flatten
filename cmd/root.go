package cmd

import (
	"github.com/spf13/cobra"
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)

func init() {
}

var rootCmd = &cobra.Command{
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

var format string

func Execute() {
	rootCmd.Flags().StringVarP(&format, "format", "f", "JSON", "times to echo the input")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
