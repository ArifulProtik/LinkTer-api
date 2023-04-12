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

func (c *Controller) SignIn(e echo.Context) error {
	if e.Request().Body != nil {
		signininput := utility.SignInInput{}
		if err := e.Bind(&signininput); err != nil {
			return e.JSON(http.StatusUnprocessableEntity, ErrorResponse{
				Msg: "JSON is Required",
			})
		}
		verr := c.validate.Struct(signininput)
		if verr != nil {
			return e.JSON(http.StatusUnprocessableEntity, ErrorResponse{
				Msg: "Ghonta",
			})
		}
		user, err := c.service.GetUserByEmail(signininput.Email)
		if err != nil || user == nil {
			return e.JSON(http.StatusBadRequest, ErrorResponse{
				Msg: "Wrong Email or Password",
			})
		}
		err = utility.VerifyPass(user.Password, signininput.Password)
		if err != nil {
			return e.JSON(http.StatusBadRequest, ErrorResponse{
				Msg: "Wrong Email or Password",
			})
		}
		AcToken, _ := c.auth.GenAccesstoken(&user.ID)
		RFToken, _ := c.auth.GenRefreshtoken(&user.ID)
		oldsession, err := c.service.GetSesssionByUser(user.ID)
		if err == nil {
			err := c.service.DeleteSession(oldsession.ID)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, ErrorResponse{
					Msg: "Internal Server Error",
				})
			}
		}
		data := utility.SessionInput{
			UserID: user.ID,
			Token:  RFToken,
			IP:     e.RealIP(),
		}
		err = c.service.SaveSession(data)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, ErrorResponse{
				Msg: "Internal Server Error",
			})
		}
		cookie := c.auth.GenCookie(c.cfg.CookieName, RFToken, c.cfg.RfTokenEXP)
		e.SetCookie(cookie)
		return e.JSON(http.StatusAccepted, echo.Map{
			"status":        "Ok",
			"refresh_token": AcToken,
		})

	}
	return e.JSON(http.StatusUnprocessableEntity, ErrorResponse{
		Msg: "Body required",
	})
}
