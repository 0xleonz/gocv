package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/0xleonz/gocv/internal/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa la configuración y los templates en ~/.config/gocv",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		configDir := filepath.Join(home, ".config", "gocv")
		templatesDir := filepath.Join(configDir, "templates")
		configFile := filepath.Join(configDir, "config.yml")

		if err := os.MkdirAll(templatesDir, 0o755); err != nil {
			fmt.Println(utils.Colorize("❌ Error creando directorio: "+err.Error(), utils.Red))
			os.Exit(1)
		}

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			url := "https://gitlab.com/0xleonz/gocv/-/raw/main/assets/config.yml"
			if err := downloadToFile(url, configFile); err != nil {
				fmt.Println(utils.Colorize("❌ No se pudo descargar config.yml: "+err.Error(), utils.Red))
				os.Exit(1)
			}
			fmt.Println(utils.Colorize("✅ Configuración descargada: "+configFile, utils.Green))
		} else {
			fmt.Println(utils.Colorize("⚠️  Config ya existe: "+configFile, utils.Yellow))
		}

		templateFiles := []string{
			"cvBase.typ",
			"cvFarmer.typ",
		}

		for _, name := range templateFiles {
			dest := filepath.Join(templatesDir, name)
			if _, err := os.Stat(dest); os.IsNotExist(err) {
				url := "https://gitlab.com/0xleonz/gocv/-/raw/main/assets/templates/" + name
				if err := downloadToFile(url, dest); err != nil {
					fmt.Println(utils.Colorize("❌ Error descargando "+name+": "+err.Error(), utils.Red))
					continue
				}
				fmt.Println(utils.Colorize("✅ Template descargado: "+dest, utils.Green))
			} else {
				fmt.Println(utils.Colorize("⚠️  Template ya existe: "+dest, utils.Yellow))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func downloadToFile(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("respuesta HTTP %d", resp.StatusCode)
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

