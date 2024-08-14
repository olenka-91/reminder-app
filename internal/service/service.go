package service

import (
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/olenka--91/reminder-app/internal/repository"
)

type Remind interface {
	Create(userID int, remind domain.Remind) (int, error)
	GetByID(userID int, remindID int) (domain.Remind, error)
	GetAll(userID int) ([]domain.Remind, error)
	Delete(userID, remindID int) error
	Update(userID, remindID int, input domain.RemindUpdateInput) error
}

type Service struct {
	Remind
}

func NewService(r *repository.Repository) *Service {
	return &Service{Remind: NewRemindService(r)}
}
