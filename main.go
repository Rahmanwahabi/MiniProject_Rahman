package main

import (
	route "mini-project/route"
)

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
