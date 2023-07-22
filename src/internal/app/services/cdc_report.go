package services

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/rweebs/cdc-bowo/internal/app/config"
	"github.com/rweebs/cdc-bowo/internal/app/lib"
	"github.com/rweebs/cdc-bowo/internal/app/utils"
)

type CDCReportService struct {
	dbSource           lib.Database
	dbDest             lib.Database
	rdb                lib.Cache
	timeStampCutOff    int64
	transformationList map[string]config.DDLTransform
	config             config.Config
}

func NewCDCReportServices(dbSource lib.Database, dbDest lib.Database, rdb lib.Cache, config config.Config) *CDCReportService {
	sqlFile, err := ioutil.ReadFile(config.SQLFile)
	if err != nil {
		log.Panic(`Error reading SQL file`)
	}
	return &CDCReportService{
		dbSource:           dbSource,
		dbDest:             dbDest,
		rdb:                rdb,
		timeStampCutOff:    time.Now().UnixNano() / int64(time.Millisecond),
		config:             config,
		transformationList: utils.InitTransformListNew(string(sqlFile), config.DDLTransform),
	}
}
func (s *CDCReportService) getTimeStampCutOff() (int64, error) {
	res, err := s.rdb.Cache.Get(context.Background(), "cdc-source-stop-timestamp").Result()
	if err != nil {
		return 0, nil
	}
	result, err := strconv.ParseInt(res, 10, 64)
	return result, err
}

func (s *CDCReportService) GenerateReport() {
	timeStampCutOff, err := s.getTimeStampCutOff()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Generate Report")

	utils.GenerateReport(s.dbSource.Db, s.dbDest.Db, s.transformationList, timeStampCutOff)
}
