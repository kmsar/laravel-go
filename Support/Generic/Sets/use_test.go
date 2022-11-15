package Sets

func use_test() {
	// Create a set of ints.
	// We pass int as a type argument.
	// Then we write () because Make does not take any non-type arguments.
	// We have to pass an explicit type argument to Make.
	// Function argument type inference doesn't work because the
	// type argument to Make is only used for a result parameter type.
	s := Make[int]()

	// Add the value 1 to the set s.
	s.Add(1)

	// Check that s does not contain the value 2.
	if s.Contains(2) {
		panic("unexpected 2")
	}
}
