// https://stackoverflow.com/questions/27198193/how-to-react-to-keypress-events-in-go
package main

import (
	"fmt"

	"gopkg.in/azul3d/keyboard.v1"
)

func main() {
	watcher := keyboard.NewWatcher()
	// Query for the map containing information about all keys
	status := watcher.States()
	left := status[keyboard.A]
	if left == keyboard.Down {
		fmt.Println("...")
	}
	fmt.Println("...2", left)
}
