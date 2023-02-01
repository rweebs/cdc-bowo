package config

// import "fmt"

// type cacheConfig struct {
// 	CacheConfig *cacheConfigImpl `mapstructure:"cache"`
// }

// type cacheConfigImpl struct {
// 	Host     string `mapstructure:"host"`
// 	Port     int    `mapstructure:"port"`
// 	Password string `mapstructure:"password"`
// 	Database string `mapstructure:"database"`
// }

// type CacheConfig interface {
// 	Address() string
// 	Password() string
// 	Database() string
// }

// func NewCache(viperLoader ViperLoader) CacheConfig {
// 	cfg := &cacheConfig{}

// 	viperLoader.Unmarshal(cfg)

// 	return cfg
// }

// func (a *cacheConfig) Password() string {
// 	return a.CacheConfig.Password
// }

// func (a *cacheConfig) Database() string {
// 	return a.CacheConfig.Database
// }

// func (a *cacheConfig) Address() string {
// 	return fmt.Sprintf("%s:%d", a.CacheConfig.Host, a.CacheConfig.Port)
// }
