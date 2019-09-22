package main

import (
	"gin-gen-markdown/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "main"}

	rootCmd.AddCommand(cmd.Cmd)
	rootCmd.Execute()
}
