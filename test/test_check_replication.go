package test

// import (
// 	"fmt"

// 	"github.com/rweebs/cdc-bowo/internal/app/config"
// 	"github.com/rweebs/cdc-bowo/internal/app/lib"
// 	"github.com/rweebs/cdc-bowo/internal/app/services"
// )

// func main() {
// 	config, err := config.LoadConfig("./")
// 	if err != nil {
// 		panic(err)
// 	}
// 	sourceDb := lib.NewDatabase(config.SourceConfig.Host, config.SourceConfig.Username, config.SourceConfig.Password, config.SourceConfig.Name, config.SourceConfig.Port)
// 	destDb := lib.NewDatabase(config.DestConfig.Host, config.DestConfig.Username, config.DestConfig.Password, config.DestConfig.Name, config.DestConfig.Port)
// 	cache := lib.NewCache()
// 	cdcSourceService := services.NewCDCSourceServices(sourceDb, destDb, cache, config)
// 	result, err := cdcSourceService.CheckReplicationCatchUp()
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		log.Println(result)
// 	}

// }
