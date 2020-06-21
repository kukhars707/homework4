package services

import (
	"fmt"
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type CommentService struct {
	repository repositories.CommentRepository
}

func NewCommentService(repository repositories.CommentRepository) CommentService {
	return CommentService{repository: repository}
}

func (s *CommentService) GetComments() ([]model.Comment, error) {
	fmt.Println("Not implement")
}

func (s *CommentService) GetComment(commentID uint) (model.Comment, error) {
	fmt.Println("Not implement")
}

func (s *CommentService) CreateComment() error {
	fmt.Println("Not implement")
}

func (s *CommentService) EditComments(commentID uint) error {
	fmt.Println("Not implement")
}

func (s *CommentService) RemoveComments(commentID uint) error {
	fmt.Println("Not implement")
}