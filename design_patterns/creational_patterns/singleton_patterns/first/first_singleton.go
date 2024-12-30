package first

// A simple demonstration of singlet design pattern
// A unique counter
// As an example of an object of which we must ensure that there is only one instance we will counter
// that holds the number of time it has been called during program execution
// with these requirements... we can write the following acceptance criteria
// Requirements and acceptance criteria
// When no counter has been created before a new one is created with the value 0
// If a counter has already been created, return this instance that holds the actual count
// If we call the method AddOne, the count must be increment by 1

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	// create an instance of the singleton struct
	// and ensure it is created only once through out the application lifecycle
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

// use the same instantiated object to increment the count method
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
