package main

import (
	observer "github.com/giap9416/observer"
)

func main() {
	post := observer.NewPost("Golang Design Patterns: Observer")

	subscriber1 := &observer.Subscriber{ID: "subscriber1@gmail.com"}
	subscriber2 := &observer.Subscriber{ID: "subscriber2@gmail.com"}

	post.Register(subscriber1)
	post.Register(subscriber2)

	post.UpdateStatus()

}
