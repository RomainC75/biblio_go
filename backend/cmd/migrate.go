package cmd

import (
	"gitub.com/RomainC75/biblio/api/bootstrap"
	"gitub.com/RomainC75/biblio/config"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "table migration",
	Long:  `Application will be served on host and port defined`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	config.Set()
	bootstrap.Migrate()
}
