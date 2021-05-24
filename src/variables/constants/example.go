package main

func main() {
	// [keyword] [name] [data_type] = [value]
	const name string = "Amir"

	const message = "Golang is a programming language" // string constant
	const rating = 2.4 // floating constant
	const percentage = 100 // integer constant

	// Oops!, we cannot modify the value of a constant
	message = "Golang is not a programming language"

	// Oops!, the declaration and assignment must happen in the same line
	const address
	address = "some address"

}
