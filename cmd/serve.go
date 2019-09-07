package cmd

import (
	"github.com/sqsinformatique/backend/srv"
	"github.com/sqsinformatique/backend/utils"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Parent command for starting public HTTP/2 API",
	Run:   serveHandler,
}

func serveHandler(cmd *cobra.Command, args []string) {
	err := srv.Start()
	if err != nil {
		utils.Fatal(err)
	}
	exitLoop()
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
