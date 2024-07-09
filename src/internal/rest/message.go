package rest

import (
	"github.com/gofiber/fiber/v2"
	"main/domains/message"
	domainMessage "main/domains/message"
	"main/pkg/utils"
	"main/pkg/whatsapp"
)

type Message struct {
	Service domainMessage.IMessageService
}

func InitRestMessage(app *fiber.App, service domainMessage.IMessageService) Message {
	rest := Message{Service: service}
	app.Post("/message/:message_id/reaction", rest.ReactMessage)
	app.Post("/message/:message_id/revoke", rest.RevokeMessage)
	app.Post("/message/:message_id/delete", rest.DeleteMessage)
	app.Post("/message/:message_id/update", rest.UpdateMessage)
	return rest
}

func (controller *Message) RevokeMessage(c *fiber.Ctx) error {
	var request domainMessage.RevokeRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	request.MessageID = c.Params("message_id")
	whatsapp.SanitizePhone(&request.Phone)

	response, err := controller.Service.RevokeMessage(c.UserContext(), request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Status:  200,
		Code:    "SUCCESS",
		Message: response.Status,
		Results: response,
	})
}

func (controller *Message) DeleteMessage(c *fiber.Ctx) error {
	var request domainMessage.DeleteRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	request.MessageID = c.Params("message_id")
	whatsapp.SanitizePhone(&request.Phone)

	err = controller.Service.DeleteMessage(c.UserContext(), request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Status:  200,
		Code:    "SUCCESS",
		Message: "Message deleted successfully",
		Results: nil,
	})
}

func (controller *Message) UpdateMessage(c *fiber.Ctx) error {
	var request domainMessage.UpdateMessageRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	request.MessageID = c.Params("message_id")
	whatsapp.SanitizePhone(&request.Phone)

	response, err := controller.Service.UpdateMessage(c.UserContext(), request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Status:  200,
		Code:    "SUCCESS",
		Message: response.Status,
		Results: response,
	})
}

func (controller *Message) ReactMessage(c *fiber.Ctx) error {
	var request message.ReactionRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	request.MessageID = c.Params("message_id")
	whatsapp.SanitizePhone(&request.Phone)

	response, err := controller.Service.ReactMessage(c.UserContext(), request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Status:  200,
		Code:    "SUCCESS",
		Message: response.Status,
		Results: response,
	})
}
