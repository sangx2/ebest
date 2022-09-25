package model

import (
	"errors"
	"strconv"
)

// AppSettings :
type AppSettings struct {
	ID         string
	Passwd     string
	CertPasswd string
	Server     string
	ResPath    string
	QueueSize  int
}

func (as *AppSettings) isValid() error {
	if len(as.ID) == 0 {
		return errors.New("Config.AppSettings.ID length is 0")
	}
	if len(as.Passwd) == 0 {
		return errors.New("Config.AppSettings.Passwd is 0")
	}
	if len(as.CertPasswd) == 0 {
		return errors.New("Config.AppSettings.CertPasswd is 0")
	}
	if len(as.Server) == 0 {
		return errors.New("Config.AppSettings.Server is 0")
	}
	if len(as.ResPath) == 0 {
		return errors.New("Config.AppSettings.ResPath is 0")
	}
	if as.QueueSize == 0 {
		return errors.New("Config.AppSettings.QueueSize is 0")
	}
	return nil
}

type AccountSettings struct {
	Accounts map[string]string
}

func (as *AccountSettings) isValid() error {
	return nil
}

// APISettings :
type APISettings struct {
	KeyAuth struct {
		Enable bool
		Key    string
	}
	TLS struct {
		Enable      bool
		CertPEMPath string
		KeyPEMPath  string
	}
	Port string
}

func (as *APISettings) isValid() error {
	if as.KeyAuth.Enable {
		if len(as.KeyAuth.Key) == 0 {
			return errors.New("Config.APISetting.KeyAuth.Key is nil")
		}
	}
	if as.TLS.Enable {
		if len(as.TLS.CertPEMPath) == 0 {
			return errors.New("Config.APISetting.TLS.CertPEMPath is nil")
		}
		if len(as.TLS.KeyPEMPath) == 0 {
			return errors.New("Config.APISetting.TLS.KeyPEMPath is nil")
		}
	}
	if _, e := strconv.Atoi(as.Port); e != nil {
		return e
	}

	return nil
}

type ManagerSettings struct {
	DelayStartTrader string
	LogLevel         string
	LogFileName      string
}

func (ms *ManagerSettings) isValid() error {
	if len(ms.DelayStartTrader) == 0 {
		return errors.New("ManagerSettings.DelayStartTrader is nil")
	}
	if len(ms.LogLevel) == 0 {
		return errors.New("ManagerSettings.LogLevel is nil")
	}
	if len(ms.LogFileName) == 0 {
		return errors.New("ManagerSettings.LogFilename is nil")
	}
	return nil
}

type TraderSettings struct {
	LogLevel    string
	LogFileName string
}

func (ts *TraderSettings) isValid() error {
	if len(ts.LogLevel) == 0 {
		return errors.New("TraderSettings.LogLevel is nil")
	}
	if len(ts.LogFileName) == 0 {
		return errors.New("TraderSettings.LogFileName is nil")
	}
	return nil
}

// Config :
type Config struct {
	AppSettings     *AppSettings
	AccountSettings *AccountSettings
	APISettings     *APISettings
	ManagerSettings *ManagerSettings
	TraderSettings  *TraderSettings
}

func (c *Config) IsValid() error {
	if c.AppSettings != nil {
		if err := c.AppSettings.isValid(); err != nil {
			return err
		}
	}
	if c.AccountSettings != nil {
		if err := c.AccountSettings.isValid(); err != nil {
			return err
		}
	}
	if c.APISettings != nil {
		if err := c.APISettings.isValid(); err != nil {
			return err
		}
	}
	if c.ManagerSettings != nil {
		if err := c.ManagerSettings.isValid(); err != nil {
			return err
		}
	}
	if c.TraderSettings != nil {
		if err := c.TraderSettings.isValid(); err != nil {
			return err
		}
	}

	return nil
}
