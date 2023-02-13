package structure

var StatusOK int = 200
var StatusErrData int = 404
var StatusNoData int = 300

type Utilisateur struct {
	Id    int    `json:"id"`
	Nom   string `json:"nom"`
	Email string `json:"email"`
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

type Forum struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    Utilisateur `json:"data"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Message_Forum struct {
	Id      int    `json:"id"`
	Titre   string `json:"titre"`
	Contenu string `json:"contenu"`
	Date    string `json:"date"`
	Id_Frm  string `json:"id_frm"`
	Id_user string `json:"id_user"`
}

type Message struct {
	Id      int    `json:"id"`
	Obj_msg string `json:"obj_msg"`
	Contenu string `json:"contenu"`
}

type Admin struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

type Connection struct {
	Host     string
	Port     string
	Password string
	DbName   string
}
