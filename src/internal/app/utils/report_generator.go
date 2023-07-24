package utils

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/rweebs/cdc-bowo/internal/app/config"
)

type TableComparisonResult struct {
	Schema          string
	Table           string
	RecordsInDB1    int
	RecordsInDB2    int
	Difference      int
	MissingInDB1    bool
	MissingInDB2    bool
	ChangeOperation string
	Expected        bool
}

func convertDestTableName(renameTables []config.RenameTable, sourceTableName string) string {
	for _, table := range renameTables {
		if table.OldName == sourceTableName {
			return table.NewName
		}
	}
	return sourceTableName
}
func getTransformationList(data config.DDLTransform) string {
	var result string
	valueOfData := reflect.ValueOf(data)
	typeOfData := valueOfData.Type()

	for i := 0; i < valueOfData.NumField(); i++ {
		fieldValue := valueOfData.Field(i)
		fieldName := typeOfData.Field(i).Name

		// Check if the length of the slice for the field is more than 0
		if fieldValue.Len() > 0 {
			// If yes, add the name of the field to the result string
			fieldName = convertToTitleCase(fieldName)
			result += fieldName + ", "
		}
	}

	// Trim the trailing ", " from the result string
	if len(result) > 2 {
		result = result[:len(result)-2]
	}

	return result
}

func GenerateReport(db1 *sql.DB, db2 *sql.DB, transformationList map[string]config.DDLTransform, timestampCutOff int64) {
	// Get the list of schemas and their tables in both databases
	schemasDB1, err := getSchemas(db1)
	if err != nil {
		log.Fatal("Error fetching schema names from the first database:", err)
	}

	schemasDB2, err := getSchemas(db2)
	if err != nil {
		log.Fatal("Error fetching schema names from the second database:", err)
	}

	results := []TableComparisonResult{}

	// Compare the data in each table for each schema in the first database
	for schema, tablesDB1 := range schemasDB1 {
		for _, tableName := range tablesDB1 {
			result := TableComparisonResult{Schema: schema, Table: tableName}
			result.RecordsInDB1, _ = getRowCount(db1, schema, tableName)
			db2Config := transformationList[schema+"."+tableName]
			db2TableName := convertDestTableName(db2Config.RenameTable, tableName)
			result.RecordsInDB2, _ = getRowCount(db2, schema, db2TableName)
			result.Difference = result.RecordsInDB1 - result.RecordsInDB2

			tablesDB2, foundInDB2 := schemasDB2[schema]
			result.ChangeOperation = getTransformationList(transformationList[schema+"."+tableName])
			result.MissingInDB2 = !foundInDB2 || !contains(tablesDB2, tableName)

			if result.Difference == 0 || result.MissingInDB2 {
				result.Expected = true
			} else {
				result.Expected = false
			}
			results = append(results, result)
		}
	}

	// Check for tables present in the second database but missing in the first database
	for schema, tablesDB2 := range schemasDB2 {
		for _, tableName := range tablesDB2 {
			tablesDB1, foundInDB1 := schemasDB1[schema]
			if !foundInDB1 || !contains(tablesDB1, tableName) {
				result := TableComparisonResult{Schema: schema, Table: tableName}
				result.RecordsInDB1 = 0
				result.RecordsInDB2, _ = getRowCount(db2, schema, tableName)
				result.Difference = -result.RecordsInDB2
				result.MissingInDB1 = !foundInDB1 || !contains(tablesDB1, tableName)
				if result.Difference == 0 || result.MissingInDB1 {
					result.Expected = true
				} else {
					result.Expected = false
				}
				results = append(results, result)
			}
		}
	}
	if timestampCutOff == 0 {
		fmt.Println("You are not starting the blue green deployment, yet")
	} else {
		result := fmt.Sprintf("'%s'", time.Unix(0, timestampCutOff*int64(time.Millisecond)).UTC().Format(time.RFC3339))

		fmt.Println("You are starting the blue green deployment from timestamp: ", result)
	}

	fmt.Println("Report generated")
	// Print the results
	fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "Schema", "Table", "Source Records", "Dest Records", "Difference", "Change Operation", "Expected")
	fmt.Printf("%-20s-%-20s-%-20s-%-20s-%-20s-%-20s-%-20s\n", "-------------------", "-------------------", "-------------------", "-------------------", "-------------------", "-------------------", "-------------------")

	for _, result := range results {
		fmt.Printf("%-20s %-20s %-20d %-20d %-20d %-20s %-20v\n", result.Schema, result.Table, result.RecordsInDB1, result.RecordsInDB2, result.Difference, result.ChangeOperation, result.Expected)
	}
}

// getSchemas retrieves the list of schemas and their table names in the database.
func getSchemas(db *sql.DB) (map[string][]string, error) {
	schemas := make(map[string][]string)

	rows, err := db.Query(`
		SELECT table_schema, table_name
		FROM information_schema.tables
where table_schema not in ('pg_catalog','information_schema');
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema, tableName string
		err := rows.Scan(&schema, &tableName)
		if err != nil {
			return nil, err
		}
		schemas[schema] = append(schemas[schema], tableName)
	}

	return schemas, nil
}

// getRowCount retrieves the total number of records in a specific table in a schema.
func getRowCount(db *sql.DB, schema, tableName string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", schema, tableName))
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// contains checks if a slice contains a specific string element.
func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

func convertToTitleCase(input string) string {
	// Use a regular expression to insert a space before each uppercase letter following a lowercase letter
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	output := re.ReplaceAllString(input, "$1 $2")

	// Convert the first letter to uppercase
	output = strings.Title(output)

	return output
}
