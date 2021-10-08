package controller

import (
	"app/model"
	"app/service"
)

type GreetController struct {
	Service service.GreetService
	// Ctx iris.Context
}

func (c *GreetController) Get(req model.Request) (model.Response, error) {
	message, err := c.Service.Say(req.Name)
	if err != nil {
		return model.Response{}, err
	}

	return model.Response{Message: message}, nil
}
