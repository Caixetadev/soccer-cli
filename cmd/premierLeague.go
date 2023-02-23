/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Caixetadev/soccer-cli/config"
	"github.com/Caixetadev/soccer-cli/utils"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

// premierLeagueCmd represents the premierLeague command
var premierLeagueCmd = &cobra.Command{
	Use:   "premierLeague",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		premierLeague()
	},
}

func init() {
	rootCmd.AddCommand(premierLeagueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// premierLeagueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// premierLeagueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func premierLeague() {
	c := config.Colly()

	blue := color.New(color.FgBlue).SprintFunc()
	orange := color.New(color.FgMagenta).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	fmt.Println()

	c.OnHTML(".col-main table tbody tr", func(h *colly.HTMLElement) {
		name := h.ChildText(".team-name")
		points := h.ChildText(".points")

		index := h.Index + 1

		switch {
		case index <= 4:
			fmt.Printf("%02d | %-30s | %v\n", index, blue(name), points)
		case index == 5:
			fmt.Printf("%02d | %-30s | %v\n", index, orange(name), points)
		case index >= 6 && index <= 17:
			fmt.Printf("%02d | %-30s | %v\n", index, white(name), points)
		case index >= 18 && index <= 20:
			fmt.Printf("%02d | %-30s | %v\n", index, red(name), points)
		}

	})

	if err := c.Visit("https://www.terra.com.br/esportes/futebol/internacional/inglaterra/campeonato-ingles/tabela/"); err != nil {
		utils.UnCache("https://www.terra.com.br/esportes/futebol/brasileiro-serie-a/tabela/")
		log.Fatal(err)
	}

	fmt.Printf("\n%v - %v - %v\n\n", blue("Liga dos Campeões"), orange("Liga Europa"), red("Rebaixados"))
}
