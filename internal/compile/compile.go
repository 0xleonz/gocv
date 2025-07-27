package compile

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gitlab.com/0xleonz/gocv/internal/config"
	"gitlab.com/0xleonz/gocv/internal/utils"
)

func CV(name string, cv config.CVConfig, outputDir string, templatePath string) error {
	// source := filepath.Join(filepath.Dir(templatePath), name+".typ")
	output := filepath.Join(outputDir, name+".pdf")

	fmt.Println(utils.Colorize("🛠️  Compilando "+name+"...", utils.Pink))
	// fmt.Println(templatePath)
	// fmt.Println(source)
	// fmt.Println(output)
	cmd := exec.Command("typst", "compile", templatePath, output)

	// cmd.Env = append(os.Environ(),
	// 	"TYPST_FONT_PATH="+filepath.Dir(templatePath),
	// )

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(utils.Colorize("❌ Falló la compilación: "+err.Error(), utils.Red))
		return err
	}

	fmt.Println(utils.Colorize("✅ Compilado: "+output, utils.Green))
	return nil
}
