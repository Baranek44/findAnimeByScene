package cmd

import "github.com/spf13/cobra"

var SearchingByFile = &cobra.Command{
	Use:   "file",
	Short: "Finding anime by file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SearchingAnimeByFile(args[0])
	},
}

var SearchingByLink = &cobra.Command{
	Use:   "link",
	Short: "Finding anime by link",
	Run: func(cmd *cobra.Command, args []string) {
		SearchingAnimeByLink(args[0])
	},
}

func uploadAllCommands() {
	routes.AddCommand(SearchingByFile)
	routes.AddCommand(SearchingByLink)
}
