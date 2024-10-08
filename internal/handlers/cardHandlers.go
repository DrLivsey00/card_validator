package handlers

import (
	"errors"
	"github.com/DrLivsey00/card_checker/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Code    string `query:"code"`
	Message string `query:"message"`
}

type Response struct {
	Valid bool           `query:"valid"`
	Error *ErrorResponse `query:"error"`
}

func (h *Handlers) IsValidCard(c echo.Context) error {
	var card models.Card
	err := c.Bind(&card)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Valid: false,
			Error: &ErrorResponse{
				Code:    http.StatusText(http.StatusBadRequest),
				Message: errors.New("invalid params").Error(),
			},
		})
	}
	valid, err, code := h.service.IsValid(card.Number, card.ExpirationMonth, card.ExpirationYear)
	if err != nil {
		return c.JSON(http.StatusOK, &Response{
			Valid: false,
			Error: &ErrorResponse{
				Code:    code,
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Valid: valid,
		Error: nil,
	})

}
