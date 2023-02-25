package utils

import (
	"fmt"

	"github.com/rweebs/cdc-bowo/internal/app/config"
)

func InitTransformList(configs config.DDLTransform) map[string]config.DDLTransform {
	transformList := map[string]config.DDLTransform{}
	//TODO:
	for _, v := range configs.AddColumn {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.AddColumn = append(object.AddColumn, v)
		transformList[schemaTable] = object
	}
	//DONE:
	for _, v := range configs.DropColumn {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.DropColumn = append(object.DropColumn, v)
		transformList[schemaTable] = object
	}
	//TODO:
	for _, v := range configs.DropNotNullConstraint {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.DropNotNullConstraint = append(object.DropNotNullConstraint, v)
		transformList[schemaTable] = object
	}
	//DONE:
	for _, v := range configs.ModifyDataType {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.ModifyDataType = append(object.ModifyDataType, v)
		transformList[schemaTable] = object
	}
	//DONE:
	for _, v := range configs.RenameColumn {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.RenameColumn = append(object.RenameColumn, v)
		transformList[schemaTable] = object
	}
	//DONE:
	for _, v := range configs.RenameTable {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.OldName)
		object := transformList[schemaTable]
		object.RenameTable = append(object.RenameTable, v)
		transformList[schemaTable] = object
	}
	//DONE:
	for _, v := range configs.DropTable {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.DropTable = append(object.DropTable, v)
		transformList[schemaTable] = object
	}
	for _, v := range configs.CreateTable {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.CreateTable = append(object.CreateTable, v)
		transformList[schemaTable] = object
	}
	return transformList
}
