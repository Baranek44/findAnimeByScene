package routes

import (
	"github.com/spf13/cobra"
)

var RoutesCommand = &cobra.Command{
	Use:     "anime",
	Short:   "Find anime only by one image",
	Example: "anime file image/anime-1.jpg",
}
