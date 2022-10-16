package test

import (
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	db := InitDbTest()
	userRepo := repository.GetUserRepositoryDB(db)

	newUser := repository.CreateUserModelMoc()
	newUser.Id = 0

	createdUser, err := userRepo.CreateByModel(newUser)
	if assert.NoError(t, err, "Create new user error") {
		defer func() {
			err = userRepo.ForceUserDelete(createdUser)
			assert.NoError(t, err, "Delete created user error")
		}()
		assert.NotNil(t, createdUser, "User not created without error")
		t.Log("User created")
		t.Log(createdUser)

		userById, err := userRepo.GetById(createdUser.Id)
		assert.NoError(t, err, "Can not get created user by Id")
		assert.Equal(t, createdUser.PublicId, userById.PublicId, "Created user and received by Id user is not equals")

		userByPublicId, err := userRepo.GetByPublicId(createdUser.PublicId)
		assert.NoError(t, err, "Can not get created user by PublicId")
		assert.Equal(t, userByPublicId.Id, userById.Id, "Created user and received by PublicId user is not equals")

	}
}
