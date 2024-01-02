package cmd

import (
	"fmt"
	"github.com/kitabisa/kibitalk/cmd/api"
	"github.com/kitabisa/kibitalk/cmd/migration"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(api.RestApiCmd, migration.MigrateUpCmd)
}
