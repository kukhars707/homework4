package repositories

import (
	"database/sql"
	model "github.com/kukhars707/homework4/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return CommentRepository{db: db}
}

func (r *CommentRepository) GetComments(taskID string) (*[]model.Comment, error) {
	var c model.Comment
	sqlStatement := `SELECT * FROM comment WHERE projectID = $1`

	result, err := r.db.Query(sqlStatement, taskID)

	if err != nil {
		panic(err.Error())
	}

	var cSlice []model.Comment
	for result.Next() {
		result.Scan(&c.ID, &c.Text, &c.TaskID)
		cSlice = append(cSlice, c)
	}

	return &cSlice, nil
}

func (r *CommentRepository) CreateComment(comment *model.Comment) error {
	sqlStatement := `INSERT INTO comment (text, taskID) VALUES ($1, $2);`

	_, err := r.db.Exec(sqlStatement, comment.Text, comment.TaskID)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *CommentRepository) EditComments(comment *model.Comment) error {
	sqlStatement := `UPDATE comment SET text = $1 WHERE id = $2;`

	_, err := r.db.Exec(sqlStatement, comment.Text, comment.ID)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *CommentRepository) RemoveComments(commentID string) error {
	sqlStatement := `DELETE FROM comment WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, commentID)

	if err != nil {
		panic(err.Error())
	}

	return err
}