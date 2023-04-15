package main

import (
	api "back_togther/api"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// write ReadMe of this code in github format
func Handlers() {

	http.HandleFunc("/user", api.GetUserData)
	http.HandleFunc("/allUser", api.GetAllUserData)

	http.HandleFunc("/forum", api.GetForumData)
	http.HandleFunc("/message_forum", api.GetMessageForumData)
	http.HandleFunc("/message", api.GetMessageData)
	http.HandleFunc("/admin", api.GetAdminData)

	port := 8090
	fmt.Println("Server is running on ", port)
	addr := fmt.Sprintf(":%v", port)
	http.ListenAndServe(addr, nil)
}

func main() {

	Handlers()
}
