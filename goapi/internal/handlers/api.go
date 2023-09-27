package handlers

import(
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "goapi/internal/middleware"
)

func Hander(r *chi.Mux) {
    r.Use(chimiddle.StripSlashes)
    r.Route("/account", func (router chi.Router){
        router.Use(middleware.Authorization)

        router.Get("/coins", GetCoinBalance)
    })
}
