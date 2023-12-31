package client

import (
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
)

func (api *ClientAPI) Welcome(w http.ResponseWriter, r *http.Request) {
	// projects, err := api.store.Project.GetAll(context.TODO(), 4, 1)
	// if err != nil {
	// 	helpers.ErrResponse(w, err, http.StatusInternalServerError)
	// 	return
	// }

	err := api.temp.ExecuteTemplate(w, "welcome.html", nil)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}
