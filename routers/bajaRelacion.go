package routers

import (
	"net/http"

	"github.com/grak0s/twittor/db"
	"github.com/grak0s/twittor/models"
)

/*BajaRelacion realiza el borrado de la relacion entre usuarios */
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
