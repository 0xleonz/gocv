package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/0xleonz/gocv/internal/utils"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Muestra los CVs configurados en ~/.config/gocv",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := AppConfig // ya fue cargado en root

		if len(cfg.Data.CVs) == 0 {
			fmt.Println(utils.Colorize("⚠️  No hay CVs definidos en config.yml", utils.Yellow))
			return
		}

		// Pasar CVs a slice para indexarlos
		keys := make([]string, 0, len(cfg.Data.CVs))
		for name := range cfg.Data.CVs {
			keys = append(keys, name)
		}

		// Mostrar lista numerada
		fmt.Println(utils.Colorize("📁 CVs configurados:\n", utils.Cyan))
		for i, name := range keys {
			cv := cfg.Data.CVs[name]
			fmt.Printf("%d. %s: %s\n", i+1,
				utils.Colorize(name, utils.Green),
				utils.Colorize(cv.Description, utils.Yellow),
			)
		}

		// Preguntar selección al usuario
		fmt.Print(utils.Colorize("\n👉 Selecciona un número para ver más detalles: ", utils.Purple))

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > len(keys) {
			fmt.Println(utils.Colorize("❌ Selección inválida.", utils.Red))
			return
		}

		// Mostrar descripción larga y detalles
		selectedName := keys[choice-1]
		selectedCV := cfg.Data.CVs[selectedName]

		fmt.Printf("\n%s %s\n", utils.Colorize("📄 CV seleccionado:", utils.Cyan), utils.Colorize(selectedName, utils.Green))
		fmt.Printf("%s\n", utils.Colorize(selectedCV.LongDescription, utils.Yellow))

		if t := selectedCV.LastCompileTime(); t != nil {
			fmt.Printf("%s %s\n", utils.Colorize("🕒 Última compilación:", utils.Cyan), t.Format(time.RFC822))
		} else {
			fmt.Println(utils.Colorize("🕒 Nunca compilado", utils.Red))
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
