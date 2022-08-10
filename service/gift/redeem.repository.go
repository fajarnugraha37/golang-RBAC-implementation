package gift

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type RedeemRepository interface {
	Create(redeem *entity.Redeem) error
}
