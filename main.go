package main

import (
	api "back_togther/api"
	"fmt"
	"net/http"
)

func Handlers() {
	http.HandleFunc("/user", api.GetUserData)
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
