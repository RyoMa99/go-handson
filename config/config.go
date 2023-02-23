package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err)
	}

	return &cfg, nil
}
