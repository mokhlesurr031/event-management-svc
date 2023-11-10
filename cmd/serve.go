package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/cobra"

	"github.com/mokhlesurr031/event-management-svc/internal/connection"

	_reservationHttp "github.com/mokhlesurr031/event-management-svc/reservations/delivery/http"
	_reservationRepository "github.com/mokhlesurr031/event-management-svc/reservations/repository"
	_reservationUseCase "github.com/mokhlesurr031/event-management-svc/reservations/usecase"

	_eventHttp "github.com/mokhlesurr031/event-management-svc/events/delivery/http"
	_eventRepository "github.com/mokhlesurr031/event-management-svc/events/repository"
	_eventUseCase "github.com/mokhlesurr031/event-management-svc/events/usecase"

	_workshopHttp "github.com/mokhlesurr031/event-management-svc/workshops/delivery/http"
	_workshopRepository "github.com/mokhlesurr031/event-management-svc/workshops/repository"
	_workshopUseCase "github.com/mokhlesurr031/event-management-svc/workshops/usecase"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting Server...",
	Long:  `Starting Server...`,
	Run:   server,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func server(cmd *cobra.Command, args []string) {
	log.Println("Connecting database")
	if err := connection.ConnectDB(); err != nil {
		log.Println("database connection failed", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	srv := buildHTTP(cmd, args)
	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}(srv)
	<-stop
}

func buildHTTP(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Println(err)
		}
	})

	db := connection.DefaultDB()
	_ = db

	reservationRepo := _reservationRepository.New(db)
	reservationUsecase := _reservationUseCase.New(reservationRepo)
	_reservationHttp.NewHttpHandler(r, reservationUsecase)

	eventRepo := _eventRepository.New(db)
	eventUsecase := _eventUseCase.New(eventRepo)
	_eventHttp.NewHttpHandler(r, eventUsecase)

	workshopRepo := _workshopRepository.New(db)
	workshopUsecase := _workshopUseCase.New(workshopRepo)
	_workshopHttp.NewHttpHandler(r, workshopUsecase)

	return &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8081),
		Handler: r,
	}
}
