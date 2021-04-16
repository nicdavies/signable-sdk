package http

import (
	"github.com/labstack/echo/v4"
	"signable-sdk/http/branding"
	"signable-sdk/http/contact"
	"signable-sdk/http/envelope"
	"signable-sdk/http/settings"
	"signable-sdk/http/team"
	"signable-sdk/http/template"
	"signable-sdk/http/user"
	"signable-sdk/http/webhook"
)

func Init() {
	main()
}

func main() {
	e := echo.New()

	// TODO: add /v1/ prefix to all routes

	// contacts
	e.GET("/contacts", contact.List)
	e.POST("/contacts", contact.Create)
	e.GET("/contacts/:id", contact.Read)
	e.PUT("/contacts/:id", contact.Update)
	e.DELETE("/contacts/:id", contact.Delete)
	e.GET("/contacts/:id/envelopes", contact.EnvelopeList)

	// envelope
	e.GET("/envelopes", envelope.List)
	e.POST("/envelopes", envelope.Create)
	e.GET("/envelopes/:id", envelope.Read)
	e.PUT("/envelopes/:id/remind", envelope.Reminder)
	e.PUT("/envelopes/:id/cancel", envelope.Cancel)
	e.PUT("/envelopes/:id/expire", envelope.Expire)
	e.DELETE("/envelopes/:id", envelope.Delete)

	// template
	e.GET("/templates", template.List)
	e.GET("/templates/:id", template.Read)
	e.DELETE("/templates/:id", template.Delete)

	// user
	e.GET("/users", user.List)
	e.POST("/users", user.Create)
	e.GET("/users/:id", user.Read)
	e.PUT("/users/:id", user.Update)
	e.DELETE("/users/:id", user.Delete)

	// webhook
	e.GET("/webhooks", webhook.List)
	e.POST("/webhooks", webhook.Create)
	e.GET("/webhooks/:id", webhook.Read)
	e.PUT("/webhooks/:id", webhook.Update)
	e.DELETE("/webhooks/:id", webhook.Delete)

	// team
	e.GET("/teams", team.List)
	e.POST("/teams", team.Create)
	e.GET("/teams/:id", team.Read)
	e.PUT("/teams/:id", team.Update)
	e.DELETE("/teams/:id", team.Delete)

	// branding
	e.GET("/branding", branding.List)
	e.GET("/branding/emails", branding.ListEmails)
	e.PUT("/branding", branding.Update)
	e.PUT("/branding/emails/email-signed", branding.UpdateEmails)

	// settings
	e.GET("/settings", settings.List)
	e.PUT("/settings", settings.Update)

	// TODO: partner

	e.Logger.Fatal(e.Start(":8080"))
}
