package listener

import "go-fiber-react/config"

func Test2Listener(data string) {
	config.Log("Listener 2: " + data)
}
