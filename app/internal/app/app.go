package app

import (
	_ "Cinema/docs"
	"Cinema/internal/config"
	"Cinema/pkg/common/errors"
	"Cinema/pkg/common/logging"
	psql "Cinema/pkg/postgresql"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg *config.Config

	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(ctx context.Context, config *config.Config) (App, error) {
	logging.L(ctx).Info("router initial")
	router := httprouter.New()

	logging.L(ctx).Info("swagger doc initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logging.WithFields(ctx,
		logging.StringField("username", config.PostgreSQL.Username),
		logging.StringField("password", "<REMOVED>"),
		logging.StringField("host", config.PostgreSQL.Host),
		logging.StringField("port", config.PostgreSQL.Port),
		logging.StringField("database", config.PostgreSQL.Database),
	).Info("PostgreSQL initializing")

	pgDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.PostgreSQL.Username,
		config.PostgreSQL.Password,
		config.PostgreSQL.Host,
		config.PostgreSQL.Port,
		config.PostgreSQL.Database,
	)

	pgClient, err := psql.NewClient(ctx, 5, 3*time.Second, pgDsn, false)
	if err != nil {
		return App{}, errors.Wrap(err, "psql.NewClient")

	}
	defer pgClient.Close()

	return App{
		cfg:    config,
		router: router,
	}, nil
}

func (app *App) Run(ctx context.Context) error {
	return app.startHttp(ctx)
}

func (app *App) startHttp(ctx context.Context) error {
	logger := logging.WithFields(ctx,
		logging.StringField("IP", app.cfg.HTTP.IP),
		logging.IntField("Port", app.cfg.HTTP.Port),
	)

	logger.Info("HTTP Server initializing")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", app.cfg.HTTP.IP, app.cfg.HTTP.Port))
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to create listener")
	}

	logger.With(
		logging.StringsField("AllowedMethods", app.cfg.HTTP.CORS.AllowedMethods),
		logging.StringsField("AllowedOrigins", app.cfg.HTTP.CORS.AllowedOrigins),
		logging.BoolField("AllowCredentials", app.cfg.HTTP.CORS.AllowCredentials),
		logging.StringsField("AllowedHeaders", app.cfg.HTTP.CORS.AllowedHeaders),
		logging.BoolField("OptionsPassthrough", app.cfg.HTTP.CORS.OptionsPassthrough),
		logging.StringsField("ExposedHeaders", app.cfg.HTTP.CORS.ExposedHeaders),
		logging.BoolField("Debug", app.cfg.HTTP.CORS.Debug),
	).Info("CORS initializing")

	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut,
			http.MethodOptions, http.MethodDelete},
		AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              false,
	})

	logging.L(ctx).Info("initial handler")
	handler := c.Handler(app.router)

	logging.L(ctx).Info("http server serve")
	app.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: app.cfg.HTTP.WriteTimeout,
		ReadTimeout:  app.cfg.HTTP.ReadTimeout,
	}

	logging.L(ctx).Info("serve")
	if err = app.httpServer.Serve(listener); err != nil {
		logging.L(ctx).Info("Aaaaaaaaaaaaaaaaaaaaaa")
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.With(logging.ErrorField(err)).Fatal("failed to start server")
		}
	}

	logging.L(ctx).Info("application completely initialized and started")
	err = app.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to shutdown server")
	}

	return err
}
