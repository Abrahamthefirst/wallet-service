package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/config"
	"github.com/Abrahamthefirst/finecore-practice/internal/db/repository"
	endpoints "github.com/Abrahamthefirst/finecore-practice/internal/http"
	"github.com/Abrahamthefirst/finecore-practice/internal/http/middleware"
	authservice "github.com/Abrahamthefirst/finecore-practice/internal/service/auth-service"
	walletservice "github.com/Abrahamthefirst/finecore-practice/internal/service/wallet-service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"gorm.io/gorm"
)

type application struct {
	router *gin.Engine
	db     *gorm.DB
	cfg    *config.Config
	logger *slog.Logger
}

func NewApp(db *gorm.DB, cfg *config.Config, logger *slog.Logger) *application {
	router := gin.Default()

	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	}
	return &application{
		router: router,
		db:     db,
		cfg:    cfg,
		logger: logger,
	}

}

func (a *application) Bootstrap() {

	slog.SetDefault(a.logger)

	a.startServer()

}

func (a *application) startServer() {

	a.router.GET("/openapi.json", func(c *gin.Context) {
		instance := swag.GetSwagger("swagger")
		if instance == nil {
			c.JSON(500, gin.H{"error": "Swagger instance not found. Ensure the docs package is imported."})
			return
		}
		c.Data(200, "application/json", []byte(instance.ReadDoc()))
	})

	a.router.GET("/api/docs", func(c *gin.Context) {
		c.Data(200, "text/html", []byte(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>Insight API Docs</title>
            <meta charset="utf-8"/>
            <meta name="viewport" content="width=device-width, initial-scale=1">
        </head>
        <body>
            <script
                id="api-reference"
                data-url="/openapi.json"
                data-configuration='{"theme": "purple"}'
                src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
        </body>
        </html>
    `))
	})

	a.router.Group("api/v1")

	v1 := a.router.Group("/api/v1")



	v1.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost:5173",
            "http://localhost:4000",
            "https://book-app-cyan.vercel.app",
        },
        AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Idempotency-Key", "ngrok-skip-browser-warning"},
        ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
        AllowCredentials: true,
        MaxAge:           48 * time.Hour,
    }))

	transactor := repository.NewGormTransactor(a.db)

	// Repositories bootstrapping
	walletRepository := repository.NewWalletRepository(a.db)
	userRepository := repository.NewUserRepository(a.db)

	// Service bootstrapping
	authService := authservice.NewAuthService(userRepository)
	walletService := walletservice.NewWalletService(transactor, walletRepository)

	// Controllers initialization
	authController := endpoints.NewAuthController(authService)
	walletController := endpoints.NewWalletController(walletService)

	endpoints.RegisterAuthRoutes(v1, authController)
	endpoints.RegisterWalletRoutes(v1, walletController)

	v1.Use(middleware.AuthMiddleware())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + a.cfg.PORT,
		Handler: a.router,
	}

	a.router.GET("/test-shutdown", func(c *gin.Context) {
		fmt.Println("Request started...")
		time.Sleep(7 * time.Second)
		c.JSON(200, gin.H{"message": "I finished despite the shutdown!"})
		fmt.Println("Request finished!")
	})

	a.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Listen: %s\n", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("[System] Shutdown error: %v\n", err)
	}

	fmt.Println("[System] Server gracefully stopped.")
}
