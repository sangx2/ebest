package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/middleware"
	"github.com/sangx2/ebest/interfaces"
	"github.com/sangx2/ebest/model"
	"sync"
	"time"

	"github.com/labstack/echo"
	log "github.com/sangx2/golog"
)

type Server struct {
	echo *echo.Echo

	api *API

	settings *model.APISettings

	wg sync.WaitGroup
}

func NewServer(settings *model.APISettings) *Server {
	e := echo.New()

	a := Server{
		echo: e,

		api: NewAPI(e),

		settings: settings,
	}

	return &a
}

func (s *Server) Init(serverInterface interfaces.EBestServer) error {
	if s.echo == nil {
		return fmt.Errorf("Init: echo is nil")
	}

	if s.settings.KeyAuth.Enable {
		s.echo.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
			return key == s.settings.KeyAuth.Key, nil
		}))
	}

	s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tc := &Context{c}
			tc.SetServer(serverInterface)

			return next(tc)
		}
	})

	return nil
}

func (s *Server) SetData(key string, val interface{}) {
	s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(key, val)

			return next(c)
		}
	})
}

func (s *Server) Start() error {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if s.settings.TLS.Enable {
			if e := s.echo.StartTLS(":"+s.settings.Port, s.settings.TLS.CertPEMPath, s.settings.TLS.KeyPEMPath); e != nil {
				log.Error("StartTLS", log.Err(e))
			}
		} else {
			if e := s.echo.Start(":" + s.settings.Port); e != nil {
				log.Error("Start", log.Err(e))
			}
		}
	}()

	return nil
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if e := s.echo.Shutdown(ctx); e != nil {
		log.Info(e.Error())
	}

	s.wg.Wait()
}
