package main

import (
	"net/http"
)

//type ResponseWriter interface {
//	Header() Header
//	Write([]byte) (int, error)
//	WriteHeader(statusCode int)
//}
//
//
//type Handler interface {
//	ServeHTP(ResponseWriter, *Request)
//}

type CustomHandler struct {}

func (handler CustomHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	data := []byte("hello world")
	response.Write(data)
}

// добавляем маршруты и обработчики к DefaultServeMux
func  apiPage(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Это страница API"))
}

func docsPage(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Это страница документации к API"))
}

func main() {
	//var handler CustomHandler

	//// В HandleFunc нельзя использовать регулярные выражения и маски - поэтому луче не использовать http.DefaultServeMux
	//// Порядок не важен
	//http.HandleFunc("/api", apiPage)
	//http.HandleFunc("/docs", docsPage)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/", apiPage) // обработает и /api и /api/v1/...
	mux.HandleFunc("/docs", docsPage)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
