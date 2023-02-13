package routes

import (
	"net/http"

	"lms/controller/CAuth"
	"lms/controller/CHome"
	"lms/controller/CPeserta"
	"lms/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", middlewares.RenderKeJSON(CAuth.Login)).Methods("POST")
	router.HandleFunc("/refresh", CAuth.RefreshToken).Methods("POST")
	router.HandleFunc("/logout", middlewares.RenderKeJSON(CAuth.Logout)).Methods("POST")
	router.HandleFunc("/api/home", middlewares.RenderKeJSON(middlewares.HarusAuth(CHome.Home))).Methods("GET")
	router.HandleFunc("/api/refresh", middlewares.RenderKeJSON(middlewares.HarusAuth(CAuth.RefreshToken))).Methods("POST")

	router.HandleFunc("/api/peserta", middlewares.RenderKeJSON(middlewares.HarusAuth(CPeserta.ViewPeserta))).Methods("GET")
	router.HandleFunc("/api/peserta", middlewares.RenderKeJSON(middlewares.HarusAuth(CPeserta.CreatePeserta))).Methods("POST")
	router.HandleFunc("/api/peserta", middlewares.RenderKeJSON(middlewares.HarusAuth(CPeserta.UpdatePeserta))).Methods("PUT")
	router.HandleFunc("/api/peserta", middlewares.RenderKeJSON(middlewares.HarusAuth(CPeserta.DeletePeserta))).Methods("DELETE")

	// serving file server
	var imgServer = http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", imgServer))
	return router
}
