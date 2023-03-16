package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/labstack/echo"
)

func FetchAllPegawai(c echo.Context) error {

	result, err := models.FetchAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.StorePegawai(nama, alamat, telepon, username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")
	username := c.FormValue("username")
	password := c.FormValue("password")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePegawai(nama, alamat, telepon, username, password, conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeletePegawai(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
