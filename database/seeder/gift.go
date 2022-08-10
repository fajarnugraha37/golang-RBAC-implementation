package seeder

import (
	"strconv"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"github.com/jaswdr/faker"
)

func GetGifts() []entity.Gift {
	faker := faker.New()
	gifts := make([]entity.Gift, 0)
	for i := 1; i <= 25; i++ {
		gifts = append(gifts, entity.Gift{
			Name:        faker.App().Name(),
			Point:       uint(faker.IntBetween(50_000, 500_000)),
			Excerpt:     faker.Lorem().Sentence(5),
			Description: faker.Lorem().Sentence(12),
			Image:       "/assets/images/" + strconv.Itoa(faker.IntBetween(1, 18)) + ".png",
			Stock:       uint16(faker.IntBetween(1, 250)),
		})
	}
	return gifts
}
