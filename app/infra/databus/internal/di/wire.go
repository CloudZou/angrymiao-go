// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"angrymiao-go/app/infra/databus/internal/dao"
	"angrymiao-go/app/infra/databus/internal/service"
	"angrymiao-go/app/infra/databus/internal/server/grpc"
	"angrymiao-go/app/infra/databus/internal/server/http"

	"github.com/google/wire"
)

//go:generate punk t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
