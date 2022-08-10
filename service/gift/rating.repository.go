package gift

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type RatingRepository interface {
	Upsert(rating *entity.Rating) error
}
