package main

import "hyphen-hellog/initializer"

func main() {
	config := initializer.NewConfig()
	initializer.NewDatabase(config)

}
