package config

// import "fmt"

// type sourceConfig struct {
// 	SourceConfig *sourceConfigImpl `mapstructure:"source"`
// }

// type sourceConfigImpl struct {
// 	Name         string `mapstructure:"database"`
// 	Host         string `mapstructure:"host"`
// 	Port         int    `mapstructure:"port"`
// 	Username     string `mapstructure:"username"`
// 	Password     string `mapstructure:"password"`
// 	MaxOpenConns int    `mapstructure:"maxOpenConns"`
// 	MaxIdleConns int    `mapstructure:"maxIdleConns"`
// }

// type SourceConfig interface {
// }

// func NewSource(viperLoader ViperLoader) SourceConfig {
// 	cfg := &sourceConfig{}

// 	viperLoader.Unmarshal(cfg)

// 	return cfg
// }

// func (a *sourceConfig) DSN() string {
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", a.SourceConfig.Host, a.SourceConfig.Username, a.SourceConfig.Password, a.SourceConfig.Name, a.SourceConfig.Port)
// }
