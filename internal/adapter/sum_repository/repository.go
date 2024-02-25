package sum_repository

import (
	"gorm.io/gorm"
)

type Data struct {
	ValueA   int
	ValueB   int
	Operator string
	Result   int
}

type SumRepository struct {
	db *gorm.DB
}

func (rep SumRepository) GetSum(a, b int) (int, error) {
	data := Data{
		ValueA:   a,
		ValueB:   b,
		Operator: "+",
	}
	rep.db.First(&data)
	return data.Result, rep.db.Error
}

func (rep SumRepository) SetSum(a, b, res int) error {
	data := Data{
		ValueA:   a,
		ValueB:   b,
		Operator: "+",
	}
	rep.db.Create(&data)
	return rep.db.Error
}

func NewSumRepository(db *gorm.DB) *SumRepository {
	db.AutoMigrate(&Data{})

	return &SumRepository{
		db: db,
	}
}
