package controllers

import (
	"encoding/json"
	"net/http"
	"yaonkey/inhouse/pkg/models"
	"yaonkey/inhouse/pkg/utils"

	"github.com/gorilla/mux"
)

// Контроллер, позволяющий получить время доступа к определенному сайту
// от имени обычного пользователя
func GetSiteTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	site := &models.Site{Url: params["sitename"]}
	endpoint := models.Endpoint{Url: "/{sitename}/"}
	if err := endpoint.Increment(); err != nil {
		utils.SendHTTPNotFoundError(w, err)
	}

	err := site.GetTime()
	if err != nil {
		utils.SendHTTPNotFoundError(w, err)
		return
	}

	res, err := json.Marshal(site)
	if err != nil {
		utils.SendHTTPInternalError(w, err)
	}
	utils.SendHTTPSuccessResult(w, res)
}

// Контроллер, позволяющий получить минимальное время доступа
// к определенному сайту от имени обычного пользователя
func GetMinTimeSite(w http.ResponseWriter, r *http.Request) {
	site := &models.Site{}
	endpoint := models.Endpoint{Url: "/min/"}
	if err := endpoint.Increment(); err != nil {
		utils.SendHTTPNotFoundError(w, err)
	}

	if err := site.GetMinTime(); err != nil {
		utils.SendHTTPInternalError(w, err)
		return
	}
	res, err := json.Marshal(site.Url)
	if err != nil {
		utils.SendHTTPInternalError(w, err)
	}
	utils.SendHTTPSuccessResult(w, res)
}

// Контроллер, позволяющий получить максимальное время доступа
// к определенному сайту от имени обычного пользователя
func GetMaxTimeSite(w http.ResponseWriter, r *http.Request) {
	site := &models.Site{}
	endpoint := models.Endpoint{Url: "/max/"}
	if err := endpoint.Increment(); err != nil {
		utils.SendHTTPNotFoundError(w, err)
	}

	if err := site.GetMaxTime(); err != nil {
		utils.SendHTTPInternalError(w, err)
		return
	}

	res, err := json.Marshal(site.Url)
	if err != nil {
		utils.SendHTTPInternalError(w, err)
	}

	utils.SendHTTPSuccessResult(w, res)
}
