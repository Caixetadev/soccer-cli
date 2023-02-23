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

// brasileiraoCmd represents the brasileirao command
var brasileiraoCmd = &cobra.Command{
	Use:   "brasileirao",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		brasileirao()
	},
}

func init() {
	rootCmd.AddCommand(brasileiraoCmd)
}

func brasileirao() {
	c := config.Colly()

	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	orange := color.New(color.FgMagenta).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	fmt.Println()

	c.OnHTML(".col-main table tbody tr", func(h *colly.HTMLElement) {
		name := h.ChildAttr("a", "title")
		points := h.ChildText(".points")

		index := h.Index + 1

		switch {
		case index <= 4:
			fmt.Printf("%02d | %-25s | %v\n", index, blue(name), points)
		case index == 5 || index == 6:
			fmt.Printf("%02d | %-25s | %v\n", index, orange(name), points)
		case index >= 7 && index <= 12:
			fmt.Printf("%02d | %-25s | %v\n", index, green(name), points)
		case index >= 13 && index <= 16:
			fmt.Printf("%02d | %-25s | %v\n", index, white(name), points)
		case index >= 17 && index <= 20:
			fmt.Printf("%02d | %-25s | %v\n", index, red(name), points)
		}

	})

	if err := c.Visit("https://www.terra.com.br/esportes/futebol/brasileiro-serie-a/tabela/"); err != nil {
		utils.UnCache("https://www.terra.com.br/esportes/futebol/brasileiro-serie-a/tabela/")
		log.Fatal(err)
	}

	fmt.Printf("\n%v - %v - %v - %v\n\n", blue("Libertadores"), orange("Pré-libertadores"), green("Copa Sul-Americana"), red("Rebaixados"))
}
