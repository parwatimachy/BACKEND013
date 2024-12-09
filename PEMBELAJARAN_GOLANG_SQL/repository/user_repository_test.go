package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-database-user/config"
	"golang-database-user/model"
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserInsert_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	userModel := model.MstUser{}

	UserRepository := NewUserRepositoryImpl(sql)
	RoleRepository := NewRoleRepositoryImpl(sql)

	ctx := context.Background()
	emailExists, err := UserRepository.EmailExists(ctx, userModel.Email)

	if emailExists {
		panic("Email sudah terdaftar")
	}

	uuidUser := uuid.New().String()

	theRole, err := RoleRepository.FindMstRole(ctx, "ROLE001")
	if err != nil {
		panic(err)
	}

	user := model.MstUser{
		IdUser:      uuidUser,
		Name:        userModel.Name,
		Email:       userModel.Email,
		Password:    userModel.Password,
		PhoneNumber: userModel.PhoneNumber,
		Role:        theRole,
	}

	insertUser, err := UserRepository.InsertUser(ctx, user)

	assert.NotNil(t, insertUser)
	assert.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}
	defer sql.Close()

	UserRepository := NewUserRepositoryImpl(sql)

	ctx := context.Background()

	// userModel := model.MstUser{}

	existingUserID := "194eca94-b078-4232-a004-9907e081daf1"
	user := model.MstUser{
		Name:        "Purwanto",
		Email:       "purwanto@gmail.com",
		Password:    "123",
		PhoneNumber: "0987654321",
	}

	UpdateUser, err := UserRepository.UpdateUser(ctx, user, existingUserID)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, UpdateUser)
	assert.Nil(t, err)
}

func TestReadUser(t *testing.T){
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)

	users, err := UserRepository.ReadUsers(ctx)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, users)
	assert.Nil(t, err)
}

func TestDeleteUser(t *testing.T){
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)

	userId := "194eca94-b078-4232-a004-9907e081daf1"

	if userId == "" {
		fmt.Print(errors.New("ID user tidak boleh kosong"))
	}

	deleteUser, err := UserRepository.DeleteUser(ctx, userId)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, deleteUser)
	assert.Nil(t, err)
}