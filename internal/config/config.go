package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type CVConfig struct {
	Description     string            `mapstructure:"description"`
	LongDescription string            `mapstructure:"long_description"`
	Template        string            `mapstructure:"template"`
	Source          string            `mapstructure:"source"`
	LastCompile     string            `mapstructure:"last_compile"`
	Vars            map[string]string `mapstructure:"vars"`
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
	OutputDir       string              `mapstructure:"output_dir"`
	DefaultTemplate string              `mapstructure:"default_template"`
	TemplatesDir    string              `mapstructure:"templates"`
	CVs             map[string]CVConfig `mapstructure:"cvs"`
}

type LoadedConfig struct {
	Data  *Config
	Viper *viper.Viper
}

func Load() (*LoadedConfig, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "gocv", "config.yml")

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error leyendo config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
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
		Data:  &cfg,
		Viper: v,
	}, nil
}

func (lc *LoadedConfig) Save() error {
	return saveWithYamlV3(lc.Data, lc.Viper.ConfigFileUsed())
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
	if strings.HasPrefix(p, "~") {
		home, _ := os.UserHomeDir()
		p = strings.Replace(p, "~", home, 1)
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
	return info.ModTime().After(lastCompile.Add(-15 * time.Second))
}
