package moisturepipe

import (
	"fmt"
	"os"
	"sync"
	"syscall"
)

func ReadFromMoisturePipe(wait_group *sync.WaitGroup, moisture_value *float32) {

	pipeName := "moisturepipe"

	_ = os.Remove(pipeName)

	err := syscall.Mkfifo(pipeName, 0666)

	if err != nil {
		fmt.Println("Error from creating pipe:", err)

	}

	namedPipe, err := os.OpenFile(pipeName, os.O_RDONLY, os.ModeNamedPipe)

	if err != nil {
		fmt.Println("Error from opening pipe:", err)
	}

	defer namedPipe.Close()

	buf := make([]byte, 50)
	for {
		bytes, err := namedPipe.Read(buf)

		if err != nil {
			fmt.Printf("Error reading %d from named pipe: %s", bytes, err)
		}
		*moisture_value = 2.2 + float32(bytes)
	}

}
