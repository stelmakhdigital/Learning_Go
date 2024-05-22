package main

import (
	"io"
	"net/http"
)

const formCONST = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`


func Auth(login, password string) bool {
	return login == "admin" && password == "admin"
}

func mainPage(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		login := request.FormValue("login")
		password := request.FormValue("password")
		if Auth(login, password) {
			io.WriteString(response, "Success log in")
		} else {
			http.Error(response, "Uncorrect password or login", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(response, formCONST)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}