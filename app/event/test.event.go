package event

import "go-fiber-react/app/listener"

func TestEvent(data string) {
	listener.Test1Listener(data)
	go listener.Test2Listener(data)
}
