package main

import (
	"feed/meta/library/logs"
	"feed/meta/routers"
)

func main() {
	r := routers.RegisterRouter()
	r.Use(logs.LoggerToFile())
	_ = r.Run(":8000")
}
