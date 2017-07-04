package main

import (
	"fmt"
	"net/http"
	"os"

	"cmdr/cmdr"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Version is app version
const Version = "0.1"

// CommandList is command sent to service
type CommandList struct {
	// Commands is a list of command
	Commands []string `json:"commands"`
}

func main() {
	cnf := cmdr.NewConfig()
	conf, err := cnf.LoadConfig("./cmdr.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var e = echo.New()
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:API-Key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == conf.APIToken, nil
		},
	}))
	e.HideBanner = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello. This is CMD Service.")
	})
	e.POST("/dispatch", func(c echo.Context) error {
		clist := new(CommandList)
		if err := c.Bind(clist); err != nil {
			return err
		}

		var output []string

		d := cmdr.NewDispatcher(conf.Commands)
		for _, v := range clist.Commands {
			output = append(output, d.Do(v))
		}

		return c.JSON(http.StatusCreated, output)
	})

	var addr = fmt.Sprintf("%s:%s", conf.IP, conf.Port)

	fmt.Printf("CMDR v%s is serving on %s\n", Version, addr)
	e.Logger.Fatal(e.Start(addr))
}
