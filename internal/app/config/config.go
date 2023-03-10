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
	Name             string `mapstructure:"database"`
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	MaxOpenConns     int    `mapstructure:"maxOpenConns"`
	MaxIdleConns     int    `mapstructure:"maxIdleConns"`
	SubscriptionName string `mapstructure:"subscriptionName"`
}

type cacheConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type CreateTable struct {
	Schema string `mapstructure:"schema"`
	Table  string `mapstructure:"table"`
}
type DropTable struct {
	Schema string `mapstructure:"schema"`
	Table  string `mapstructure:"table"`
}
type AddColumn struct {
	Schema string `mapstructure:"schema"`
	Table  string `mapstructure:"table"`
	Column string `mapstructure:"column"`
}
type DropColumn struct {
	Schema     string `mapstructure:"schema"`
	Table      string `mapstructure:"table"`
	Column     string `mapstructure:"column"`
	IsNullable bool   `mapstructure:"isNullable"`
	Default    int    `mapstructure:"default"`
}
type RenameColumn struct {
	Schema  string `mapstructure:"schema"`
	Table   string `mapstructure:"table"`
	OldName string `mapstructure:"oldName"`
	NewName string `mapstructure:"newName"`
}
type RenameTable struct {
	Schema  string `mapstructure:"schema"`
	OldName string `mapstructure:"oldName"`
	NewName string `mapstructure:"newName"`
}
type DropNotNullConstraint struct {
	Schema  string `mapstructure:"schema"`
	Table   string `mapstructure:"table"`
	Column  string `mapstructure:"column"`
	Default int    `mapstructure:"default"`
}
type ModifyDataType struct {
	Schema  string `mapstructure:"schema"`
	Table   string `mapstructure:"table"`
	Column  string `mapstructure:"column"`
	OldType string `mapstructure:"oldType"`
	NewType string `mapstructure:"newType"`
}

type HorizontalSplitting struct {
	Schema      string `mapstructure:"schema"`
	SourceTable string `mapstructure:"sourceTable"`
	DestTable   string `mapstructure:"destTable"`
	Criteria    string `mapstructure:"criteria"`
}
type DDLTransform struct {
	CreateTable           []CreateTable           `mapstructure:"createTable"`
	DropTable             []DropTable             `mapstructure:"dropTable"`
	AddColumn             []AddColumn             `mapstructure:"addColumn"`
	DropColumn            []DropColumn            `mapstructure:"dropColumn"`
	RenameColumn          []RenameColumn          `mapstructure:"renameColumn"`
	RenameTable           []RenameTable           `mapstructure:"renameTable"`
	DropNotNullConstraint []DropNotNullConstraint `mapstructure:"dropNotNullConstraint"`
	ModifyDataType        []ModifyDataType        `mapstructure:"modifyDataType"`
	HorizontalSplitting   []HorizontalSplitting   `mapstructure:"horizontalSplitting"`
}
type Config struct {
	SourceConfig SourceConfig `mapstructure:"source"`
	DestConfig   destConfig   `mapstructure:"dest"`
	CacheConfig  cacheConfig  `mapstructure:"redis"`
	SQLScript    string       `mapstructure:"sqlScript"`
	SQLFile      string       `mapstructure:"sqlFile"`
	DDLTransform DDLTransform `mapstructure:"ddlTransform"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
