package task

import "fmt"

type repo interface {
	GetAll() ([][]string, error)
	Add(input []string) error
	UpdateStatus(i int, active bool) error
	DeleteDisabled() error
}

type Service struct {
	db repo
}

func New(db repo) *Service {
	return &Service{
		db: db,
	}
}

func (s Service) AddTask(input []string) error {
	return s.db.Add(input)
}

func (s Service) GetTasks() ([]string, error) {
	result, err := s.db.GetAll()
	if err != nil {
		return nil, err
	}

	formattedResult := []string{}
	for _, r := range result {
		formattedResult = append(formattedResult, fmt.Sprintf("%s: %s", r[0], r[1]))
	}

	return formattedResult, err
}

func (s Service) ChangeTaskStatus(i int, active bool) error {
	return s.db.UpdateStatus(i, active)
}

func (s Service) CleanDoneTasks() error {
	return s.db.DeleteDisabled()
}
