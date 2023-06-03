package routes

import (
	"yaonkey/inhouse/pkg/controllers"

	"github.com/gorilla/mux"
)

// Инициация роутов для сервиса, перенаправляющихся на контроллеры
// # Доступные роуты:
// * POST "/get/{sitename}" => GetSiteTime - отображает время доступа к сайту sitename
// * GET "/getmin/" => GetMinTimeSite - отображает сайт с минимальным временем доступа
// * GET "/getmax/" => GetMaxTimeSite - отображает сайт с максимальным временем доступа
// * GET "/getstats/" => GetAdminStatistics - отображает статистику по роутам для администратора
// * POST "/getstats/{endpoint}" => GetAdminStatisticsByEndpoint - отображает статистику по определенному эндпоинту endpoint для администратора
var RegisterRoutes = func(router mux.Router) {
	router.HandleFunc("/get/{sitename}", controllers.GetSiteTime).Methods("POST")
	router.HandleFunc("/getmin/", controllers.GetMinTimeSite).Methods("GET")
	router.HandleFunc("/getmax/", controllers.GetMaxTimeSite).Methods("GET")
	router.HandleFunc("/getstats/", controllers.GetAdminStatistics).Methods("GET")
	router.HandleFunc("/getstats/{endpoint}", controllers.GetAdminStatisticsByEndpoint).Methods("POST")
}
