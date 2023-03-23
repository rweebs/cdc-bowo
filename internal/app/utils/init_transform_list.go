package utils

import (
	"encoding/json"
	"fmt"
	"log"

	pg_query "github.com/pganalyze/pg_query_go/v4"
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

func InitTransformListFromQuery(query string, configs config.DDLTransform) map[string]config.DDLTransform {
	transformList := map[string]config.DDLTransform{}
	parse, err := pg_query.Parse(query)
	if err != nil {
		log.Println(err)
	}
	for _, v := range parse.GetStmts() {
		//AddColumn
		alterTables := v.Stmt.GetAlterTableStmt()
		renames := v.Stmt.GetRenameStmt()
		createTables := v.Stmt.GetCreateStmt()
		dropTables := v.Stmt.GetDropStmt()
		if alterTables != nil && alterTables.Cmds[0].GetAlterTableCmd().Subtype == pg_query.AlterTableType_AT_AddColumn {
			var schemaName string
			if alterTables.Relation.Schemaname != "" {
				schemaName = alterTables.Relation.Schemaname
			} else {
				schemaName = "public"
			}
			schemaTable := fmt.Sprintf("%s.%s", schemaName, alterTables.Relation.Relname)
			object := transformList[schemaTable]
			object.AddColumn = append(object.AddColumn, config.AddColumn{
				Schema: schemaName,
				Table:  alterTables.Relation.Relname,
				Column: alterTables.Cmds[0].GetAlterTableCmd().GetDef().GetColumnDef().Colname,
			})
			transformList[schemaTable] = object

		} else if alterTables != nil && alterTables.Cmds[0].GetAlterTableCmd().Subtype == pg_query.AlterTableType_AT_DropColumn {
			//DropColumn
			var schemaName string
			if alterTables.Relation.Schemaname != "" {
				schemaName = alterTables.Relation.Schemaname
			} else {
				schemaName = "public"
			}
			schemaTable := fmt.Sprintf("%s.%s", schemaName, alterTables.Relation.Relname)
			object := transformList[schemaTable]
			object.DropColumn = append(object.DropColumn, config.DropColumn{
				Schema: schemaName,
				Table:  alterTables.Relation.Relname,
				Column: alterTables.Cmds[0].GetAlterTableCmd().Name,
			})
			transformList[schemaTable] = object
		} else if renames != nil && renames.RenameType == pg_query.ObjectType_OBJECT_COLUMN {
			//RenameColumn
			var schemaName string
			if renames.Relation.Schemaname != "" {
				schemaName = renames.Relation.Schemaname
			} else {
				schemaName = "public"
			}
			schemaTable := fmt.Sprintf("%s.%s", schemaName, renames.Relation.Relname)
			object := transformList[schemaTable]
			object.RenameColumn = append(object.RenameColumn, config.RenameColumn{
				Schema:  schemaName,
				Table:   renames.Relation.Relname,
				OldName: renames.Subname,
				NewName: renames.Newname,
			})
			transformList[schemaTable] = object
		} else if renames != nil && renames.RenameType == pg_query.ObjectType_OBJECT_TABLE {
			//RenameTable
			var schemaName string
			if renames.Relation.Schemaname != "" {
				schemaName = renames.Relation.Schemaname
			} else {
				schemaName = "public"
			}
			schemaTable := fmt.Sprintf("%s.%s", schemaName, renames.Relation.Relname)
			object := transformList[schemaTable]
			object.RenameTable = append(object.RenameTable, config.RenameTable{
				Schema:  schemaName,
				OldName: renames.Relation.Relname,
				NewName: renames.Newname,
			})
			transformList[schemaTable] = object
			schemaTable = fmt.Sprintf("%s.%s", schemaName, renames.Newname)
			object = transformList[schemaTable]
			object.RenameTable = append(object.RenameTable, config.RenameTable{
				Schema:  schemaName,
				OldName: renames.Relation.Relname,
				NewName: renames.Newname,
			})
			transformList[schemaTable] = object
		} else if alterTables != nil && alterTables.Cmds[0].GetAlterTableCmd().Subtype == pg_query.AlterTableType_AT_DropColumn {
			//DropColumn
			var schemaName string
			if alterTables.Relation.Schemaname != "" {
				schemaName = alterTables.Relation.Schemaname
			} else {
				schemaName = "public"
			}
			schemaTable := fmt.Sprintf("%s.%s", schemaName, alterTables.Relation.Relname)
			object := transformList[schemaTable]
			object.DropColumn = append(object.DropColumn, config.DropColumn{
				Schema: schemaName,
				Table:  alterTables.Relation.Relname,
				Column: alterTables.Cmds[0].GetAlterTableCmd().Name,
			})
			transformList[schemaTable] = object
		} else if createTables != nil {
			//CreateTable
			var schemaName string
			if createTables.Relation.Schemaname != "" {
				schemaName = createTables.Relation.Schemaname
			} else {
				schemaName = "public"
			}
			type CreateTableColumn struct {
				Node struct {
					ColumnDef struct {
						Colname string `json:"colname"`
					} `json:"ColumnDef"`
				} `json:"Node"`
			}
			tableItems, _ := json.Marshal(createTables.TableElts)
			tableItemsColumns := []CreateTableColumn{}
			_ = json.Unmarshal(tableItems, &tableItemsColumns)
			var columnName []string
			for _, v := range tableItemsColumns {
				columnName = append(columnName, v.Node.ColumnDef.Colname)
			}
			schemaTable := fmt.Sprintf("%s.%s", schemaName, createTables.Relation.Relname)
			object := transformList[schemaTable]
			object.CreateTable = append(object.CreateTable, config.CreateTable{
				Schema:  schemaName,
				Table:   createTables.Relation.Relname,
				Colname: columnName,
			})
			transformList[schemaTable] = object
		} else if dropTables != nil && dropTables.RemoveType == pg_query.ObjectType_OBJECT_TABLE {

			//DropTable
			var schemaName string
			var tableName string

			dropTableItems, _ := json.Marshal(dropTables.Objects[0].Node)
			type DropTableItems struct {
				List struct {
					Items []struct {
						Node struct {
							String struct {
								Sval string `json:"sval"`
							} `json:"String_"`
						} `json:"Node"`
					} `json:"items"`
				} `json:"List"`
			}
			var dropTableItemsObject DropTableItems
			_ = json.Unmarshal(dropTableItems, &dropTableItemsObject)
			items := dropTableItemsObject.List.Items
			if len(items) > 1 {
				schemaName = items[0].Node.String.Sval
				tableName = items[1].Node.String.Sval
			} else {
				schemaName = "public"
				tableName = items[0].Node.String.Sval
			}

			schemaTable := fmt.Sprintf("%s.%s", schemaName, tableName)
			object := transformList[schemaTable]
			object.DropTable = append(object.DropTable, config.DropTable{
				Schema: schemaName,
				Table:  tableName,
			})
			transformList[schemaTable] = object
		}
	}
	for _, v := range configs.VerticalSplitting {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.SourceTable)
		object := transformList[schemaTable]
		for _, table := range v.DerivedTable {
			v.DerivedTableDetails = append(v.DerivedTableDetails, transformList[fmt.Sprintf("%s.%s", v.Schema, table)].CreateTable...)
		}
		object.VerticalSplitting = append(object.VerticalSplitting, v)
		transformList[schemaTable] = object
	}
	for _, v := range configs.AddColumnInTheMiddle {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.DropTable = []config.DropTable{}
		object.RenameTable = []config.RenameTable{}
		object.AddColumn = append(object.AddColumn, config.AddColumn{
			Schema: v.Schema,
			Table:  v.Table,
			Column: v.Column,
		})
		transformList[schemaTable] = object
	}
	// data, _ := json.MarshalIndent(transformList, "", "  ")
	// fmt.Println(string(data))
	return transformList
}

