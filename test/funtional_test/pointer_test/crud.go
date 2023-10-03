package pointer

// ICrud is an interface for CRUD operations.
type ICrud interface {
	Create(item any)
	Read(id int) any
	Update(item any)
	Delete(id int)
}

// User is a struct for user data.
type User struct {
	ID   int
	Name string
}

// UserCrud is a struct for user CRUD operations.
type UserCrud struct {
	users []User
}

// Create creates a new user.
func (crud *UserCrud) Create(item any) {
	user := item.(User)
	crud.users = append(crud.users, user)
}

// Read reads a user by id.
func (crud *UserCrud) Read(id int) any {
	for _, user := range crud.users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

// Update updates a user.
func (crud *UserCrud) Update(item any) {
	user := item.(User)
	for i, u := range crud.users {
		if u.ID == user.ID {
			crud.users[i] = user
			return
		}
	}
}

// Delete deletes a user by id.
func (crud *UserCrud) Delete(id int) {
	for i, user := range crud.users {
		if user.ID == id {
			crud.users = append(crud.users[:i], crud.users[i+1:]...)
			return
		}
	}
}

func (crud *UserCrud) GetUsers() *[]User {
	return &crud.users
}
