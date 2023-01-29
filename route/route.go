package route

import (
	"database/sql"
	"net/http"

	"goapps/handlers"

	"github.com/gorilla/mux"
)

type CreateHandlers struct {
	conn *sql.Conn
}

func NewCreateHandlers(conn *sql.Conn) *CreateHandlers {
	return &CreateHandlers{
		conn: conn,
	}
}

func (h *CreateHandlers) Routes() *mux.Router {
	article_handlers := handlers.NewArticleHandlers(h.conn)

	r := mux.NewRouter()
	r.HandleFunc("/article", article_handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list/{page:[0-9]+}", article_handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", article_handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", article_handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/comment", article_handlers.PostCommentHandler).Methods(http.MethodPost)

	// if you add a new handler, write this area.

	return r
}
