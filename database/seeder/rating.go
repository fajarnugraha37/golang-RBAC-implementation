package seeder

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"github.com/jaswdr/faker"
)

func GetRatings() []entity.Rating {
	faker := faker.New()
	ratings := make([]entity.Rating, 0)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			ratings = append(ratings, entity.Rating{
				UserID: uint(i),
				GiftID: uint(faker.IntBetween(1, 25)),
				Score:  uint(faker.IntBetween(1, 5)),
			})
		}
	}
	return ratings
}
