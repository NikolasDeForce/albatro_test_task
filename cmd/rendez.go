/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"parser/rendezvous/parse"

	"github.com/spf13/cobra"
)

// rendezCmd represents the rendez command
var rendezCmd = &cobra.Command{
	Use:   "rendez",
	Short: "Парсер Rendez-Vous",
	Long: `Запускает парсер сайта Rendez-Vous с параметрами:
	
	--sort - вид сортировки. Алфавитная 'asc'. Обратная 'desc'. Если нужно без сортировки - 'without'.
	--file - вид файла. 'csv' либо 'xls'.
	--name - название файла.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, _ := cmd.Flags().GetString("file")

		sortFlag, _ := cmd.Flags().GetString("sort")

		nameFlag, _ := cmd.Flags().GetString("name")

		if nameFlag == "" && fileFlag == "csv" {
			nameFlag = "rendez.csv"
		} else if nameFlag == "" && fileFlag == "xls" {
			nameFlag = "rendez.xls"
		}

		if sortFlag == "asc" && fileFlag == "csv" {
			parse.ParseToCSV(nameFlag, "asc")
		} else if sortFlag == "desc" && fileFlag == "csv" {
			parse.ParseToCSV(nameFlag, "desc")
		} else if sortFlag == "without" && fileFlag == "csv" {
			parse.ParseToCSV(nameFlag, "")
		} else if sortFlag == "asc" && fileFlag == "xls" {
			parse.ParseToXLS(nameFlag, "asc")
		} else if sortFlag == "desc" && fileFlag == "xls" {
			parse.ParseToXLS(nameFlag, "desc")
		} else if sortFlag == "without" && fileFlag == "xls" {
			parse.ParseToXLS(nameFlag, "")
		}
	},
}

func init() {
	rootCmd.AddCommand(rendezCmd)
	rendezCmd.Flags().String("sort", "s", "Применяет сортировку к конечному файлу. Если она не нужна - 'without'. Сортировка по алфавиту - 'asc'. В обратном порядке - 'desc'")
	rendezCmd.Flags().String("file", "f", "В каком файле будет сохранен файл. 'csv' либо 'xls'.")
	rendezCmd.Flags().String("name", "n", "Название файла")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rendezCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rendezCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
