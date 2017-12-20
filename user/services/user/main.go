package main

import (
    "fmt"
    "log"
    "net/http"
    "user/apis/controllers"
    "user/config"

    "github.com/julienschmidt/httprouter"
    "github.com/koding/multiconfig"
    "github.com/pilu/xrequestid"
    "github.com/urfave/negroni"
)

func main() {
    var err error
    m := multiconfig.New()
    config.Config = new(config.CmdConfig)
    err = m.Load(config.Config)
    if err != nil {
        log.Fatalf("Load configuration failed. Error: %s\n", err.Error())
    }
    m.MustLoad(config.Config)

    err = config.InitializeConn()
    if err != nil {
        log.Fatalf("config.InitialzeConn() failed. Error info: %s\n", err.Error())
    }
    defer func() {
        config.DBHandle.Close()
        config.RedisHandle.Close()
    }()

    router := httprouter.New()
    router.NotFound = http.FileServer(http.Dir("../../statuc"))
    router.HandleMethodNotAllowed = false
    router.GET("/user/login", admin.Login)
    router.GET("/user/reg", admin.Reg)
    router.POST("/user/dologin", admin.DoLogin)
    router.GET("/user/index", admin.Index)
    router.POST("/user/doupdata", admin.DoUpdata)
    router.POST("/user/doreg", admin.DoReg)

    n := negroni.New(negroni.NewRecovery())
    n.Use(xrequestid.New(16))
    n.UseHandler(router)
    n.Run(fmt.Sprintf("%s:%d", config.Config.ApiConf.Host, config.Config.ApiConf.Port))
}
