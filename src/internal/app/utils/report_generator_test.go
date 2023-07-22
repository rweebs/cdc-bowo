package utils

import (
	"reflect"
	"testing"

	"github.com/rweebs/cdc-bowo/internal/app/config"
)

func TestTransformationList(t *testing.T) {
	var tests = []struct {
		testName string
		config   config.DDLTransform
		expected string
	}{
		{
			testName: "test transformation list",
			// Alter table public.test_temporal rename column time_mili to time2; Alter table public.test_temporal rename column timestamp_mili to timestamp2; Alter table public.test_temporal rename column timestamp to timestamp_with_time_zone2; Alter table  public.test_temporal alter column time2 type time; Alter table  public.test_temporal alter column timestamp2 type timestamp; Alter table  public.test_temporal alter column timestamp_with_time_zone2 type timestamp with time zone;
			config: config.DDLTransform{
				CreateTable: []config.CreateTable{config.CreateTable{}},
			},
			expected: "Create Table",
		},
		{
			testName: "test transformation list",
			// Alter table public.test_temporal rename column time_mili to time2; Alter table public.test_temporal rename column timestamp_mili to timestamp2; Alter table public.test_temporal rename column timestamp to timestamp_with_time_zone2; Alter table  public.test_temporal alter column time2 type time; Alter table  public.test_temporal alter column timestamp2 type timestamp; Alter table  public.test_temporal alter column timestamp_with_time_zone2 type timestamp with time zone;
			config: config.DDLTransform{
				CreateTable: []config.CreateTable{config.CreateTable{}}, DropTable: []config.DropTable{config.DropTable{}},
			},
			expected: "Create Table, Drop Table",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			transformList := getTransformationList(tt.config)
			if !reflect.DeepEqual(transformList, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, transformList)
			}

		})
	}
}
