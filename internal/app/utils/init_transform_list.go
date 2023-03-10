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

func InitTransformListNew(query string, configs config.DDLTransform) map[string]config.DDLTransform {
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
			schemaTable := fmt.Sprintf("%s.%s", schemaName, createTables.Relation.Relname)
			object := transformList[schemaTable]
			object.CreateTable = append(object.CreateTable, config.CreateTable{
				Schema: schemaName,
				Table:  createTables.Relation.Relname,
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
			// log.Println(items)
			// if len(items) > 1 {
			// 	myString := items[0]
			// 	// log.Println("myString", myString)
			// 	start := strings.Index(myString, "sval:\"") + len("sval:\"")
			// 	end := strings.Index(myString[start:], "\"}") + start
			// 	schemaName = myString[start:end]
			// 	myString2 := items[1]
			// 	start2 := strings.Index(myString2, "sval:\"") + len("sval:\"")
			// 	end2 := strings.Index(myString2[start2:], "\"}")
			// 	// log.Println("end2", end2)
			// 	tableName = myString2[start2 : end2+start2]
			// 	// log.Println("tableName", tableName)
			// } else {
			// 	schemaName = "public"
			// 	myString := items[0]
			// 	start := strings.Index(myString, "sval:\"") + len("sval:\"")
			// 	end := strings.Index(myString[start:], "\"}") + start
			// 	tableName = myString[start:end]
			// }
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
