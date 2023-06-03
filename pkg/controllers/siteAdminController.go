package controllers

import "net/http"

// Контроллер, позволяющий администраторам получить статистику
// по эндпоинтам от имени администратора
func GetAdminStatistics(w http.ResponseWriter, r *http.Request) {}

// Контроллер, позволяющий получить статистику по отдельному эндпоинту
// от имени администратора
func GetAdminStatisticsByEndpoint(w http.ResponseWriter, r *http.Request) {}
