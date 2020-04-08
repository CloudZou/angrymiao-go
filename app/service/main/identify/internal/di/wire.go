// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"angrymiao-go/app/service/main/identify/internal/dao"
	"angrymiao-go/app/service/main/identify/internal/service"
	"angrymiao-go/app/service/main/identify/internal/server/grpc"
	"angrymiao-go/app/service/main/identify/internal/server/http"

	"github.com/google/wire"
)

//go:generate punk t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
