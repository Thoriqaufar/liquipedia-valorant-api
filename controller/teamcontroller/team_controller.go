package teamcontroller

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/thoriqaufar/liquipedia-valorant-api/config"
	"github.com/thoriqaufar/liquipedia-valorant-api/entity"
	"github.com/thoriqaufar/liquipedia-valorant-api/helper"
	"github.com/thoriqaufar/liquipedia-valorant-api/model"
	"net/http"
)

func FindAll(c echo.Context) error {
	var teams []entity.Team
	var teamResponse []model.TeamResponse

	config.DB.Find(&teams)

	err := copier.Copy(&teamResponse, &teams)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, helper.H{
		"status": http.StatusOK,
		"teams":  teamResponse,
	})
}

func FindById(c echo.Context) error {
	var team entity.Team
	var teamResponse model.TeamResponse

	id := c.Param("id")

	err := config.DB.First(&team, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&teamResponse, &team)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, helper.H{
		"status": http.StatusOK,
		"team":   teamResponse,
	})
}

func Create(c echo.Context) error {
	var team entity.Team
	var teamCreateRequest model.CreateTeamRequest

	err := c.Bind(&teamCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	v := validator.New()
	err = v.Struct(teamCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&team, &teamCreateRequest)
	helper.PanicIfError(err)

	config.DB.Create(&team)

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"message": "Record Created",
	})
}

func Update(c echo.Context) error {
	var team entity.Team
	var teamUpdateRequest model.UpdateTeamRequest

	id := c.Param("id")

	err := c.Bind(&teamUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	v := validator.New()
	err = v.Struct(teamUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.H{"message": err.Error()})
	}

	err = copier.Copy(&team, &teamUpdateRequest)
	helper.PanicIfError(err)

	tx := config.DB.Model(&team).Where("id = ?", id).Updates(&team)
	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"message": "Record Updated",
	})
}

func Delete(c echo.Context) error {
	var team entity.Team

	id := c.Param("id")

	tx := config.DB.Delete(&team, id)
	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, helper.H{"message": "Record Not Found"})
	}

	return c.JSON(http.StatusOK, helper.H{
		"status":  http.StatusOK,
		"message": "Record Deleted",
	})
}
