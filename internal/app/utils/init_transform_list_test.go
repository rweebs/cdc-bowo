package utils

import (
	"reflect"
	"testing"

	"github.com/rweebs/cdc-bowo/internal/app/config"
)

func TestInitTransform(t *testing.T) {
	var tests = []struct {
		testName string
		query    string
		expected map[string]config.DDLTransform
	}{
		{
			testName: "test add column",
			// Alter table public.test_temporal rename column time_mili to time2; Alter table public.test_temporal rename column timestamp_mili to timestamp2; Alter table public.test_temporal rename column timestamp to timestamp_with_time_zone2; Alter table  public.test_temporal alter column time2 type time; Alter table  public.test_temporal alter column timestamp2 type timestamp; Alter table  public.test_temporal alter column timestamp_with_time_zone2 type timestamp with time zone;
			query: `ALTER TABLE test.table_name ADD COLUMN new_column_name int not null default 0;`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					AddColumn: []config.AddColumn{
						{
							Schema: "test",
							Table:  "table_name",
							Column: "new_column_name",
						}},
				}},
		},
		{
			testName: "test drop column",
			// Alter table public.test_temporal rename column time_mili to time2; Alter table public.test_temporal rename column timestamp_mili to timestamp2; Alter table public.test_temporal rename column timestamp to timestamp_with_time_zone2; Alter table  public.test_temporal alter column time2 type time; Alter table  public.test_temporal alter column timestamp2 type timestamp; Alter table  public.test_temporal alter column timestamp_with_time_zone2 type timestamp with time zone;
			query: `ALTER TABLE test.table_name Drop COLUMN old_column_name`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					DropColumn: []config.DropColumn{
						{
							Schema: "test",
							Table:  "table_name",
							Column: "old_column_name",
						}},
				}},
		},
		{
			testName: "test rename column",

			query: `ALTER TABLE test.table_name RENAME COLUMN old_column_name TO new_column_name`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					RenameColumn: []config.RenameColumn{
						{
							Schema:  "test",
							Table:   "table_name",
							OldName: "old_column_name",
							NewName: "new_column_name",
						}},
				}},
		},
		{
			testName: "test rename table",

			query: `ALTER TABLE test.table_name RENAME TO new_table_name`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					RenameTable: []config.RenameTable{
						{
							Schema:  "test",
							OldName: "table_name",
							NewName: "new_table_name",
						}},
				}},
		},
		{
			testName: "test drop column",

			query: `ALTER TABLE test.table_name DROP COLUMN old_column_name`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					DropColumn: []config.DropColumn{
						{
							Schema: "test",
							Table:  "table_name",
							Column: "old_column_name",
						}},
				}},
		},
		{
			testName: "test drop table",

			query: `DROP TABLE test.table_name`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					DropTable: []config.DropTable{
						{
							Schema: "test",
							Table:  "table_name",
						}},
				}},
		},
		{
			testName: "test create table",

			query: `CREATE TABLE test.table_name (id int)`,
			expected: map[string]config.DDLTransform{
				"test.table_name": {
					CreateTable: []config.CreateTable{
						{
							Schema: "test",
							Table:  "table_name",
						}},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			transformList := InitTransformListNew(tt.query, config.DDLTransform{})
			if !reflect.DeepEqual(transformList, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, transformList)
			}

		})
	}
}
