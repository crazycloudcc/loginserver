package main

import (
	"assist"
	"cherry/base"
	"cherry/dbproxy"
	"cherry/nethttp"
	"cherry/nettcp"
	"cherry/netwebsocket"
	"fmt"
	"handlers"
	"model"
	"os"
	"os/signal"
	"time"
)

func main() {
	base.SetBigEndian()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	base.SetLogLevel(base.LOG_LEVEL_DEBUG)

	conf := assist.LoadConfigFile("./config/dev.json")
	base.LogInfo(fmt.Sprintf("Start Server Info: ID=[%d], Name=[%v]", conf.AppConf.ID, conf.AppConf.Group))

	httpConf := conf.HTTPConf
	httpsConf := conf.HTTPSConf
	wsConf := conf.WSConf
	tcpConf := conf.TCPConf
	redisConfRemote := conf.RedisConfRemote
	redisConfLocal := conf.RedisConfLocal

	/************************************************************************/
	// database & dataproxy service module.
	/************************************************************************/
	if redisConfRemote.Flag == 1 {
		if !dbproxy.RedisConnectRemote(redisConfRemote) {
			os.Exit(1)
			return
		}

		model.InitFromRedis()
	}

	if redisConfLocal.Flag == 1 {
		if !dbproxy.RedisConnectLocal(redisConfLocal) {
			os.Exit(1)
			return
		}
	}

	/************************************************************************/
	// model service module.
	/************************************************************************/

	/************************************************************************/
	// network service module.
	/************************************************************************/

	if wsConf.Flag == 1 {
		wsServ := netwebsocket.NewService(wsConf)
		wsServ.Start()
	}

	if tcpConf.Flag == 1 {
		tcpServ := nettcp.NewService(tcpConf)

		// TODO RegHandler

		tcpServ.Start()
	}

	if httpConf.Flag == 1 {
		httpServ := nethttp.NewHTTP(httpConf)

		// httpServ.RegHandler("/", ATestsHandle)
		httpServ.RegHandler("/test", handlers.ATestsHandle)
		httpServ.RegHandler("/register_guest", handlers.RegisterGuestHandle)
		httpServ.RegHandler("/login", handlers.LoginHandle)

		httpServ.Start()
	}

	if httpsConf.Flag == 1 {
		httpsServ := nethttp.NewHTTPS(httpsConf)

		// httpsServ.RegHandler("/", ATestsHandle)
		// httpsServ.RegHandler("/test", handlers.ATestsHandle)

		httpsServ.Start()
	}

	base.LogInfo("Server Start Time:", time.Now(), time.Now().Unix())
	base.LogInfo("Server Start Done.")

	/************************************************************************/
	// service stop.
	/************************************************************************/
	s := <-interrupt

	base.LogInfo("Server Stop With Signal: ", s)
	base.LogInfo("Server Stop Done.")
}
