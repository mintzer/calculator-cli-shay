package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "calc <operation> <a> <b>",
		Short: "Simple calculator CLI.",
		Long:  "A simple CLI calculator that supports add, sub, mul, and div operations on two numbers.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			operation := args[0]
			aStr := args[1]
			bStr := args[2]

			a, err := strconv.ParseFloat(aStr, 64)
			if err != nil {
				return fmt.Errorf("invalid number: %s", aStr)
			}

			b, err := strconv.ParseFloat(bStr, 64)
			if err != nil {
				return fmt.Errorf("invalid number: %s", bStr)
			}

			var result float64
			switch operation {
			case "add":
				result = Add(a, b)
			case "sub":
				result = Subtract(a, b)
			case "mul":
				result = Multiply(a, b)
			case "div":
				val, divErr := Divide(a, b)
				if divErr != nil {
					return fmt.Errorf("Cannot divide by zero")
				}
				result = val
			default:
				return fmt.Errorf("unknown operation: %s", operation)
			}

			fmt.Println(strconv.FormatFloat(result, 'f', -1, 64))
			return nil
		},
	}

	rootCmd.Flags().SetInterspersed(false)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
