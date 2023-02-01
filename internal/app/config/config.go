package config

import "github.com/spf13/viper"

type SourceConfig struct {
	Name         string `mapstructure:"database"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type destConfig struct {
	Name         string `mapstructure:"database"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type cacheConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
type DDLTransform struct {
	CreateTable []struct {
		Schema string `mapstructure:"schema"`
		Table  string `mapstructure:"table"`
	} `mapstructure:"createTable"`
	DropTable []struct {
		Schema string `mapstructure:"schema"`
		Table  string `mapstructure:"table"`
	} `mapstructure:"dropTable"`
	AddColumn []struct {
		Schema string `mapstructure:"schema"`
		Table  string `mapstructure:"table"`
		Column string `mapstructure:"column"`
	} `mapstructure:"addColumn"`
	DropColumn []struct {
		Schema     string      `mapstructure:"schema"`
		Table      string      `mapstructure:"table"`
		Column     string      `mapstructure:"column"`
		IsNullable bool        `mapstructure:"isNullable"`
		Default    interface{} `mapstructure:"default"`
	} `mapstructure:"dropColumn"`
	RenameColumn []struct {
		Schema  string `mapstructure:"schema"`
		Table   string `mapstructure:"table"`
		OldName string `mapstructure:"oldName"`
		NewName string `mapstructure:"newName"`
	} `mapstructure:"renameColumn"`
	RenameTable []struct {
		Schema  string `mapstructure:"schema"`
		OldName string `mapstructure:"oldName"`
		NewName string `mapstructure:"newName"`
	} `mapstructure:"renameTable"`
	DropNotNullConstraint []struct {
		Schema  string      `mapstructure:"schema"`
		Table   string      `mapstructure:"table"`
		Column  string      `mapstructure:"column"`
		Default interface{} `mapstructure:"default"`
	} `mapstructure:"dropNotNullConstraint"`
	ModifyDataType []struct {
		Schema  string `mapstructure:"schema"`
		Table   string `mapstructure:"table"`
		Column  string `mapstructure:"column"`
		OldType string `mapstructure:"oldType"`
		NewType string `mapstructure:"newType"`
	} `mapstructure:"modifyDataType"`
}
type Config struct {
	SourceConfig SourceConfig `mapstructure:"source"`
	DestConfig   destConfig   `mapstructure:"dest"`
	CacheConfig  cacheConfig  `mapstructure:"cache"`
	SQLScript    string       `mapstructure:"sqlScript"`
	DDLTransform DDLTransform `mapstructure:"ddlTransform"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
