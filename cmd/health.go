package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/0xleonz/gocv/internal/compile"
	"gitlab.com/0xleonz/gocv/internal/config"
	"gitlab.com/0xleonz/gocv/internal/utils"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Verifica que el entorno esté listo para compilar los CVs",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := AppConfig
		ok := true
		needsCompile := []string{}

		fmt.Println(utils.Colorize("🩺 Verificando entorno...\n", utils.Cyan))

		if cfg != nil {
			fmt.Println(utils.Colorize("✅ Configuración cargada correctamente", utils.Green))
		} else {
			fmt.Println(utils.Colorize("❌ No se pudo cargar la configuración", utils.Red))
			os.Exit(1)
		}

		// Verificar templates
		templatesDir := cfg.Data.TemplatesDir
		if info, err := os.Stat(templatesDir); err == nil && info.IsDir() {
			fmt.Println(utils.Colorize("✅ Directorio de templates encontrado: "+templatesDir, utils.Green))
		} else {
			fmt.Println(utils.Colorize("❌ No se encontró el directorio de templates: "+templatesDir, utils.Red))
			ok = false
		}

		// Verificar recompilación
		for name, cv := range cfg.Data.CVs {
			templatePath := filepath.Join(templatesDir, cv.Template)
			if _, err := os.Stat(templatePath); err != nil {
				fmt.Println(utils.Colorize("❌ Template faltante para '"+name+"': "+templatePath, utils.Red))
				ok = false
				continue
			}

			fmt.Println(utils.Colorize("✅ Template para '"+name+"' OK", utils.Green))

			if config.TemplateNeedsRecompile(templatePath, cv.LastCompileTime()) {
				fmt.Println(utils.Colorize("🔄 Template de '"+name+"' fue modificado recientemente", utils.Yellow))
				needsCompile = append(needsCompile, name)
			}
		}

		// Verificar typst
		if _, err := exec.LookPath("typst"); err != nil {
			fmt.Println(utils.Colorize("❌ typst no está en el PATH", utils.Red))
			ok = false
		} else {
			fmt.Println(utils.Colorize("✅ typst encontrado en el PATH", utils.Green))
		}

		fmt.Println()

		if ok && len(needsCompile) == 0 {
			fmt.Println(utils.Colorize("🎉 Todo está listo para compilar 🎯", utils.Cyan))
			return
		}

		if len(needsCompile) > 0 {
			fmt.Println(utils.Colorize("⚠️  Hay CVs cuyo template fue modificado recientemente:", utils.Yellow))
			for _, name := range needsCompile {
				fmt.Println("  •", utils.Colorize(name, utils.Purple))
			}

			fmt.Print(utils.Colorize("\n¿Deseas compilarlos ahora? [s/N]: ", utils.Cyan))
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))

			if input == "s" || input == "sí" || input == "si" {
				// Flag to track if any CV was compiled successfully
				compiled := false
				for _, name := range needsCompile {
					cv := cfg.Data.CVs[name]
					templatePath := filepath.Join(cfg.Data.TemplatesDir, cv.Template)

					if err := compile.CV(name, cv, cfg.Data.OutputDir, templatePath); err == nil {
						fmt.Println(utils.Colorize(fmt.Sprintf("✅ Compilado: %s", filepath.Join(cfg.Data.OutputDir, name)), utils.Green))

						// Correctly update the in-memory struct
						updatedCV := cfg.Data.CVs[name]
						updatedCV.LastCompile = utils.NowRFC3339()
						cfg.Data.CVs[name] = updatedCV

						compiled = true
					} else {
						fmt.Println(utils.Colorize(fmt.Sprintf("❌ Error compilando %s: %v", name, err), utils.Red))
					}
				}

				// Save the configuration only if there were successful compilations
				if compiled {
					if err := cfg.Save(); err != nil {
						fmt.Println(utils.Colorize(fmt.Sprintf("❌ Error al guardar la configuración: %v", err), utils.Red))
					}
				}
			} else {
				fmt.Println(utils.Colorize("ℹ️  Compilación omitida", utils.Yellow))
			}
		}

		if !ok {
			fmt.Println(utils.Colorize("\n⚠️  Hay problemas que resolver antes de compilar", utils.Red))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
