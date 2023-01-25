package main

import (
	"context"
	"dwd-api/app/routers"
	"dwd-api/global"
	"dwd-api/pkg/shutdown"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.NewHTTPServer()
	gin.SetMode(global.ServerSetting.RunMode)
}

func main() {

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//开始监听
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrAbortHandler {
			global.Logger.Fatal("http server startup errs:%v", err)
		}
	}()

	//优雅关机
	shutdown.NewHook().Close(
		//关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := s.Shutdown(ctx); err != nil {
				global.Logger.Fatal("server shutdown errs:%v", err)
			}
		},
		//关闭db
		func() {
			if global.DBEngine != nil {
				if err := global.DBEngine.Close(); err != nil {
					global.Logger.Error("mysql db close err", err)
				}
			}
		},
		//关闭redis
		func() {
			if global.RdbCliend != nil {
				if err := global.RdbCliend.Close(); err != nil {
					global.Logger.Error("redis db close err", err)
				}
			}
		},
	)
}
