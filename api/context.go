package api

import (
	"github.com/labstack/echo"
	"github.com/sangx2/ebest/interfaces"
)

type Context struct {
	echo.Context
}

func (c *Context) SetServer(serverInterface interfaces.EBestServer) {
	c.Set("server", serverInterface)
}

func (c *Context) GetServer() interfaces.EBestServer {
	server, _ := c.Get("server").(interfaces.EBestServer)

	return server
}
