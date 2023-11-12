package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kitabisa/kibitalk/api"
	"github.com/kitabisa/kibitalk/client/campaign"
	"github.com/kitabisa/kibitalk/client/payment"
	"github.com/kitabisa/kibitalk/config"
	"github.com/kitabisa/kibitalk/config/broker/rabbitmq"
	"github.com/kitabisa/kibitalk/config/cache"
	"github.com/kitabisa/kibitalk/config/database"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var RestApiCmd = &cobra.Command{
	Use:   "api",
	Short: "kibitalk api",
	Run: func(cmd *cobra.Command, args []string) {
		config.NewAppConfig()
		database.InitMySQL()
		cache.InitCache()
		payment.InitPaymentClient()
		campaign.InitCampaignClient()
		rabbitmq.NewAMQPClient()

		router := chi.NewRouter()
		router.Use(
			middleware.RedirectSlashes,
			middleware.Recoverer,
			middleware.Logger, //middleware to recover from panics
		)

		//Sets context for all requests
		router.Use(middleware.Timeout(30 * time.Second))

		// Add routes
		api.ApplyRoutes(router)

		fmt.Println("List of registered endpoints")
		walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			route = strings.Replace(route, "/*/", "/", -1)
			fmt.Printf("%s %s\n", method, route)
			return nil
		}

		if err := chi.Walk(router, walkFunc); err != nil {
			fmt.Printf("Logging err: %s\n", err.Error())
		}

		// Create a context that listens for termination signals.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Create a WaitGroup to keep track of active connections.
		var wg sync.WaitGroup

		// Start the HTTP server.
		server := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.AppCfg.App.Host, config.AppCfg.App.Port),
			Handler: router,
		}

		zlog.Info().Msgf("Starting server on %s", server.Addr)

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			zlog.Err(err).Msgf("[API] Fail to start listen and server: %v", err)
		}

		// Create a channel to listen for OS signals (e.g., SIGINT, SIGTERM).
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		// Block until a signal is received.
		sig := <-sigCh
		fmt.Printf("Received signal: %v\n", sig)

		// Cancel the context to stop accepting new requests.
		cancel()

		// Wait for existing requests to finish.
		wg.Wait()

		// Shutdown the HTTP server gracefully.
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Error during server shutdown: %v\n", err)
		} else {
			fmt.Println("Server has shut down gracefully.")
		}

	},
}
