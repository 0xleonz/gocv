package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type CVConfig struct {
	Description     string            `yaml:"description"`
	LongDescription string            `yaml:"long_description"`
	Template        string            `yaml:"template"`
	Source          string            `yaml:"source"`
	LastCompile     string            `yaml:"last_compile"`
	Vars            map[string]string `yaml:"vars"`
}

func (cv CVConfig) LastCompileTime() *time.Time {
	if cv.LastCompile == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, cv.LastCompile)
	if err != nil {
		return nil
	}
	return &t
}

type Config struct {
	OutputDir       string              `yaml:"output_dir"`
	DefaultTemplate string              `yaml:"default_template"`
	TemplatesDir    string              `yaml:"templates"`
	CVs             map[string]CVConfig `yaml:"cvs"`
}

type LoadedConfig struct {
	Data *Config
	Path string
}

func Load() (*LoadedConfig, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "gocv", "config.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parseando config: %w", err)
	}

	cfg.OutputDir = expandPath(cfg.OutputDir)
	cfg.TemplatesDir = expandPath(cfg.TemplatesDir)

	for key, cv := range cfg.CVs {
		cv.Template = expandPath(cv.Template)
		cv.Source = expandPath(cv.Source)
		cfg.CVs[key] = cv
	}

	return &LoadedConfig{
		Data: &cfg,
		Path: configPath,
	}, nil
}

func (lc *LoadedConfig) Save() error {
	return saveWithYamlV3(lc.Data, lc.Path)
}

func saveWithYamlV3(cfg *Config, path string) error {
	root := yaml.Node{Kind: yaml.MappingNode}

	root.Content = append(root.Content,
		yamlScalar("output_dir"), yamlScalar(cfg.OutputDir),
		yamlScalar("default_template"), yamlScalar(cfg.DefaultTemplate),
		yamlScalar("templates"), yamlScalar(cfg.TemplatesDir),
	)

	cvsKey := yamlScalar("cvs")
	cvsMap := &yaml.Node{Kind: yaml.MappingNode}

	for name, cv := range cfg.CVs {
		key := yamlScalar(name)
		val := &yaml.Node{Kind: yaml.MappingNode}

		// Campos básicos
		descK, descV := yamlScalar("description"), yamlScalar(cv.Description)
		longK, longV := yamlScalarWithStyle("long_description", cv.LongDescription, yaml.LiteralStyle)
		templK, templV := yamlScalar("template"), yamlScalar(cv.Template)
		sourceK, sourceV := yamlScalar("source"), yamlScalar(cv.Source)
		lastK, lastV := yamlScalar("last_compile"), yamlScalar(cv.LastCompile)

		val.Content = append(val.Content,
			descK, descV,
			longK, longV,
			templK, templV,
			sourceK, sourceV,
			lastK, lastV,
		)

		// Variables extra
		if len(cv.Vars) > 0 {
			varsNode := &yaml.Node{Kind: yaml.MappingNode}
			for k, v := range cv.Vars {
				varsNode.Content = append(varsNode.Content, yamlScalar(k), yamlScalar(v))
			}
			val.Content = append(val.Content, yamlScalar("vars"), varsNode)
		}

		cvsMap.Content = append(cvsMap.Content, key, val)
	}
	root.Content = append(root.Content, cvsKey, cvsMap)

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("no se pudo crear archivo config: %w", err)
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(2)
	return enc.Encode(&root)
}

func yamlScalar(value string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: value,
	}
}

func yamlScalarWithStyle(key, value string, style yaml.Style) (*yaml.Node, *yaml.Node) {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: key},
		&yaml.Node{Kind: yaml.ScalarNode, Value: value, Style: style}
}

func expandPath(p string) string {
	if strings.HasPrefix(p, "~"+string(os.PathSeparator)) {
		home, _ := os.UserHomeDir()
		p = filepath.Join(home, p[2:])
	}
	return os.ExpandEnv(p)
}

func TemplateNeedsRecompile(templatePath string, lastCompile *time.Time) bool {
	info, err := os.Stat(templatePath)
	if err != nil {
		return true
	}
	if lastCompile == nil {
		return true
	}
	// revisa el template es más reciente (con margen de 15s)
	return info.ModTime().After(lastCompile.Add(-15 * time.Second))
}
