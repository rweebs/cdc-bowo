package services

import (
	"context"
	"time"

	"github.com/rweebs/cdc-bowo/internal/app/lib"
)

type CDCMgmtService struct {
	rdb lib.Cache
}

// NewCDCMgmtService creates a new instance of the CDC management service. Note that you must call Close on the returned service to free resources
//
// Args:
//
//	rdb: Cache to use for
func NewCDCMgmtService(rdb lib.Cache) *CDCMgmtService {
	return &CDCMgmtService{
		rdb: rdb,
	}
}

// StartSync tells cdcmgmt to start syncing. This is called when we have a connection to the master and it s time to send a request to the master.
//
// Args:
//
//	s: CDCMgmtService for this instance of
func (s *CDCMgmtService) StartSync() {
	s.rdb.Cache.Set(context.Background(), "stop-replication-ready", 1, 0)
}

// StartBlueGreen starts the cdc management service. This is a no op in order to avoid race conditions when we have multiple services running at the same time
//
// Args:
//
//	s: the service to start
func (s *CDCMgmtService) StartBlueGreen() {
	s.rdb.Cache.Set(context.Background(), "cdc-source-stop-timestamp", time.Now().UnixNano()/int64(time.Millisecond), 0)

}
