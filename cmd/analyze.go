/*
Copyright © 2025 Krishna Biswal <kshnabiswal619@gmail.com>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/KrishT0/log-analyzer/internal"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze [file]",
	Short: "Parse and analyze a file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. 
For example: log-analyzer analyze <logfile>`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		ext := strings.ToLower(filepath.Ext(filePath))

		switch ext {
		case ".json":
			internal.ParseJSON(filePath)
		// case ".log":
		// 	internal.ParseLog(filePath)
		// case ".csv":
		// 	parseCSV(filePath)
		// case ".txt":
		// 	parseTXT(filePath)
		// case ".xml":
		// 	parseXML(filePath)
		default:
			fmt.Println("Unsupported file format:", ext)
		}

	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
