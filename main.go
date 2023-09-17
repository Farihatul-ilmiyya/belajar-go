package main

import (
	"echo-golang/config"
	handler_product "echo-golang/handler/handlerProduct"
	"echo-golang/injector/injectorRepository"
	"echo-golang/pkg/logs"
	"echo-golang/service/service_product"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	logging := logs.NewLogger()
	db, err := config.DB()
	if err != nil {
		return
	}
	rp := injectorRepository.NewInjectorRepository(db, logging)
	// repositoryProduct := repository_product.NewRepositoryProduct(db, logging)
	serviceProduct := service_product.NewServiceProduct(rp, logging)
	handlerProduct := handler_product.NewHandler(serviceProduct, logging)
	e := echo.New()

	// if err = db.AutoMigrate(&product.Product{}); err != nil {
	// 	logging.Error("failed auto migrate product ", err)
	// 	return
	// }
	e.POST("/product", handlerProduct.Create)
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
