package config

// import "fmt"

// type destConfig struct {
// 	DestConfig *destConfigImpl `mapstructure:"dest"`
// }

// type destConfigImpl struct {
// 	Name         string `mapstructure:"dest"`
// 	Host         string `mapstructure:"host"`
// 	Port         int    `mapstructure:"port"`
// 	Username     string `mapstructure:"username"`
// 	Password     string `mapstructure:"password"`
// 	MaxOpenConns int    `mapstructure:"maxOpenConns"`
// 	MaxIdleConns int    `mapstructure:"maxIdleConns"`
// }

// type DestConfig interface {
// }

// func NewDest(viperLoader ViperLoader) DestConfig {
// 	cfg := &destConfig{}

// 	viperLoader.Unmarshal(cfg)

// 	return cfg
// }

// func (a *destConfig) DSN() string {
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", a.DestConfig.Host, a.DestConfig.Username, a.DestConfig.Password, a.DestConfig.Name, a.DestConfig.Port)
// }
