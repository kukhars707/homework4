package repositories

import (
	"database/sql"
	"fmt"
	model "github.com/kukhars707/homework4/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return CommentRepository{db: db}
}

func (r *CommentRepository) GetComments(projectID string, columnID string, taskID string) (*[]model.Comment, error) {
	fmt.Println("Not implement")
}

func (r *CommentRepository) CreateComment(comment *model.Comment) error {
	fmt.Println("Not implement")
}

func (r *CommentRepository) EditComments(comment *model.Comment) error {
	fmt.Println("Not implement")
}

func (r *CommentRepository) RemoveComments(commentID string, taskID string, projectID string, columnID string) error {
	fmt.Println("Not implement")
}