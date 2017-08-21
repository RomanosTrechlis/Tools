package array

// ArrayInterface describes Head and Tail functionality for array types.
type ArrayInterface interface {
	// Head returns the first element of an array
	Head() interface{}
	// Tail returns a new array that excludes the first element
	Tail() interface{}
}
