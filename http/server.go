package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"signable-sdk/http/branding"
	"signable-sdk/http/contact"
	"signable-sdk/http/envelope"
	"signable-sdk/http/settings"
	"signable-sdk/http/team"
	"signable-sdk/http/template"
	"signable-sdk/http/user"
	"signable-sdk/http/webhook"
	"time"
)

var (
	HttpServer  *echo.Echo
	LogRequests = false
)

func Init() {
	HttpServer = echo.New()

	HttpServer.HideBanner = true

	addMiddleware()
	addRoutes()

	go func() {
		if err := HttpServer.Start(":8080"); err != nil && err != http.ErrServerClosed {
			HttpServer.Logger.Fatal("shutting down http server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpServer.Shutdown(ctx); err != nil {
		HttpServer.Logger.Fatal(err)
	}
}

func addMiddleware() {
	HttpServer.Use(middleware.Recover())
	HttpServer.Use(SDKHeaderMiddleware)
	HttpServer.Use(BearerAuthMiddleware)

	if LogRequests == true {
		HttpServer.Use(middleware.Logger())
	}
}

func addRoutes() {
	// prefix all routes with /v1
	g := HttpServer.Group("/v1")

	// contacts
	g.GET("/contacts", contact.List)
	g.POST("/contacts", contact.Create)
	g.GET("/contacts/:id", contact.Read)
	g.PUT("/contacts/:id", contact.Update)
	g.DELETE("/contacts/:id", contact.Delete)
	g.GET("/contacts/:id/envelopes", contact.EnvelopeList)

	// envelope
	g.GET("/envelopes", envelope.List)
	g.POST("/envelopes", envelope.Create)
	g.GET("/envelopes/:id", envelope.Read)
	g.PUT("/envelopes/:id/remind", envelope.Reminder)
	g.PUT("/envelopes/:id/cancel", envelope.Cancel)
	g.PUT("/envelopes/:id/expire", envelope.Expire)
	g.DELETE("/envelopes/:id", envelope.Delete)

	// template
	g.GET("/templates", template.List)
	g.GET("/templates/:id", template.Read)
	g.DELETE("/templates/:id", template.Delete)

	// user
	g.GET("/users", user.List)
	g.POST("/users", user.Create)
	g.GET("/users/:id", user.Read)
	g.PUT("/users/:id", user.Update)
	g.DELETE("/users/:id", user.Delete)

	// webhook
	g.GET("/webhooks", webhook.List)
	g.POST("/webhooks", webhook.Create)
	g.GET("/webhooks/:id", webhook.Read)
	g.PUT("/webhooks/:id", webhook.Update)
	g.DELETE("/webhooks/:id", webhook.Delete)

	// team
	g.GET("/teams", team.List)
	g.POST("/teams", team.Create)
	g.GET("/teams/:id", team.Read)
	g.PUT("/teams/:id", team.Update)
	g.DELETE("/teams/:id", team.Delete)

	// branding
	g.GET("/branding", branding.List)
	g.GET("/branding/emails", branding.ListEmails)
	g.PUT("/branding", branding.Update)
	g.PUT("/branding/emails/email-signed", branding.UpdateEmails)

	// settings
	g.GET("/settings", settings.List)
	g.PUT("/settings", settings.Update)

	// TODO: partner routes
}
