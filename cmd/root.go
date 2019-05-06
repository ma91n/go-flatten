package cmd

import (
	"github.com/spf13/cobra"
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)

func init() {
	rootCmd.Flags().StringVarP(&format, "format", "f", "JSON", "output format")
}

var rootCmd = &cobra.Command{
	Use: "flatten",
	RunE: func(cmd *cobra.Command, args []string) error {

		line, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		input := map[string]interface{}{}
		if err := json.Unmarshal(line, &input); err != nil {
			fmt.Fprintln(os.Stderr, string(line))
			//file, err := os.Create(`test.log`)
			//if err != nil {
			//	fmt.Fprintln(os.Stderr, string(line))
			//}
			//defer file.Close()
			//
			//file.Write(line)

			return err
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
			return err
		}

		_, err = cmd.OutOrStdout().Write(output)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		return nil
	},
}

var format string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
