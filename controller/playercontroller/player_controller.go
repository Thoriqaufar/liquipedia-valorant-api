package playercontroller

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/thoriqaufar/liquipedia-valorant-api/config"
	"github.com/thoriqaufar/liquipedia-valorant-api/entity"
	"github.com/thoriqaufar/liquipedia-valorant-api/helper"
	"github.com/thoriqaufar/liquipedia-valorant-api/model"
	"net/http"
	"strconv"
)

func FindAll(c echo.Context) error {
	var players []entity.Player
	var playerResponse []model.PlayerResponse

	config.DB.Find(&players)

	err := copier.Copy(&playerResponse, &players)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"players": playerResponse,
	})
}

func FindById(c echo.Context) error {
	var player entity.Player
	var playerResponse model.PlayerResponse

	id := c.Param("id")

	err := config.DB.First(&player, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&playerResponse, &player)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, helper.H{
		"status": http.StatusOK,
		"player": playerResponse,
	})
}

func Create(c echo.Context) error {
	var team entity.Team
	var player entity.Player
	var playerCreateRequest model.CreatePlayerRequest

	err := c.Bind(&playerCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	teamIDString := strconv.FormatUint(uint64(playerCreateRequest.TeamID), 10)

	err = config.DB.First(&team, playerCreateRequest.TeamID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": "There is no team_id = " + teamIDString})
	}

	v := validator.New()
	err = v.Struct(playerCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&player, playerCreateRequest)
	helper.PanicIfError(err)

	config.DB.Create(&player)

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"message": "Record Created",
	})
}

func Update(c echo.Context) error {
	var team entity.Team
	var player entity.Player
	var playerUpdateRequest model.UpdatePlayerRequest

	id := c.Param("id")

	err := c.Bind(&playerUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	teamIDString := strconv.FormatUint(uint64(playerUpdateRequest.TeamID), 10)

	err = config.DB.First(&team, playerUpdateRequest.TeamID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": "There is no team_id = " + teamIDString})
	}

	v := validator.New()
	err = v.Struct(playerUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&player, &playerUpdateRequest)
	helper.PanicIfError(err)

	tx := config.DB.Model(&player).Where("id = ?", id).Updates(&player)
	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"message": "Record Updated",
	})
}

func Delete(c echo.Context) error {
	var player entity.Player

	id := c.Param("id")

	tx := config.DB.Delete(&player, id)
	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": "Record Not Found"})
	}

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"message": "Record Deleted",
	})
}
