package main

func main() {
	app := InitApp()
	app.WebServer.Spin()
	panic("WebServer fail")
}
