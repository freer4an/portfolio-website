package client

import (
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
)

func (api *ProjectAPI) Welcome(w http.ResponseWriter, r *http.Request) {
	// projects, err := api.store.Project.GetAll(context.TODO(), 4, 1)
	// if err != nil {
	// 	helpers.ErrResponse(w, err, http.StatusInternalServerError)
	// 	return
	// }

	err := temp.ExecuteTemplate(w, "welcome.html", nil)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}
