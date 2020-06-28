package services

import (
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type CommentService struct {
	repository repositories.CommentRepository
}

func NewCommentService(repository repositories.CommentRepository) CommentService {
	return CommentService{repository: repository}
}

func (s *CommentService) GetComments(projectID string, columnID string, taskID string) (*[]model.Comment, error) {
	return s.repository.GetComments(projectID, columnID, taskID)
}

func (s *CommentService) CreateComment(projectID string, columnID string, taskID string, text string) (*model.Comment, error) {
	comment := &model.Comment{
		Text: text,
		ProjectID: projectID,
		ColumnID: columnID,
		TaskID: taskID,
	}

	if err := s.repository.CreateComment(comment); err != nil {
		return comment, err
	}

	return comment, nil
}

func (s *CommentService) EditComments(commentID string, text string, projectID string, columnID string, taskID string) (*model.Comment, error) {
	comment := &model.Comment{
		ID: commentID,
		Text: text,
		ProjectID: projectID,
		ColumnID: columnID,
		TaskID: taskID,
	}

	if err := s.repository.CreateComment(comment); err != nil {
		return comment, err
	}

	return comment, nil
}

func (s *CommentService) RemoveComments(commentID string, projectID string, columnID string, taskID string) (*model.Comment, error) {
	return s.repository.RemoveComments(commentID, projectID, columnID, taskID)
}