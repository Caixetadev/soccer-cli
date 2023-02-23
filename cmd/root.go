/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "soccer-cli",
	Short: "Soccer-CLI: Uma ferramenta CLI para obter informações e resultados de jogos de futebol.",
	Long: `Soccer-CLI é uma ferramenta CLI em Go que permite aos usuários obter informações atualizadas sobre as tabelas de todos os campeonatos de futebol em tempo real. Com o Soccer-CLI, os usuários podem acessar rapidamente as tabelas de seus campeonatos favoritos, bem como informações detalhadas sobre equipes, jogadores, partidas e resultados. A ferramenta é ideal para fãs de futebol que desejam acompanhar de perto os resultados de seus times e campeonatos preferidos através de uma interface de linha de comando fácil de usar.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.soccer-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
