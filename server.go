package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"mohamadelabror.com/posapi/config"
	"mohamadelabror.com/posapi/delivery/api"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (a *appServer) initHandlers() {
	a.v1()
}

func (a *appServer) v1() {
	cashierApiGroup := a.routerEngine.Group("/cashiers")
	api.NewCashierApi(cashierApiGroup, a.cfg.UseCaseManager.ShowAllCashierUseCase(), a.cfg.UseCaseManager.GetCashierDetail(), a.cfg.UseCaseManager.CreateCashier(), a.cfg.UseCaseManager.UpdateCashier(), a.cfg.UseCaseManager.DeleteCashier())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.routerEngine.Run("localhost:3000")
	if err != nil {
		fmt.Println("test")
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig()
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
