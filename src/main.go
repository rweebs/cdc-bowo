package main

import (
	"github.com/rweebs/cdc-bowo/internal/app/services"
)

func main() {

	cdcSourceService := services.CDCSourceServices{}
	cdcSourceService.ExecuteDDLChange()
}
