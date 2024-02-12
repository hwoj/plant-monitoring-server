package main

import (
	"fmt"
	"log"
	"net/http"
	"plant-monitoring-server/internal/readmoisture"
	"plant-monitoring-server/internal/websocket"
	"sync"
)

func main() {

	http.HandleFunc("/ws", websocket.HandleWebSocket)
	http.HandleFunc("/", HandleRoot)

	var wait_group sync.WaitGroup

	var moisture_value float32

	wait_group.Add(1)
	go readmoisture.ReadFromSharedMemory(&wait_group, &moisture_value)

	wait_group.Add(1)
	go log.Fatal(http.ListenAndServe(":8080", nil))

	wait_group.Wait()

}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLOOOOO WORLD")
}
