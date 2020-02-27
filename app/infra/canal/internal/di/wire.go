// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"angrymiao-go/app/infra/canal/internal/dao"
	"angrymiao-go/app/infra/canal/internal/server/http"
	"angrymiao-go/app/infra/canal/internal/service"

	"github.com/google/wire"
)

//go:generate punk t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, NewApp))
}
