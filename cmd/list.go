package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "lists available relays",
		Run: func(cmd *cobra.Command, args []string) {
			for port, device := range relays {
				fmt.Printf("%s: %s\n", port, device.String())
			}

			if len(relays) == 0 {
				fmt.Printf("No relay boards found.\n")
			}

		},
	})
}
