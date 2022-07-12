package main

import (
	"fmt"
	"github.com/sangx2/ebest/model"
	"github.com/sangx2/go-servers/scheduling"
	log "github.com/sangx2/golog"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	ProcessNameTrader = "trader.exe"

	ConfigFilePath = "./config"
	ConfigFileName = "config_manager"
	ConfigFileType = "json"
)

func startTrader() {
	if e := exec.Command("cmd.exe", "/C", "start", ProcessNameTrader).Run(); e != nil {
		log.Error(ProcessNameTrader+" 프로세스 시작 에러", log.Err(e))
		return
	}
	fmt.Printf("%s: trader.exe 시작\n", time.Now().Format("2006-01-02T15:04:05"))
}

func shutdownTrader() {
	if e := exec.Command("taskkill", "/IM", ProcessNameTrader).Run(); e != nil {
		if strings.Contains(e.Error(), "exit status 128") {
			log.Warn(ProcessNameTrader + " 프로세스가 실행중이 아님")
		} else {
			log.Error(ProcessNameTrader+" 프로세스 종료 에러", log.Err(e))
		}
		return
	}
	fmt.Printf("%s: trader.exe 종료\n", time.Now().Format("2006-01-02T15:04:05"))
}

func main() {
	// config 설정
	var config *model.Config

	viper.AddConfigPath(".")
	viper.AddConfigPath(ConfigFilePath)
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)

	if e := viper.ReadInConfig(); e != nil {
		fmt.Printf("config 파일 읽기 에러: %v", e)
		return
	}

	if e := viper.Unmarshal(&config); e != nil {
		fmt.Printf("config 파일 unmarshal 에러: %v", e)
	}

	if e := config.IsValid(); e != nil {
		fmt.Printf("config 파일이 유효하지 않음: %v", e)
		return
	}

	// logger 설정
	log.InitGlobalLogger(log.NewLogger(config.ManagerSettings.LogFileName, config.ManagerSettings.LogLevel))
	log.Info("manager 시작")

	// trader 프로세스 종료
	shutdownTrader()

	// trader 프로세스 시작
	duration, e := time.ParseDuration(config.ManagerSettings.DelayStartTrader)
	if e != nil {
		fmt.Printf("DelayStartTrader가 유효하지 않음: %s", config.ManagerSettings.DelayStartTrader)
		return
	}
	fmt.Printf("Start "+ProcessNameTrader+": delay %s...\n", config.ManagerSettings.DelayStartTrader)
	time.Sleep(duration)
	startTrader()

	// scheduler 설정
	ss := scheduling.NewServer()
	if e := ss.AddSchedulerWithFunc("shutdown trader",
		scheduling.NewScheduler(7, 00, 0,
			[]time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}),
		shutdownTrader); e != nil {
		log.Error("AddSchedulerWithFunc", log.Err(e))
	}

	if e := ss.AddSchedulerWithFunc("start trader",
		scheduling.NewScheduler(7, 30, 0,
			[]time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}),
		startTrader); e != nil {
		log.Error("AddSchedulerWithFunc", log.Err(e))
	}

	// scheduler 시작
	ss.Start()
	defer ss.Shutdown()

	// interrupt 설정
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-interruptChan:
		log.Debug("interrupt 수신")
	}

	log.Info("manager 종료")
}
