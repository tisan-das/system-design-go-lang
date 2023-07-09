package behavioral_patterns

type User struct {
	FirstName string
	LastName  string
}

// Iterator
type Iterator interface {
	HasNext() bool
	GetNext() interface{}
}

type UserIterator struct {
	currentIndex int
	UserArray    []User
}

func (userItr *UserIterator) HasNext() bool {
	return userItr.currentIndex < len(userItr.UserArray)
}

func (userItr *UserIterator) GetNext() interface{} {
	if !userItr.HasNext() {
		return nil
	}
	user := userItr.UserArray[userItr.currentIndex]
	userItr.currentIndex++
	return &user
}

// Collection
type Collection interface {
	CreateIterator() Iterator
}

type UserCollection struct {
	UserArray []User
}

func (collection *UserCollection) CreateIterator() Iterator {
	return &UserIterator{UserArray: collection.UserArray}
}
