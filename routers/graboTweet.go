package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/grak0s/twittor/db"
	"github.com/grak0s/twittor/models"
)

//GraboTweet permite grabar el twiit en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado insertar en el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
