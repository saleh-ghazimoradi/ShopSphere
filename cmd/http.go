package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/ShopSphere/config"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/gateway"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Launching ShopSphere App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server is running on %s", config.AppConfig.ServerConfig.Port)
		if err := gateway.Server(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
