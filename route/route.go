package route

import (
	"github.com/labstack/echo/v4"
	"github.com/thoriqaufar/liquipedia-valorant-api/controller/detailcontroller"
	"github.com/thoriqaufar/liquipedia-valorant-api/controller/playercontroller"
	"github.com/thoriqaufar/liquipedia-valorant-api/controller/teamcontroller"
	"github.com/thoriqaufar/liquipedia-valorant-api/helper"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	// Index
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.H{"message": "Inspired by https://liquipedia.net/valorant"})
	})

	// Teams
	e.GET("/api/teams", teamcontroller.FindAll)
	e.GET("/api/team/:id", teamcontroller.FindById)
	e.POST("/api/team", teamcontroller.Create)
	e.PUT("/api/team/:id", teamcontroller.Update)
	e.DELETE("/api/team/:id", teamcontroller.Delete)

	// Players
	e.GET("/api/players", playercontroller.FindAll)
	e.GET("/api/player/:id", playercontroller.FindById)
	e.POST("/api/player", playercontroller.Create)
	e.PUT("/api/player/:id", playercontroller.Update)
	e.DELETE("/api/player/:id", playercontroller.Delete)

	// Show Details
	e.GET("/api/team/details", detailcontroller.TeamDetail)
	e.GET("/api/player/details", detailcontroller.PlayerDetail)

	return e
}
