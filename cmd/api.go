package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Chandra5468/basic-ecom/internal/adapters/postgresql/sqlc"
	"github.com/Chandra5468/basic-ecom/internal/orders"
	"github.com/Chandra5468/basic-ecom/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

// mount
func (app *application) mount() http.Handler {
	// gorilla mux or chi. Most packages implement this interface
	r := chi.NewRouter()

	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // important for rate limiting, analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good."))
	})

	// Dependency injection of service (postgres methods of products) into handler
	productService := products.NewService(repo.New(app.db))
	// dependency injection of handlers(w, r) into api (transport http,grpc layer)
	productsHandler := products.NewHandler(productService)
	r.Get("/products", productsHandler.ListProducts)

	ordersService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(ordersService)
	r.Post("/orders", ordersHandler.PlaceOrder)
	return r
}

// run

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	// db driver
	db *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
