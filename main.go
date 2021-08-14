package main

import (
	"filesystem/config"
	"filesystem/dig"
	_ "filesystem/docs"
	router "filesystem/routers"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// @title 文件系统 api
// @description filesystem api
func main() {

	dig.Container()
	addr := config.GetInstance().GetString("application.host")
	if err := http.ListenAndServe(addr, router.GetInstance()); err != nil {
		logrus.Errorf("Failed to ListenAndServer at %v, err = %v", addr, err)
		os.Exit(1)
	}
}
