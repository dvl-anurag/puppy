package puppy

import "github.com/dvl-anurag/dog"

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof woof woof woof"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func FromTwo() string {
	return "this is version 1.2"
}
