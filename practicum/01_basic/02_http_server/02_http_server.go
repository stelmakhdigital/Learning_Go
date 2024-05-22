package main

import (
	"encoding/json"
	"fmt"
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

type Sybject struct {
	Product string `json:"name"`
	Price int `json:"price"`
}

func (handler CustomHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	data := []byte("hello world")
	response.Write(data)
}

// добавляем маршруты и обработчики к DefaultServeMux
func  apiPage(response http.ResponseWriter, request *http.Request) {
	body := fmt.Sprint("Method: %s\r\n", request.Method)
	body += "Header ===============\r\n"
	for key, value := range request.Header {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	body += "Query parameters ===============\r\n"
	for key, value := range request.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", key, value)
	}
	response.Write([]byte(body))
}

func docsPage(response http.ResponseWriter, request *http.Request) {
	// Разрещаем только GET метод
	if request.Method != http.MethodGet {
		http.Error(response, "Разрешен только GET метод", http.StatusMethodNotAllowed)
		return
	}
	response.Write([]byte("Это страница документации к API"))
}

func mainPage(response http.ResponseWriter, request *http.Request) {
	body := "Form Params ==============\r\n"
	if err := request.ParseForm(); err != nil {
		response.Write([]byte(err.Error()))
		return
	}
	for key, value := range request.Form {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	response.Write([]byte(body))
}

func JSONHandler(response http.ResponseWriter, request *http.Request) {
	// собаираем данные в структуру
	subj := Sybject{"Миска риса", 500}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	// записываем заголовок Content-type
	// для передачи клиенту информации о json
	response.Header().Set("Content-Type", "application/json")
	// устанавливаем код 200
	response.WriteHeader(http.StatusOK)
	// пишем тело запроса json
	response.Write(resp)
}


func main() {
	//var handler CustomHandler

	//// В HandleFunc нельзя использовать регулярные выражения и маски - поэтому луче не использовать http.DefaultServeMux
	//// Порядок не важен
	//http.HandleFunc("/api", apiPage)
	//http.HandleFunc("/docs", docsPage)

	mux := http.NewServeMux()
	// вызывается обработчик с максимальным количеством совподающих символов с URL
	mux.HandleFunc("/api/", apiPage) // обработает и /api и /api/v1/...
	mux.HandleFunc("/docs", docsPage)
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/json", JSONHandler)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
