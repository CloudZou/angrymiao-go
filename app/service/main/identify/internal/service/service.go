package service

import (
	"angrymiao-go/punk/cache"
	"context"
	"fmt"

	pb "angrymiao-go/app/service/main/identify/api"
	"angrymiao-go/app/service/main/identify/internal/dao"
	"angrymiao-go/punk/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
)
// Service service.
type Service struct {
	d *dao.Dao
	cache *cache.Cache
}

// New new a service and return.
func New(d *dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		d: d,
	}
	cf = s.Close
	return
}

// Close close the resource.
func (s *Service) Close() {
}
