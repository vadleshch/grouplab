package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/spf13/viper"

    "github.com/vadleshch/grouplab/internal/bottle"
    "github.com/vadleshch/grouplab/internal/server"
    "github.com/vadleshch/grouplab/internal/user"
)

func main() {
    viper.SetConfigName("env")   
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("config read error: %v", err)
    }

    port := viper.GetString("server.port")
    dbURL := viper.GetString("postgres.dsn")
    if port == "" || dbURL == "" {
        log.Fatal("PORT or DSN not set in config/env.yaml")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    dbpool, err := pgxpool.New(ctx, dbURL)
    if err != nil {
        log.Fatalf("can't connect to postgres: %v", err)
    }
    defer dbpool.Close()


    bottleStorage := bottle.NewStorage(dbpool)
    userStorage := user.NewStorage(dbpool)

    handler := &server.Handler{
        BottleStorage: bottleStorage,
        UserStorage:   userStorage,
    }
    router := server.NewRouter(handler)


    addr := fmt.Sprintf(":%s", port)
    srv := &http.Server{
        Addr:    addr,
        Handler: router,
    }
    log.Printf("Server running on %s", addr)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server error: %v", err)
    }
}
