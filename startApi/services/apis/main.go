package main

import (
    "fmt"
    "log"
    //"net/http"
    "startApi/apis/controllers"
    "startApi/config"

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

    router.HandleMethodNotAllowed = false
    router.GET("/api/query", controllers.Query)
    router.GET("/api/queryrow", controllers.QueryRow)
    router.POST("/api/create", controllers.Create)
    router.POST("/api/update", controllers.Update)
    router.POST("/api/delete", controllers.Delete)
    router.POST("/api/business", controllers.Business)


    n := negroni.New(negroni.NewRecovery())
    n.Use(xrequestid.New(16))
    n.UseHandler(router)
    n.Run(fmt.Sprintf("%s:%d", config.Config.ApiConf.Host, config.Config.ApiConf.Port))
}
