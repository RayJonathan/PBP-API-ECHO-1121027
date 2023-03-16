package routes

import (
	"net/http"

	"example.com/m/controllers"
	"example.com/m/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil menggunakan Echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAutenticated)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/logout", controllers.Logout)
	e.POST("/users", controllers.StoreUsers)
	return e
}
