package main

import (
	"fmt"
	"os"

	"github.com/FearLessSaad/SNFOK/cli/commands/auth"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "snfokctl",
	Short: "SANFOK CLI",
	Long:  `SNFOK Developed By HashX Pvt Ltd. This CLI Is Used To Controlling This Security System`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to SNFOK! Use 'snfokctl status' to get status of this system.")
	},
}

func init() {
	rootCmd.AddCommand(auth.AddUserCmd)
	rootCmd.AddCommand(auth.RotateTokenCmd)
}
