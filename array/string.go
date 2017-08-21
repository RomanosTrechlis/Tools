package array

type StringArray []string

// Head returns the first element of a StringArray.
func (a StringArray) Head() interface{} {
	return a[0]
}

// Tail returns a new StringArray that excludes the first element
func (a StringArray) Tail() interface{} {
	return a[1 :]
}
