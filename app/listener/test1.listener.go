package listener

import "go-fiber-react/config"

func Test1Listener(data string) {
	config.Log("Listener 1: " + data)
}