func InitTransformListNew(query string, configs config.DDLTransform) map[string]config.DDLTransform {
	transformList := InitTransformListFromQuery(query, configs)

	for _, v := range configs.ModifyDataType {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
		object := transformList[schemaTable]
		object.ModifyDataType = append(object.ModifyDataType, v)
		transformList[schemaTable] = object
	}

	for _, v := range configs.HorizontalSplitting {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.SourceTable)
		object := transformList[schemaTable]
		object.HorizontalSplitting = append(object.HorizontalSplitting, v)
		transformList[schemaTable] = object
		//HorizontalSplittingDest
	}

	// //TODO:
	// //DONE:
	//
	// //DONE:
	// for _, v := range configs.RenameTable {
	// 	schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.OldName)
	// 	object := transformList[schemaTable]
	// 	object.RenameTable = append(object.RenameTable, v)
	// 	transformList[schemaTable] = object
	// }
	// //DONE:
	// for _, v := range configs.DropTable {
	// 	schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
	// 	object := transformList[schemaTable]
	// 	object.DropTable = append(object.DropTable, v)
	// 	transformList[schemaTable] = object
	// }
	// for _, v := range configs.CreateTable {
	// 	schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)
	// 	object := transformList[schemaTable]
	// 	object.CreateTable = append(object.CreateTable, v)
	// 	transformList[schemaTable] = object
	// }
	return transformList
}

func InitTransformListDestNew(query string, configs config.DDLTransform) map[string]config.DDLTransform {
	transformList := InitTransformListFromQuery(query, configs)
	for _, v := range configs.ModifyDataType {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.Table)

		for _, renameColumn := range transformList[schemaTable].RenameColumn {
			if renameColumn.Schema == v.Schema && renameColumn.NewName == v.Column {
				v.Column = renameColumn.OldName
			}
		}
		object := transformList[schemaTable]
		object.ModifyDataType = append(object.ModifyDataType, v)
		transformList[schemaTable] = object
	}
	for _, v := range configs.HorizontalSplitting {
		schemaTable := fmt.Sprintf("%s.%s", v.Schema, v.DestTable)
		object := transformList[schemaTable]
		object.RenameTable = append(object.RenameTable, config.RenameTable{
			Schema:  v.Schema,
			OldName: v.SourceTable,
			NewName: v.DestTable,
		})
		transformList[schemaTable] = object
	}
	for _, v := range configs.VerticalSplitting {
		if v.SourceDeleted {
			for _, table := range v.DerivedTable {
				schemaTable := fmt.Sprintf("%s.%s", v.Schema, table)
				object := transformList[schemaTable]
				object.CreateTable[0].BeingUsed = true
				object.CreateTable[0].RelatedTable = fmt.Sprintf("%s.%s", v.Schema, v.SourceTable)
				object.CreateTable[0].RelatedTablePrimaryKeys = v.PrimaryKey
				transformList[schemaTable] = object
			}
		}
	}

	return transformList
}
