package main

func main() {

	ch := make(chan struct{})
	MyGlobal().Server.Run()
	ch <- struct{}{}
	MyGlobal().Box.Run()
}