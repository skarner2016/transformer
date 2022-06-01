package kernel

import (
	"fmt"
	"github.com/spf13/cobra"
	"transformer/library/config"
)

var MainCmd = &cobra.Command{
	Use: "go-origin",
	Short: "",
	Long: "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// construct
		config.InitConfig()

		fmt.Println("PersistentPreRun")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// destroy

		fmt.Println("PersistentPostRun")
	},
}

func init()  {
	MainCmd.AddCommand(versionCmd)
	MainCmd.AddCommand(liveCmd)
}