/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Caixetadev/soccer-cli/config"
	"github.com/Caixetadev/soccer-cli/utils"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

type Seila struct {
	TeamHome string
	TeamAway string
}

// championsLeagueCmd represents the championsLeague command
var championsLeagueCmd = &cobra.Command{
	Use:   "championsLeague",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		championsLeague()
	},
}

func init() {
	rootCmd.AddCommand(championsLeagueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// championsLeagueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// championsLeagueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Match struct {
	HomeTeam string
	AwayTeam string
	Result   string
	Date     string
}

func extractMatch(h *colly.HTMLElement) Match {
	home := h.ChildText(".home .fullname")
	away := h.ChildText(".away .fullname")
	result := h.ChildText(".goals.home") + " x " + h.ChildText(".goals.away")
	date := h.ChildText(".time")
	return Match{HomeTeam: home, AwayTeam: away, Result: result, Date: date}
}

func championsLeague() {
	c := config.Colly()

	var matchPlayed []Match
	var unplayedMatch []Match

	c.OnHTML("#mod-606-standings-playoff-list .group li ul .match", func(h *colly.HTMLElement) {
		if h.ChildText(".goals") != "" {
			matchPlayed = append(matchPlayed, extractMatch(h))
		} else {
			unplayedMatch = append(unplayedMatch, extractMatch(h))
		}
	})

	if err := c.Visit("https://www.terra.com.br/esportes/futebol/internacional/liga-dos-campeoes/tabela/"); err != nil {
		utils.UnCache("https://www.terra.com.br/esportes/futebol/internacional/liga-dos-campeoes/tabela/")
		log.Fatal(err)
	}

	fmt.Printf("\nJogos de Ida\n\n")

	for _, match := range matchPlayed {
		fmt.Printf("%-20s  %-8s  %-20s\n", match.HomeTeam, match.Result, match.AwayTeam)
	}

	fmt.Printf("\nJogos de Volta\n\n")

	for _, team := range unplayedMatch {
		fmt.Printf("%-20s  %-8s  %-20s |   %v\n", team.HomeTeam, team.Result, team.AwayTeam, team.Date)
	}

	fmt.Println()
}
