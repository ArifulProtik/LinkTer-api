package controller

import (
	"LinkTer-api/internel/utility"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) Signup(e echo.Context) error {
	if e.Request().Body != nil {
		newInput := utility.UserInput{}
		if err := e.Bind(&newInput); err != nil {
			return e.JSON(http.StatusUnprocessableEntity, ErrorResponse{
				Msg: "You Must Provide JSON",
			})
		}
		ferr := c.validate.Struct(newInput)
		if ferr != nil {
			return e.JSON(http.StatusUnprocessableEntity, ErrorResponse{
				Msg: ferr,
			})
		}
		newuser, err := c.service.SaveUser(&newInput)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, ErrorResponse{
				Msg: err.Error(),
			})
		}
		return e.JSON(http.StatusCreated, newuser)

	}
	return e.JSON(http.StatusUnprocessableEntity, ErrorResponse{
		Msg: "Body Required",
	})
}
