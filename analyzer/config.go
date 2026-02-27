package analyzer

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Rules         RulesConfig         `yaml:"rules"`
	SensitiveData SensitiveDataConfig `yaml:"sensitive_data"`
}

type RulesConfig struct {
	CheckFirstLetter   bool `yaml:"check_first_letter"`
	CheckEnglish       bool `yaml:"check_english"`
	CheckSpecialChars  bool `yaml:"check_special_chars"`
	CheckSensitiveData bool `yaml:"check_sensitive_data"`
}
type SensitiveDataConfig struct {
	UseDefaultPatterns bool     `yaml:"use_default_patterns"`
	CustomPatterns     []string `yaml:"custom_patterns"`
}

func default_config() *Config {
	return &Config{
		Rules: RulesConfig{
			CheckFirstLetter:   true,
			CheckEnglish:       true,
			CheckSpecialChars:  true,
			CheckSensitiveData: true,
		},
		SensitiveData: SensitiveDataConfig{
			UseDefaultPatterns: true,
			CustomPatterns:     []string{},
		},
	}
}

func load_config(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) get_sensitive_patterns() []string {
	patterns := []string{}

	if c.SensitiveData.UseDefaultPatterns {
		defaultPatterns := []string{
			"password",
			"secret",
			"token",
			"key",
		}
		patterns = append(patterns, defaultPatterns...)
	}

	patterns = append(patterns, c.SensitiveData.CustomPatterns...)

	return patterns
}
