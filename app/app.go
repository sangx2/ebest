package app

import (
	"fmt"
	"github.com/sangx2/ebest/api"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
	"sync"
)

type App struct {
	// servers
	as *api.Server
	es *EBestServer

	// config
	config *model.Config

	wg sync.WaitGroup
}

func NewApp(config *model.Config) *App {
	a := &App{
		config: config,
	}

	es := NewEBestServer(config)
	if es == nil {
		log.Error("NewEBestServer is nil")
		return nil
	}
	a.es = es

	as := api.NewServer(config.APISettings)
	if as == nil {
		log.Error("NewApiServer is nil")
		return nil
	}
	a.as = as

	return a
}

func (a *App) InitApp() error {
	if e := a.es.Init(); e != nil {
		return fmt.Errorf("InitApp: %w", e)
	}
	log.Info("eBest 서버 초기화 완료")

	if e := a.as.Init(a.es); e != nil {
		return fmt.Errorf("InitApp: %w", e)
	}
	log.Info("api 서버 초기화 완료")

	return nil
}

func (a *App) StartApp() error {
	if e := a.es.Start(); e != nil {
		return fmt.Errorf("StartApp error: %v", e)
	}
	log.Info("eBest 서버 시작")

	if e := a.as.Start(); e != nil {
		return fmt.Errorf("StratApp error: %w", e)
	}
	log.Info("api 서버 시작")

	return nil
}

func (a *App) ShutdownApp() {
	a.as.Shutdown()
	log.Info("api 서버 종료")

	a.es.Shutdown()
	log.Info("eBest 서버 종료")
}
