package client

//структура запроса изменения данных клиента
type RequestUpdateClient struct {
	Client_Login      string `json:"login" validate:"required,min=3,max=20"`
	Client_Name       string `json:"name" validate:"required"`
	Client_Firstname  string `json:"firstname" validate:"required"`
	Client_Patronymic string `json:"patronymic"`
	Client_Email      string `json:"email" validate:"required,email"`
}

//структура запроса регистрации клиента
type RequestCreateClient struct {
	Client_Login           string `json:"login" validate:"required,min=3,max=20"`
	Client_Password        string `json:"password" validate:"required,min=6,max=50"`
	Client_Repeat_Password string `json:"repeat_password" validate:"required,min=6,max=50"`
	Client_Name            string `json:"name" validate:"required"`
	Client_Firstname       string `json:"firstname" validate:"required"`
	Client_Patronymic      string `json:"patronymic"`
	Client_Phone_Number    string `json:"phone" validate:"required,numeric"`
	Client_Email           string `json:"email" validate:"required,email"`
}

//структура запроса авторизации клиента
type RequestLoginClient struct {
	Client_Login    string `json:"login" validate:"required,min=3,max=20"`
	Client_Password string `json:"password" validate:"required,min=6,max=50"`
}

//структура запроса смены пароля
type RequestChangePassword struct {
	Client_Old_Password    string `json:"client_old_password" validate:"required,min=6, max=50"`
	Client_New_Password    string `json:"client_new_password" validate:"required,min=6, max=50"`
	Client_Repeat_Password string `json:"client_repeat_password" validate:"required,min=6, max=50"`
}
