package server

import (
    "github.com/go-chi/chi/v5"
    "net/http"
)

func NewRouter(h *Handler) http.Handler {
    r := chi.NewRouter()

    r.Post("/bottles", h.CreateBottle)
    r.Get("/bottles", h.ListBottles)
    r.Get("/bottles/{id}", h.GetBottleByID)
    r.Put("/bottles/{id}", h.UpdateBottle)
    r.Delete("/bottles/{id}", h.DeleteBottle)

    r.Post("/users", h.CreateUser)
    r.Get("/users", h.ListUsers)
    r.Get("/users/{id}", h.GetUserByID)
    r.Put("/users/{id}", h.UpdateUser)
    r.Delete("/users/{id}", h.DeleteUser)

    return r
}
