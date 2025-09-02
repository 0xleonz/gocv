/*
Copyright © 2025 leonz <0xleonz@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/0xleonz/gocv/internal/utils"
)

var showCmd = &cobra.Command{
	Use:   "get",
	Short: "Muestra los CVs configurados en ~/.config/gocv",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := AppConfig // ya fue cargado en root

		if len(cfg.Data.CVs) == 0 {
			fmt.Println(utils.Colorize("⚠️  No hay CVs definidos en config.yml", utils.Yellow))
			return
		}

		fmt.Println(utils.Colorize("📁 CVs configurados:\n", utils.Cyan))

		for name, cv := range cfg.Data.CVs {
			fmt.Printf("%s %s\n", utils.Colorize("•", utils.Purple), utils.Colorize(name, utils.Green))
			fmt.Printf("  %s\n", utils.Colorize(cv.Description, utils.Yellow))
			if t := cv.LastCompileTime(); t != nil {
				fmt.Printf("  %s %s\n\n", utils.Colorize("🕒 Última compilación:", utils.Cyan), t.Format(time.RFC822))
			} else {
				fmt.Printf("  %s\n\n", utils.Colorize("🕒 Nunca compilado", utils.Red))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
