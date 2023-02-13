package api

import (
	conf "back_togther/config"
	model "back_togther/structure"
	"database/sql"
	"fmt"
)

var (
	ID          int
	Nom         string
	Email       string
	Login       string
	Pwd         string
	Description string
	Titre       string
	Contenu     string
	Date        string
	Id_Frm      string
	Obj_msg     string
	Id_user     string
)

func CheckError(err error) bool {
	return err != nil
}

func GetUser(login string) (bool, model.Utilisateur) {
	userData := model.Utilisateur{}
	db, _ := sql.Open("mysql", conf.ConnectionParameters)
	defer db.Close()
	sql := fmt.Sprintf("select * from %v.utilisateur where login = '%v'", conf.Mysql.DbName, login)
	result, err := db.Query(sql)
	if CheckError(err) {
		fmt.Println("Check Error Failled: ", err)
	}
	for result.Next() {
		result.Scan(&ID, &Nom, &Email, &Login, &Pwd)
		userData = model.Utilisateur{
			Id:    ID,
			Nom:   Nom,
			Email: Email,
			Login: Login,
			Pwd:   Pwd,
		}
	}
	return len(userData.Login) > 0, userData
}

func GetForm(id string) (bool, model.Forum) {
	formData := model.Forum{}
	db, _ := sql.Open("mysql", conf.ConnectionParameters)
	defer db.Close()
	sql := fmt.Sprintf("select * from %v.forum where id = '%v'", conf.Mysql.DbName, id)
	result, err := db.Query(sql)
	if CheckError(err) {
		fmt.Println("Check Error Failled: ", err)
	}
	for result.Next() {
		result.Scan(&ID, &Description)
		formData = model.Forum{
			Id:          ID,
			Description: Description,
		}
	}
	return len(formData.Description) > 0, formData
}

func GetMessageForum(id string) (bool, model.Message_Forum) {
	msgFrmData := model.Message_Forum{}
	db, _ := sql.Open("mysql", conf.ConnectionParameters)
	defer db.Close()
	sql := fmt.Sprintf("select * from %v.message_forum where id = '%v'", conf.Mysql.DbName, id)
	result, err := db.Query(sql)
	if CheckError(err) {
		fmt.Println("Check Error Failled: ", err)
	}
	for result.Next() {
		result.Scan(&ID, &Titre, &Contenu, &Date, &Id_Frm, &Id_user)
		msgFrmData = model.Message_Forum{
			Id:      ID,
			Titre:   Titre,
			Contenu: Contenu,
			Date:    Date,
			Id_Frm:  Id_Frm,
			Id_user: Id_user,
		}
	}
	return len(msgFrmData.Titre) > 0, msgFrmData
}

func GetMessage(id string) (bool, model.Message) {
	msgData := model.Message{}
	db, _ := sql.Open("mysql", conf.ConnectionParameters)
	defer db.Close()
	sql := fmt.Sprintf("select * from %v.message where id = '%v'", conf.Mysql.DbName, id)
	result, err := db.Query(sql)
	if CheckError(err) {
		fmt.Println("Check Error Failled: ", err)
	}
	for result.Next() {
		result.Scan(&ID, &Obj_msg, &Contenu)
		msgData = model.Message{
			Id:      ID,
			Obj_msg: Obj_msg,
			Contenu: Contenu,
		}
	}
	return len(msgData.Obj_msg) > 0, msgData
}

func GetAdmin(id string) (bool, model.Admin) {
	adminData := model.Admin{}
	db, _ := sql.Open("mysql", conf.ConnectionParameters)
	defer db.Close()
	sql := fmt.Sprintf("select * from %v.admin where id = '%v'", conf.Mysql.DbName, id)
	result, err := db.Query(sql)
	if CheckError(err) {
		fmt.Println("Check Error Failled: ", err)
	}
	for result.Next() {
		result.Scan(&ID, &Login, &Pwd)
		adminData = model.Admin{
			Id:    ID,
			Login: Login,
			Pwd:   Pwd,
		}
	}
	return len(adminData.Login) > 0, adminData
}
