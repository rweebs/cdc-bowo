package config

import "github.com/spf13/viper"

type SourceConfig struct {
	Name                string `mapstructure:"database"`
	Host                string `mapstructure:"host"`
	Port                int    `mapstructure:"port"`
	Username            string `mapstructure:"username"`
	Password            string `mapstructure:"password"`
	MaxOpenConns        int    `mapstructure:"maxOpenConns"`
	MaxIdleConns        int    `mapstructure:"maxIdleConns"`
	RedisTopicPrefix    string `mapstructure:"redisTopicPrefix"`
	DebeziumPublication string `mapstructure:"debeziumPublication"`
}

type destConfig struct {
	Name                string `mapstructure:"database"`
	Host                string `mapstructure:"host"`
	Port                int    `mapstructure:"port"`
	Username            string `mapstructure:"username"`
	Password            string `mapstructure:"password"`
	MaxOpenConns        int    `mapstructure:"maxOpenConns"`
	MaxIdleConns        int    `mapstructure:"maxIdleConns"`
	SubscriptionName    string `mapstructure:"subscriptionName"`
	RedisTopicPrefix    string `mapstructure:"redisTopicPrefix"`
	DebeziumPublication string `mapstructure:"debeziumPublication"`
}

type cacheConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type CreateTable struct {
	Schema                  string   `mapstructure:"schema"`
	Table                   string   `mapstructure:"table"`
	Colname                 []string `mapstructure:"colname"`
	BeingUsed               bool     `mapstructure:"beingUsed"`
	RelatedTable            string   `mapstructure:"relatedTable"`
	RelatedTablePrimaryKeys []string `mapstructure:"relatedTablePrimaryKeys"`
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
type AddColumnInTheMiddle struct {
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
type VerticalSplitting struct {
	Schema              string        `json:"schema"`
	SourceTable         string        `json:"sourceTable"`
	DerivedTable        []string      `json:"derivedTable"`
	DerivedTableDetails []CreateTable `json:"derivedTableDetails"`
	SourceDeleted       bool          `json:"sourceDeleted"`
	PrimaryKey          []string      `json:"primaryKey"`
}
type VerticalSplittingDest struct {
	Schema        string   `json:"schema"`
	Table         string   `json:"sourceTable"`
	DerivedTable  []string `json:"derivedTable"`
	SourceDeleted bool     `json:"sourceDeleted"`
	PrimaryKey    string   `json:"primaryKey"`
}

type DDLTransform struct {
	CreateTable           []CreateTable           `mapstructure:"createTable"`
	DropTable             []DropTable             `mapstructure:"dropTable"`
	AddColumn             []AddColumn             `mapstructure:"addColumn"`
	AddColumnInTheMiddle  []AddColumnInTheMiddle  `mapstructure:"addColumnInTheMiddle"`
	DropColumn            []DropColumn            `mapstructure:"dropColumn"`
	RenameColumn          []RenameColumn          `mapstructure:"renameColumn"`
	RenameTable           []RenameTable           `mapstructure:"renameTable"`
	DropNotNullConstraint []DropNotNullConstraint `mapstructure:"dropNotNullConstraint"`
	ModifyDataType        []ModifyDataType        `mapstructure:"modifyDataType"`
	HorizontalSplitting   []HorizontalSplitting   `mapstructure:"horizontalSplitting"`
	VerticalSplitting     []VerticalSplitting     `mapstructure:"verticalSplitting"`
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

func ConvertToAddColumn(config AddColumnInTheMiddle) AddColumn {
	return AddColumn{
		Schema: config.Schema,
		Table:  config.Table,
		Column: config.Column,
	}
}
