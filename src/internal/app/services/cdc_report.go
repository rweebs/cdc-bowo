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

// NewCDCReportServices creates a new instance of CDCReportService. Note that you must call Close on the service to release resources
//
// Args:
//
//	dbSource: the source database to read from
//	dbDest: the destination database to write to. It is up to the caller to make sure the destination is closed before returning
//	rdb: the RDB to
//	config
func NewCDCReportServices(dbSource lib.Database, dbDest lib.Database, rdb lib.Cache, config config.Config) *CDCReportService {
	sqlFile, err := ioutil.ReadFile(config.SQLFile)
	// Read the SQL file and log error
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

// getTimeStampCutOff returns the timestamp to cut off when reading cdc report. This is used to determine when to stop reading
//
// Args:
//
//	s
func (s *CDCReportService) getTimeStampCutOff() (int64, error) {
	res, err := s.rdb.Cache.Get(context.Background(), "cdc-source-stop-timestamp").Result()
	// Returns 0 nil if there is an error.
	if err != nil {
		return 0, nil
	}
	result, err := strconv.ParseInt(res, 10, 64)
	return result, err
}

// GenerateReport generates CDC report. The time stamp cut off is set in config. TimeStampCutOff
//
// Args:
//
//	s
func (s *CDCReportService) GenerateReport() {
	timeStampCutOff, err := s.getTimeStampCutOff()
	// If err is not nil log. Panic.
	if err != nil {
		log.Panic(err)
	}
	log.Println("Generate Report")

	utils.GenerateReport(s.dbSource.Db, s.dbDest.Db, s.transformationList, timeStampCutOff)
}
