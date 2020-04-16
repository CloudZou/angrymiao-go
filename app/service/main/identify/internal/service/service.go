package service

import (
	"angrymiao-go/app/service/main/identify/conf"
	"angrymiao-go/app/service/main/identify/internal/dao"
	"angrymiao-go/punk/cache"
)
// Service service.
type Service struct {
	d *dao.Dao
	cache *cache.Cache
}

// New new a service and return.
func New(c *conf.Config) (s *Service, cf func(), err error) {
	s = &Service{
		d: dao.New(c),
	}
	cf = s.Close
	return
}

// Close close the resource.
func (s *Service) Close() {
}
