package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"yaonkey/inhouse/pkg/models"
	"yaonkey/inhouse/pkg/utils"
)

// Контроллер, позволяющий администраторам получить статистику
// по эндпоинтам от имени администратора.
//
// Для получения статистики администратором должны быть переданы
// логин admin_login и хеш пароля admin_password_hash.
//
// Так как клиентская часть отсутствует, где формировался бы хеш,
// можно использовать запрос с прописанным телом
// из корня в файле admin.http
func GetAdminStatistics(w http.ResponseWriter, r *http.Request) {
	admin := &models.Admin{}
	utils.ParseBody(r, &admin)

	if !models.IsAdmin(admin) {
		utils.SendCustomError(w, http.StatusForbidden, errors.New("denied access"))
		return
	}

	endpoints, err := models.GetAllEndpoints()
	if err != nil {
		utils.SendHTTPInternalError(w, err)
		return
	}
	res, err := json.Marshal(endpoints)
	if err != nil {
		utils.SendHTTPInternalError(w, err)
	}
	utils.SendHTTPSuccessResult(w, res)
}
