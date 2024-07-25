package detailcontroller

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/thoriqaufar/liquipedia-valorant-api/config"
	"github.com/thoriqaufar/liquipedia-valorant-api/entity"
	"github.com/thoriqaufar/liquipedia-valorant-api/helper"
	"github.com/thoriqaufar/liquipedia-valorant-api/model"
	"net/http"
)

func TeamDetail(c echo.Context) error {
	var teams []entity.Team
	var response []model.TeamDetailsResponse

	err := config.DB.Model(&entity.Team{}).Preload("Player").Find(&teams).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&response, &teams)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, helper.H{"teams": response})
}

func PlayerDetail(c echo.Context) error {
	var players []entity.Player
	var response []model.PlayerDetailsResponse

	err := config.DB.Model(&entity.Player{}).Preload("Team").Find(&players).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&response, &players)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, helper.H{"players": response})
}
