package main

import (
	"fmt"
	"strings"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

func parse_vertical_splitting() {
	parse, err := pg_query.Parse(`
	INSERT INTO public.vertical_splitting_derived_1 (id, message) SELECT id, message FROM public.vertical_splitting;
	INSERT INTO public.vertical_splitting_derived_2 (id, message2) SELECT id, message2 FROM public.vertical_splitting;
`)
	if err != nil {
		panic(err)
	}
	for _, v := range parse.GetStmts() {
		// cols := []string{}
		//InsertStatement
		insertStatement := v.Stmt.GetInsertStmt()
		// childTableSchema := insertStatement.Relation.Schemaname
		// childTableRel := insertStatement.Relation.Relname
		// childTableName := fmt.Sprintf("%s.%s", childTableSchema, childTableRel)
		// fmt.Println(childTableName)
		// childCols := insertStatement.Cols
		// // fmt.Println(childCols)
		// for _, col := range childCols {
		// 	// fmt.Println(col.GetResTarget().GetName())
		// 	cols = append(cols, col.GetResTarget().GetName())
		// }
		// fmt.Println(cols)
		parentTableRel := insertStatement.SelectStmt.GetSelectStmt().GetFromClause()[0].GetRangeVar().GetRelname()
		parentTableSchema := insertStatement.SelectStmt.GetSelectStmt().GetFromClause()[0].GetRangeVar().GetSchemaname()
		parentTableName := fmt.Sprintf("%s.%s", parentTableSchema, parentTableRel)
		fmt.Println(parentTableName)

	}

}

func parse_horizontal_splitting() {
	parse, err := pg_query.Parse(`
	INSERT INTO public.table_id_after_2000 (id, message) SELECT id, message FROM public.table_id WHERE id > 2000;
`)
	if err != nil {
		panic(err)
	}
	for _, v := range parse.GetStmts() {
		// cols := []string{}
		//InsertStatement
		// insertStatement := v.Stmt.GetInsertStmt()
		// childTableSchema := insertStatement.Relation.Schemaname
		// childTableRel := insertStatement.Relation.Relname
		// childTableName := fmt.Sprintf("%s.%s", childTableSchema, childTableRel)
		// fmt.Println(childTableName)
		// childCols := insertStatement.Cols
		// // fmt.Println(childCols)
		// for _, col := range childCols {
		// 	// fmt.Println(col.GetResTarget().GetName())
		// 	cols = append(cols, col.GetResTarget().GetName())
		// }
		// fmt.Println(cols)
		// parentTableRel := insertStatement.SelectStmt.GetSelectStmt().GetFromClause()[0].GetRangeVar().GetRelname()
		// parentTableSchema := insertStatement.SelectStmt.GetSelectStmt().GetFromClause()[0].GetRangeVar().GetSchemaname()
		// parentTableName := fmt.Sprintf("%s.%s", parentTableSchema, parentTableRel)
		statement, _ := pg_query.Deparse(&pg_query.ParseResult{Stmts: []*pg_query.RawStmt{v}})
		fmt.Println(statement)
		whereClause := strings.Split(statement, "WHERE")[1]
		fmt.Println(whereClause)

	}

}
func main() {
	// parse_vertical_splitting()
	parse_horizontal_splitting()
	// parse_comment()
}
