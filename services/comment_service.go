package services

import (
	"database/sql"

	"goapps/apperrors"
	"goapps/models"
	"goapps/repositories"
)

type CommentService struct {
	CommentsInterface repositories.CommentInterface
}

func NewCommentService(conn *sql.Conn) *CommentService {
	return &CommentService{
		CommentsInterface: repositories.NewComments(conn),
	}
}

func (s *CommentService) InsertComment(comment models.Comment) error {
	if err := s.CommentsInterface.Insert(comment); err != nil {
		return apperrors.InsertErr.Wrap("Error Comment table insert. func: InsertComment", err)
	}
	return nil
}
