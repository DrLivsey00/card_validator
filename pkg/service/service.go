package service

type CardService interface {
	IsValid(number string, expMonth, expYear int) (bool, error, string)
}

type Service struct {
	CardService
}

func NewService() *Service {
	return &Service{
		CardService: NewCardSrv(),
	}
}
