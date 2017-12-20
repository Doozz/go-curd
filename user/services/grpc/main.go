package main

import (
    "fmt"
    "log"
    "net"
    "user/config"
    "user/grpc/info"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "github.com/koding/multiconfig"
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

    lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Config.RpcConf.Host, config.Config.RpcConf.Port))
    if err != nil {
        grpclog.Fatalf("failed to listen: %v\n", err)
    }

    grpcServer := grpc.NewServer()
    info.RegisterInfoServer(grpcServer, new(info.UinfoServer))

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

}