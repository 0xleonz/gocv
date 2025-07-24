package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/0xleonz/gocv/internal/config"
)

// cnfig cargada
var AppConfig *config.LoadedConfig

var rootCmd = &cobra.Command{
	Use:   "gocv",
	Short: "Generador de CVs con Typst",
	Long: `gocv es una herramienta CLI para compilar y gestionar múltiples currículums escritos en Typst.
Carga configuración desde ~/.config/gocv/config.yml y compila sólo cuando es necesario.`,
}

func Execute() {
	var err error
	AppConfig, err = config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error cargando configuración: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// globalFlags here
func init() {
	cobra.OnInitialize()
}

