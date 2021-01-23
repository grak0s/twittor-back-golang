package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/grak0s/twittor/middlew"
	"github.com/grak0s/twittor/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo de puerto y pone a escuchar el servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")

	router.HandleFunc("/verperfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminoTweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoDB(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoDB(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
