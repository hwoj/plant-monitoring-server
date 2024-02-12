package readmoisture

import (
	"fmt"
	"os"
	"sync"
	"syscall"
)

func ReadFromSharedMemory(wait_group *sync.WaitGroup, moisture_value *float32) {

	pipeName := "moisturepipe"

	namedPipe, err := os.OpenFile(pipeName, syscall.O_RDONLY|syscall.O_NONBLOCK, os.ModeNamedPipe)
	fmt.Println("Pipe opened")

	if err != nil {
		fmt.Println("Error from opening pipe:", err)
	}

	defer namedPipe.Close()

	buf := make([]byte, 50)
	for {
		fmt.Println("HERE")
		bytes, err := namedPipe.Read(buf)

		if err != nil {
			fmt.Printf("Error reading %d from named pipe: %s", bytes, err)
		}
		*moisture_value = 2.2 + float32(bytes)
	}

}
