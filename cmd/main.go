package main

import (
	"plant-monitoring-server/internal/moisturepipe"
	"sync"
)

func main() {
	var wait_group sync.WaitGroup

	var moisture_value float32

	wait_group.Add(1)
	go moisturepipe.ReadFromMoisturePipe(&wait_group, &moisture_value)

}
