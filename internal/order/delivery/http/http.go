package http

import "github.com/labstack/echo/v4"

type OrderHandlers interface {
	CreateOrder() echo.HandlerFunc
	PayOrder() echo.HandlerFunc
	SubmitOrder() echo.HandlerFunc
	UpdateOrder() echo.HandlerFunc

	GetOrderByID() echo.HandlerFunc
	Search() echo.HandlerFunc
}
