package cmd

import (
	"gitub.com/RomainC75/biblio/pkg/bootstrap"
	"gitub.com/RomainC75/biblio/pkg/configu"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "db seeder",
	Long:  `db seeder `,
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	configu.Set()
	bootstrap.Seed()

}
