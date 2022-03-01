package database

import (
	"test-task/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	gorm *gorm.DB
}

func NewDatabse(dsn string) (*Database, error) {
	gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	gorm.AutoMigrate(&domain.User{})
	return &Database{gorm: gorm}, nil
}

func (db *Database) AddUser(newUser domain.User) {
	db.gorm.Create(&newUser)
}

func (db *Database) DeleteUser(user domain.User) {
	db.gorm.Where("name = ?", user.Name).Delete(&domain.User{})
}

func (db *Database) GetUsers() []domain.User {
	var users []domain.User

	db.gorm.Find(&users)

	return users
}
