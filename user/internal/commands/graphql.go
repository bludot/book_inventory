/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package commands

import (
	"github.com/bludot/tempmee/user/http"

	"github.com/spf13/cobra"
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return http.StartGraphQLServer()
	},
}

func init() {
	serveCmd.AddCommand(graphqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
