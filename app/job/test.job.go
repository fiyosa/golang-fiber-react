package job

import "go-fiber-react/config"

func Test1Job(data string) {
	config.Log("Event: " + data)
}
