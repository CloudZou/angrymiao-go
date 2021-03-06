// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"angrymiao-go/app/infra/databus/internal/dao"
	"angrymiao-go/app/infra/databus/internal/server/http"
	"angrymiao-go/app/infra/databus/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	redisClient, err := dao.NewRedisClient()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := dao.NewDB()
	if err != nil {
		return nil, nil, err
	}
	daoDao, cleanup2, err := dao.New(redisClient, db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serviceService := service.New(daoDao)
	engine, err := http.New(serviceService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup3, err := NewApp(serviceService, engine)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
