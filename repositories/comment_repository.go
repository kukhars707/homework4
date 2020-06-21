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

func (r *CommentRepository) GetComments() (*[]model.Comment, error) {
	fmt.Println("Not implement")
}

func (r *CommentRepository) GetComment(commentID uint) (*model.Comment, error) {
	fmt.Println("Not implement")
}

func (r *CommentRepository) CreateComment() error {
	fmt.Println("Not implement")
}

func (r *CommentRepository) EditComments(commentID uint) error {
	fmt.Println("Not implement")
}

func (r *CommentRepository) RemoveComments(commentID uint) error {
	fmt.Println("Not implement")
}