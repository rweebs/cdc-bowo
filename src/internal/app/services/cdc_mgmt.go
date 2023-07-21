package services

import (
	"context"
	"time"

	"github.com/rweebs/cdc-bowo/src/internal/app/lib"
)

type CDCMgmtService struct {
	rdb lib.Cache
}

func NewCDCMgmtService(rdb lib.Cache) *CDCMgmtService {
	return &CDCMgmtService{
		rdb: rdb,
	}
}

func (s *CDCMgmtService) StartSync() {
	s.rdb.Cache.Set(context.Background(), "stop-replication-ready", 1, 0)
}

func (s *CDCMgmtService) StartBlueGreen() {
	s.rdb.Cache.Set(context.Background(), "cdc-source-stop-timestamp", time.Now().UnixNano()/int64(time.Millisecond), 0)

}
