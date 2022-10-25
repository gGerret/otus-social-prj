package test

import (
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndGetUser(t *testing.T) {
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

func TestGetTestUser(t *testing.T) {
	db := InitDbTest()
	userRepo := repository.GetUserRepositoryDB(db)

	userByPublicId, err := userRepo.GetByPublicId("53722bba-4030-453f-8578-dc1d3941069c")
	assert.NoError(t, err, "Can not get created user by PublicId")
	assert.Equal(t, userByPublicId.Id, int64(3), "Created user and received by PublicId user is not equals")
}

func TestFriendship(t *testing.T) {
	db := InitDbTest()
	userRepo := repository.GetUserRepositoryDB(db)

	newUser1 := repository.CreateUserModelMoc()
	newUser1.Id = 0
	createdUser1, err := userRepo.CreateByModel(newUser1)
	assert.NoError(t, err, "Create new user1 error")
	defer func() {
		err = userRepo.ForceUserDelete(createdUser1)
		assert.NoError(t, err, "Delete created user error")
	}()

	newUser2 := repository.CreateUserModelMoc()
	newUser2.Id = 0
	newUser2.FirstName = "Витёк"
	createdUser2, err := userRepo.CreateByModel(newUser2)
	assert.NoError(t, err, "Create new user1 error")
	defer func() {
		err = userRepo.ForceUserDelete(createdUser2)
		assert.NoError(t, err, "Delete created user error")
	}()

	newUser3 := repository.CreateUserModelMoc()
	newUser3.Id = 0
	newUser3.FirstName = "Лёха"
	createdUser3, err := userRepo.CreateByModel(newUser3)
	assert.NoError(t, err, "Create new user1 error")
	defer func() {
		err = userRepo.ForceUserDelete(createdUser3)
		assert.NoError(t, err, "Delete created user error")
	}()

	err = userRepo.CreateFriendshipLink(createdUser1, createdUser2, "Приветики 1")
	assert.NoError(t, err, "Create friendship 1 -> 2 error")

	err = userRepo.CreateFriendshipLink(createdUser1, newUser3, "Приветики 2")
	assert.NoError(t, err, "Create friendship 1 -> 3 error by userB.PublicId")

	err = userRepo.CreateFriendshipLink(createdUser1, createdUser3, "Приветики 1")
	assert.Error(t, err, "Must be an error - friendship 1 -> 3 already exists")

	err = userRepo.CreateFriendshipLink(createdUser2, createdUser3, "Приветики 3")
	assert.NoError(t, err, "Create friendship 2 -> 3 error")

	user1Friends, err := userRepo.GetUserFriends(createdUser1)
	assert.NoError(t, err, "Get user 1 friends error")
	assert.Len(t, user1Friends, 2, "Wrong user 1 friends count")
	assert.Equal(t, createdUser2.Id, user1Friends[0].Id, "Wrong first friend of user 1")
	assert.Equal(t, createdUser3.Id, user1Friends[1].Id, "Wrong second friend of user 1")

}
