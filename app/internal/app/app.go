package app

import (
	_ "Cinema/docs"
	"Cinema/internal/config"
	"Cinema/pkg/logging"
	"Cinema/pkg/metric"
	"context"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("router initial")
	router := httprouter.New()

	logger.Println("swagger doc initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	return App{
		cfg:    config,
		logger: logger,
		router: router,
	}, nil
}

func (a *App) Run() {
	a.startHttp()
}

func (a *App) startHttp() {
	a.logger.Info("start HTTP")
	var listener net.Listener

	if a.cfg.Listen.Type == config.ListenTypeSock {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			a.logger.Fatal(err)
		}
		socketPath := path.Join(appDir, a.cfg.Listen.SocketFile)
		a.logger.Infof("socket path: #{socketPath}")

		a.logger.Info("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	} else {
		a.logger.Infof("bind application to host: #{a.cfg.Listen.BindIP} and port: #{a.cfg.Listen.Port}")
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
		if err != nil {
			a.logger.Fatal(err)
		}

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

		handler := c.Handler(a.router)

		a.httpServer = &http.Server{
			Handler:      handler,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		a.logger.Println("application completely initialized and started")

		if err := a.httpServer.Serve(listener); err != nil {
			switch {
			case errors.Is(err, http.ErrServerClosed):
				a.logger.Warn("server shutdown")
			default:
				a.logger.Fatal(err)
			}
		}
		err = a.httpServer.Shutdown(context.Background())
		if err != nil {
			a.logger.Fatal(err)
		}
	}
}
