package seeder

import (
	"strconv"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"github.com/jaswdr/faker"
)

func GetUsers() []entity.User {
	faker := faker.New()
	users := make([]entity.User, 0)

	root := entity.
		NewUser("roots", "roots", 10_000_000).
		AddRole(entity.ROOT_ROLE).
		AddRole(entity.ADMIN_ROLE).
		AddRole(entity.CLIENT_ROLE).
		HashPassword()

	admin := entity.
		NewUser("admins", "admins", 1_000_000).
		AddRole(entity.ADMIN_ROLE).
		AddRole(entity.CLIENT_ROLE).
		HashPassword()

	users = append(users, *root, *admin)
	for i := 3; i <= 10; i++ {
		name := faker.Person().FirstName() + "-" + strconv.Itoa(i)

		client := entity.
			NewUser(name, name, faker.IntBetween(100_000, 500_000)).
			AddRole(entity.CLIENT_ROLE).
			HashPassword()
		users = append(users, *client)
	}

	return users
}
