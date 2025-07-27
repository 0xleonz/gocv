package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/0xleonz/gocv/internal/compile"
	"gitlab.com/0xleonz/gocv/internal/config"
	"gitlab.com/0xleonz/gocv/internal/utils"
)

var selectFlag bool

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compila uno o más currículums usando Typst",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := AppConfig

		if selectFlag {
			selectAndCompile(cfg)
		} else {
			compileAllIfModified(cfg)
		}

		_ = cfg.Save()
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
	compileCmd.Flags().BoolVarP(&selectFlag, "select", "s", false, "Selecciona un CV a compilar")
}

// Compilar CVs (15s)
func compileAllIfModified(cfg *config.LoadedConfig) {
	for name, cv := range cfg.Data.CVs {
		templatePath := filepath.Join(cfg.Data.TemplatesDir, cv.Template)
		if config.TemplateNeedsRecompile(templatePath, cv.LastCompileTime()) {
			compile.CV(name, cv, cfg.Data.OutputDir, templatePath)
			cv.LastCompile = utils.NowRFC3339()
			cfg.Data.CVs[name] = cv
		}
	}
}

// seleccionado manualmente
func selectAndCompile(cfg *config.LoadedConfig) {
	cvs := cfg.Data.CVs

	keys := []string{}
	i := 1
	fmt.Println(utils.Colorize("📄 CVs disponibles:", utils.Cyan))
	for name, cv := range cvs {
		fmt.Printf("  %d. %s - %s\n", i, utils.Colorize(name, utils.Purple), cv.Description)
		keys = append(keys, name)
		i++
	}

	fmt.Print(utils.Colorize("\nSeleccione un número: ", utils.Yellow))
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	index, _ := strconv.Atoi(line[:len(line)-1])

	if index < 1 || index > len(keys) {
		fmt.Println(utils.Colorize("❌ Selección inválida", utils.Red))
		return
	}

	name := keys[index-1]
	cv := cvs[name]
	templatePath := filepath.Join(cfg.Data.TemplatesDir, cv.Template)
	compile.CV(name, cv, cfg.Data.OutputDir, templatePath)

	fmt.Println(utils.Colorize("\n📘 Descripción larga:\n", utils.Green))
	fmt.Println(cv.LongDescription)

	cv.LastCompile = utils.NowRFC3339()
	cfg.Data.CVs[name] = cv
}
