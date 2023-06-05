package routes

import (
	"yaonkey/inhouse/pkg/controllers"

	"github.com/gorilla/mux"
)

// Инициация роутов для сервиса, перенаправляющихся на контроллеры
// # Доступные роуты:
// * POST "/{sitename}" => GetSiteTime - отображает время доступа к сайту sitename
// * GET "/min/" => GetMinTimeSite - отображает сайт с минимальным временем доступа
// * GET "/max/" => GetMaxTimeSite - отображает сайт с максимальным временем доступа
// * GET "/stats/" => GetAdminStatistics - отображает статистику по роутам для администратора
var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/min/", controllers.GetMinTimeSite).Methods("GET")
	router.HandleFunc("/max/", controllers.GetMaxTimeSite).Methods("GET")
	router.HandleFunc("/stats/", controllers.GetAdminStatistics).Methods("POST")
	router.HandleFunc("/{sitename}/", controllers.GetSiteTime).Methods("GET")
}
