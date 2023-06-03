package controllers

import "net/http"

// Контроллер, позволяющий получить время доступа к определенному сайту
// от имени обычного пользователя
func GetSiteTime(w http.ResponseWriter, r *http.Request) {}

// Контроллер, позволяющий получить минимальное время доступа
// к определенному сайту от имени обычного пользователя
func GetMinTimeSite(w http.ResponseWriter, r *http.Request) {}

// Контроллер, позволяющий получить максимальное время доступа
// к определенному сайту от имени обычного пользователя
func GetMaxTimeSite(w http.ResponseWriter, r *http.Request) {}
