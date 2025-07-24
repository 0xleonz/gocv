package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type CVConfig struct {
	Description     string            `mapstructure:"description"`
	LongDescription string            `mapstructure:"long_description"`
	Template        string            `mapstructure:"template"`
	Source          string            `mapstructure:"source"`
	LastCompile     string            `mapstructure:"last_compile"` // chatgpt: almacenado como string
	Vars            map[string]string `mapstructure:"vars"`
}

// string LastCompile -> *time.Time
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
	TemplatesDir    string							`mapstructure:"templates"`
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
	return lc.Viper.WriteConfig()
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

