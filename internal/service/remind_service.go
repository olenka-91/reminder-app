package service

import (
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/olenka--91/reminder-app/internal/repository"
)

type RemindService struct {
	repo *repository.Repository
}

func NewRemindService(r *repository.Repository) *RemindService {
	return &RemindService{repo: r}
}

func (r *RemindService) Create(userID int, remind domain.Remind) (int, error) {
	return r.repo.Create(userID, remind)
}

func (r *RemindService) GetByID(userID int, remindID int) (domain.Remind, error) {
	return r.repo.GetByID(userID, remindID)
}
func (r *RemindService) GetAll(userID int) ([]domain.Remind, error) {
	return r.repo.GetAll(userID)
}
func (r *RemindService) Delete(userID, remindID int) error {
	return r.repo.Delete(userID, remindID)
}
func (r *RemindService) Update(userID, remindID int, input domain.RemindUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return r.repo.Update(userID, remindID, input)

}
