package handler_product

import (
	"echo-golang/model/product"
	"echo-golang/service/service_product"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type HandlerProduct interface {
	Create(c echo.Context) error
}
type handler_product struct {
	serviceProduct service_product.ServiceProduct
	logging        *logrus.Logger
}

func NewHandler(service service_product.ServiceProduct, logging *logrus.Logger) *handler_product {
	return &handler_product{serviceProduct: service, logging: logging}
}

func (h handler_product) Create(c echo.Context) error {
	time.Sleep(5 * time.Second)
	h.logging.Info("Frontend memanggil routing")

	time.Sleep(5 * time.Second)
	h.logging.Info("binding data json")

	// req := new(product.RequestProduct)
	req := product.RequestProduct{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ada yang salah nih")
	}

	time.Sleep(5 * time.Second)
	h.logging.Info("validasi input data")
	errs := req.Validate()
	if len(errs) > 0 {
		return c.JSON(http.StatusBadRequest, errs)
	}

	time.Sleep(5 * time.Second)
	h.logging.Info("kirim data ke service")
	response, status, err := h.serviceProduct.Create(req)
	if err != nil {
		return c.JSON(status, err)
	}
	time.Sleep(5 * time.Second)
	h.logging.Info("kembalikan data ke frontend")
	return c.JSON(status, response)
}
