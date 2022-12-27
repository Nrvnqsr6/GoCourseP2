package adaptor

import (
	"errors"
	"part2/internal/model"
	"sync"

	"github.com/google/uuid"
)

type ConcurrentUserStorage struct {
	db map[string]*model.User
	mx sync.RWMutex
}

func (s *ConcurrentUserStorage) GetDB() map[string]*model.User {
	return s.db
}

func CreateConcurrentUserStorage() *ConcurrentUserStorage {
	return &ConcurrentUserStorage{
		db: make(map[string]*model.User),
		mx: sync.RWMutex{},
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

func (u *ConcurrentUserStorage) Authorize(login string, password string) error {
	u.mx.RLock()
	defer u.mx.RUnlock()
	if u.db[login] == nil {
		return errors.New("There is no such user in the db")
	}
	if u.db[login].Password != password {
		return errors.New("Wrong password")
	}
	return nil
}

func (u *ConcurrentUserStorage) Get(login string) *model.User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	if u.db[login] != nil {
		return u.db[login]
	}
	return nil
}

func (u *ConcurrentUserStorage) Update(user *model.User, userUpdate *model.User) error {
	u.mx.Lock()
	defer u.mx.Unlock()
	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}
	if userUpdate.Password != "" {
		user.Password = userUpdate.Password
	}
	if userUpdate.Phone != "" {
		user.Phone = userUpdate.Phone
	}
	if userUpdate.BirthDate != "" {
		user.BirthDate = userUpdate.BirthDate
	}

	//err := validateUser(user)
	// if err != nil {
	// 	return err
	// }
	// if u.db[user.Login] == nil {
	// 	return errors.New("This login doesn't exist")
	// }
	//u.db[user.Login] = user
	return nil
}

func (u *ConcurrentUserStorage) Add(user *model.User) error {
	u.mx.Lock()
	defer u.mx.Unlock()
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
