package main

import (
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	//Получаем hostname  
	str, err := os.Hostname()
	fmt.Fprintf(w, "Hostname: %v\n", str)
	
	// Получаем все доступные сетевые интерфейсы
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// Список адресов для каждого сетевого интерфейса
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Сетевой интерфейс: %s\n", interf.Name)

		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Fprintf(w, "\tIP: %v\n", ip)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
