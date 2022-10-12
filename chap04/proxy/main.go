package proxy

import "fmt"

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for _, u := range *t {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found", id)
}

func (t *UserList) addUser(u User) {
	*t = append(*t, u)
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	//Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)
	fmt.Println("Returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}
