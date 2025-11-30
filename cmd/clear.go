package cmd

import (
	"fmt"
	"my-girlfriend/internal"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear chat memory 🧹",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := internal.ClearHistory()
		if err != nil {
			return err
		}

		fmt.Println("🧹 Chat memory cleared!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
