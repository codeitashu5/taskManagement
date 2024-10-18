package task

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"taskManagement/models"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) registerUser(register RegisterBody) (primitive.ObjectID, error) {
	userId, err := s.repository.registerUser(register)
	if err != nil {
		return userId, err
	}

	return userId, err
}

func (s *Service) login(register LoginRequest) (models.User, error) {
	user, err := s.repository.login(register)
	if err != nil {
		return user, err
	}

	return user, err
}

func (s *Service) createTask(task Request, userId primitive.ObjectID) error {
	err := s.repository.createTask(task, userId)
	return err
}

func (s *Service) getTask(taskID primitive.ObjectID) (*Task, error) {
	return s.repository.getTask(taskID)
}

func (s *Service) getAllTasks(userId primitive.ObjectID) ([]Task, error) {
	return s.repository.getAllTasks(userId)
}

func (s *Service) deleteTask(taskID primitive.ObjectID) error {
	return s.repository.deleteTask(taskID)
}

func (s *Service) updateTask(taskID primitive.ObjectID, taskData Request) error {
	return s.repository.updateTask(taskID, taskData)
}
