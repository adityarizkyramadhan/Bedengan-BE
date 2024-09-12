package repositories

import (
	"fmt"
	"time"

	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		db    *gorm.DB
		redis *redis.Client
	}
	UserInterface interface {
		FindOne(id string) (*model.User, error)
		FindEmail(email string) (*model.User, error)
		Create(user *model.UserCreate) (*model.User, error)
		Update(id string, user *model.UserUpdate) (*model.User, error)
		Delete(id string) error
		Login(email, password string) (*model.User, error)
		Logout(token string, expired time.Duration) error
	}
)

// NewUserRepository will create an object that represent the User.Repository interface
func NewUserRepository(db *gorm.DB, redis *redis.Client) UserInterface {
	return &User{db, redis}
}

// FindOne will return a user by id
func (u *User) FindOne(id string) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	return &user, nil
}

// FindEmail will return a user by email
func (u *User) FindEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	return &user, nil
}

// Create will create a new user
func (u *User) Create(user *model.UserCreate) (*model.User, error) {
	if user.Password != user.ConfirmPassword {
		return nil, utils.NewError(utils.ErrValidation, "password and confirm password must be the same")
	}

	link, err := utils.SaveFile(user.FileKTP, fmt.Sprintf("public/images/%s", user.FileKTP.Filename))
	if err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal menyimpan file")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal membuat password")
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashPassword),
		Role:     "user",
		LinkKTP:  link,
	}

	var searchUser model.User
	if err := u.db.Where("email = ?", user.Email).First(&searchUser).Error; err != nil {
		if err := u.db.Create(&newUser).Error; err != nil {
			return nil, utils.NewError(utils.ErrUnknown, "gagal membuat user")
		}
	} else {
		// update user yang sudah ada
		searchUser.Name = user.Name
		searchUser.Role = "user"
		if err := u.db.Save(&searchUser).Error; err != nil {
			return nil, utils.NewError(utils.ErrUnknown, "gagal membuat user")
		}
		newUser = searchUser
	}

	return &newUser, nil
}

// Update will update a user by id dengan check field yang tidak dirubah maka tidak diupdate
func (u *User) Update(id string, user *model.UserUpdate) (*model.User, error) {
	var oldUser model.User
	if err := u.db.First(&oldUser, id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	if user.Name != "" {
		oldUser.Name = user.Name
	}
	if err := u.db.Save(&oldUser).Error; err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal update user")
	}
	return &oldUser, nil
}

// Delete will delete a user by id
func (u *User) Delete(id string) error {
	if err := u.db.Delete(&model.User{}, id).Error; err != nil {
		return utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	return nil
}

// Login will login a user
func (u *User) Login(email, password string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "pengguna tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, utils.NewError(utils.ErrValidation, "password salah")
	}

	return &user, nil
}

func (u *User) Logout(token string, expired time.Duration) error {
	err := u.redis.Set(u.db.Statement.Context, token, true, expired).Err()
	if err != nil {
		return utils.NewError(utils.ErrUnknown, "gagal logout")
	}
	return nil
}
