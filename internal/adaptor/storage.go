package adaptor

import (
	"errors"
	"part2/internal/model"

	"github.com/google/uuid"
)

type ConcurrentUserStorage struct {
	db map[string]*model.User
	//mx sync.Mutex
}

func CreateConcurrentUserStorage() *ConcurrentUserStorage {
	return &ConcurrentUserStorage{
		db: make(map[string]*model.User),
		//mx: sync.Mutex{},
	}
}

func validateUser(user *model.User) error {
	if user.ID == uuid.Nil {
		return errors.New("ID is null")
	}
	if user.Login == "" {
		return errors.New("Login is null")
	}
	if user.Password == "" {
		return errors.New("Login is null")
	}
	return nil

}

func (u *ConcurrentUserStorage) Get(login string) *model.User {
	if u.db[login] != nil {
		return u.db[login]
	}
	return nil
}

func (u *ConcurrentUserStorage) Update(user *model.User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}
	if u.db[user.Login] == nil {
		return errors.New("This login doesn't exist")
	}
	u.db[user.Login] = user
	return nil
}

func (u *ConcurrentUserStorage) Add(user *model.User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}
	if u.db[user.Login] != nil {
		return errors.New("This login already exist")
	}
	u.db[user.Login] = user
	return nil
}
