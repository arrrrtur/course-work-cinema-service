// app.go

package app

import (
	_ "Cinema/docs"
	"Cinema/internal/config"
	cinemarepository "Cinema/internal/domain/cinema/repository"
	cinemaservice "Cinema/internal/domain/cinema/service"
	cinemaHallrepository "Cinema/internal/domain/cinemaHall/repository"
	cinemaHallservice "Cinema/internal/domain/cinemaHall/service"
	movierepository "Cinema/internal/domain/movie/repository"
	movieservice "Cinema/internal/domain/movie/service"
	sessionrepository "Cinema/internal/domain/session/repository"
	sessionservice "Cinema/internal/domain/session/service"
	ticketrepository "Cinema/internal/domain/ticket/repository"
	ticketservice "Cinema/internal/domain/ticket/service"
	userrepository "Cinema/internal/domain/user/repository"
	userservice "Cinema/internal/domain/user/service"
	"Cinema/internal/handlers"
	"Cinema/pkg/common/core/closer"
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
	closer.AddN(pgClient)

	logging.L(ctx).Info("initial repositories")
	cinemaRepository := cinemarepository.NewCinemaRepository(pgClient)
	cinemaHallRepository := cinemaHallrepository.NewCinemaHallRepository(pgClient)
	movieRepository := movierepository.NewMovieRepository(pgClient)
	sessionRepository := sessionrepository.NewSessionRepository(pgClient)
	ticketRepository := ticketrepository.NewTicketRepository(pgClient)
	userRepository := userrepository.NewUserRepository(pgClient)

	logging.L(ctx).Info("initial services")
	cinemaService := cinemaservice.NewCinemaService(cinemaRepository)
	hallService := cinemaHallservice.NewCinemaHallService(cinemaHallRepository)
	movieService := movieservice.NewMovieService(movieRepository)
	sessionService := sessionservice.NewSessionService(sessionRepository)
	ticketService := ticketservice.NewTicketService(ticketRepository)
	userService := userservice.NewUserService(userRepository)

	logging.L(ctx).Info("initial handlers")
	cinemaHandler := handlers.NewCinemaHandler(cinemaService)
	movieHandler := handlers.NewMovieHandler(movieService)
	ticketHandler := handlers.NewTicketHandler(ticketService)
	userHandler := handlers.NewUserHandler(userService)
	sessionHandler := handlers.NewSessionHandler(sessionService)
	HallHandler := handlers.NewCinemaHallHandler(hallService)

	logging.L(ctx).Info("initial routes")
	router.POST("/api/cinemas", cinemaHandler.CreateCinema)
	router.GET("/api/cinemas", cinemaHandler.GetAllCinemas)
	router.GET("/api/cinemas/:id", cinemaHandler.GetCinemaByID)
	router.PUT("/api/cinemas", cinemaHandler.UpdateCinema)
	router.DELETE("/api/cinemas/:id", cinemaHandler.DeleteCinema)

	router.POST("/api/movies", movieHandler.CreateMovie)
	router.GET("/api/movies", movieHandler.GetAllMovies)
	router.GET("/api/movies/:id", movieHandler.GetMovieByID)
	router.PUT("/api/movies", movieHandler.UpdateMovie)
	router.DELETE("/api/movies/:id", movieHandler.DeleteMovie)

	router.POST("/api/tickets/buy", ticketHandler.CreateTicket)
	router.GET("/api/tickets/:id", ticketHandler.GetTicketByID)
	router.GET("/api/user/tickets/:userId", ticketHandler.GetTicketsByUserID)
	router.PUT("/api/tickets/:id", ticketHandler.UpdateTicket)
	router.DELETE("/api/tickets/:id", ticketHandler.DeleteTicket)

	router.POST("/api/sessions", sessionHandler.CreateSession)
	router.GET("/api/sessions/:id", sessionHandler.GetSessionByID)
	router.GET("/api/cinema-hall/sessions/:cinemaHallId", sessionHandler.GetSessionsByCinemaHallID)
	router.PUT("/api/sessions/:id", sessionHandler.UpdateSession)
	router.DELETE("/api/sessions/:id", sessionHandler.DeleteSession)

	router.POST("/api/users", userHandler.CreateUser)
	router.GET("/api/users", userHandler.GetAllUsers)
	router.GET("/api/users/:id", userHandler.GetUserByID)
	router.PUT("/api/users/:id", userHandler.UpdateUser)
	router.DELETE("/api/users/:id", userHandler.DeleteUser)

	router.POST("/api/cinema-halls", HallHandler.CreateCinemaHall)
	router.GET("/api/cinema-halls/cinema/:cinema_id", HallHandler.GetAllCinemaHalls)
	router.GET("/api/cinema-halls/:id", HallHandler.GetCinemaHallByID)
	router.PUT("/api/cinema-halls/:id", HallHandler.UpdateCinemaHall)
	router.DELETE("/api/cinema-halls/:id", HallHandler.DeleteCinemaHall)

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

	// TODO: fix
	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut,
			http.MethodOptions, http.MethodDelete},
		AllowedOrigins:     []string{"*"},
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
