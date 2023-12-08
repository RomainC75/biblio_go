package cmd

import (
	"gitub.com/RomainC75/biblio/api/bootstrap"
	"gitub.com/RomainC75/biblio/config"

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
	config.Set()
	bootstrap.Seed()

}
