//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func WireApp() *App {
	wire.Build(NewConfig, NewMysqlClient, NewRedisClient, NewApp)
	return &App{}
}
