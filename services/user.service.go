package services

import "github.com/kahfiredo/test_pertama/models"

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
