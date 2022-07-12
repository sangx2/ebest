package main

import (
	"fmt"
	"github.com/sangx2/ebest/app"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	configFilePath = "./config"
	configFileName = "config_trader"
	configFileType = "json"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// set config
	var config *model.Config

	viper.AddConfigPath(".")
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)

	if e := viper.ReadInConfig(); e != nil {
		fmt.Printf("read config error: %v", e)
	}

	if e := viper.Unmarshal(&config); e != nil {
		fmt.Printf("unmarshal config error: %v", e)
	}

	if e := config.IsValid(); e != nil {
		fmt.Printf("config is not valid: %v", e)
		return
	}

	// set logger
	log.InitGlobalLogger(log.NewLogger(config.TraderSettings.LogFileName, config.TraderSettings.LogLevel))
	log.Info("start trader")

	// App
	a := app.NewApp(config)
	if a == nil {
		log.Error("NewApp is nil")
		return
	}

	if e := a.InitApp(); e != nil {
		log.Error("InitApp failed: %v", log.Err(e))
		return
	}
	defer a.ShutdownApp()

	if e := a.StartApp(); e != nil {
		log.Error("StartApp failed: %v", log.Err(e))
		return
	}

	// set interrupt
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-interruptChan:
		log.Debug("get interrupt")
	}

	log.Info("shutdown trader")
}
