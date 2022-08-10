package gift

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type GiftRepository interface {
	GetById(id int, relation *entity.GiftRelation) (*entity.Gift, error)
	GetMany(limit, offset int, order, sort string, relation *entity.GiftRelation) (*[]entity.Gift, error)
	GetCount() (int64, error)
	Create(gift *entity.Gift) error
	Update(gift *entity.Gift) error
	Delete(id int) (*entity.Gift, error)
}
