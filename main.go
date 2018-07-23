package main

import (
	"flag"

	"github.com/kooksee/srelay/config"
	"github.com/kooksee/srelay/server"
)

const Version = "1.0"

func main() {

	cfg := config.GetCfg()
	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "debug mode")
	flag.StringVar(&cfg.Host, "host", cfg.Host, "app host")
	flag.Int64Var(&cfg.Port, "kcp", cfg.Port, "app port")
	flag.Parse()

	cfg.InitLog()

	server.Init()

	us := &server.UdpServer{}
	if err := us.Listen(cfg.Port); err != nil {
		panic(err.Error())
	}

	ts := &server.TcpServer{}
	if err := ts.Listen(cfg.Port); err != nil {
		panic(err.Error())
	}

	select {}
}
