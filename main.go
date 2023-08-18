package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var savedData string

type ResponseData struct {
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, мир!")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Ожидается POST-запрос", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	savedData = string(body) // Сохраняем данные в глобальной переменной
	fmt.Printf("Получено POST-тело: %s\n", savedData)

	response := ResponseData{
		Message: "Запись создана",
		Data:    savedData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Ожидается PUT-запрос. Вы лоххххх", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	savedData = string(body) // Сохраняем данные в глобальной переменной
	fmt.Printf("Получено PUT-тело: %s\n", savedData)

	response := ResponseData{
		Message: "Запись перезаписана",
		Data:    savedData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)   // обработчик пути
	http.HandleFunc("/create", createHandler) // второй эндпоинт
	http.HandleFunc("/update", updateHandler) // третий эндпоинт
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
