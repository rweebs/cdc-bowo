package services

import (
	"context"
	"time"

	"github.com/rweebs/cdc-bowo/internal/app/lib"
)

type CDCMgmtService struct {
	rdb lib.Cache
}

func NewCDCMgmtService(rdb lib.Cache) *CDCMgmtService {
	return &CDCMgmtService{
		rdb: rdb,
	}
}

func (s *CDCMgmtService) StartBlueGreen() {
	s.rdb.Cache.Set(context.Background(), "cdc-source-expired", time.Now().UnixNano()/int64(time.Millisecond), 0)

}
