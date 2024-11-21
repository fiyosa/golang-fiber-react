package controller

import (
	"go-fiber-react/app/event"
	"go-fiber-react/app/helper"
	"go-fiber-react/app/job"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

var Test testController

type testController struct{}

func (*testController) Job(c *fiber.Ctx) error {
	go job.Test1Job("Step 1")

	return helper.Res.SendSuccess(c, lang.L.Convert(lang.L.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": "Job"}))
}

func (*testController) Event(c *fiber.Ctx) error {
	event.TestEvent("Step 1")

	return helper.Res.SendSuccess(c, lang.L.Convert(lang.L.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": "Event"}))
}
