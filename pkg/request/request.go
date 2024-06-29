package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"log"
)

type (
	ContextWrapper interface {
		Bind(data any) error
	}

	contextWrapper struct {
		context   echo.Context
		validator *validator.Validate
	}
)

func NewContextWrapper(ctx echo.Context) ContextWrapper {
	return &contextWrapper{
		context:   ctx,
		validator: validator.New(),
	}
}

func (c *contextWrapper) Bind(data any) error {
	if err := c.context.Bind(data); err != nil {
		log.Printf("failed to bind data: %v", err.Error())
		return err
	}

	if err := c.validator.Struct(data); err != nil {
		return err
	}
	return nil
}
