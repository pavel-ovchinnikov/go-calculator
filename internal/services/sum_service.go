package service

type SumRepository interface {
	GetSum(a, b int) (int, error)
	SetSum(a, b, res int) error
}

type SumService struct {
	name string
	rep  SumRepository
}

func NewSumService(rep SumRepository) *SumService {
	return &SumService{
		name: "service.sum",
		rep:  rep,
	}
}

func (s SumService) Sum(a, b int) (int, error) {
	res, err := s.rep.GetSum(a, b)

	if err == nil {
		return res, nil
	}

	res = a + b
	err = s.rep.SetSum(a, b, res)
	if err != nil {
		return 0, err
	}

	return res, nil
}
