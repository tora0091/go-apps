package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"goapps/domain"
	"goapps/models"
	"goapps/services"

	"github.com/gorilla/mux"
)

type Handlers struct {
	articleService *services.ArticleService
	commentService *services.CommentService
}

func NewArticleHandlers(conn *sql.Conn) *Handlers {
	return &Handlers{
		articleService: services.NewArticleService(conn),
		commentService: services.NewCommentService(conn),
	}
}

// curl http://localhost:8080/article -X POST -d '{"title":"Handlers test.", "article_id":9, "contents":"are you seeing anyone?", "user_name":"take-001"}'
func (h *Handlers) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// [Point!!]
	// リクエストを取得したのち、メモリに格納し、そのデータを構造体にあてはめている
	// → いったんメモリに格納している
	// 下部ロジックは、リクエストを取得、ストリームをとうして構造体にあてはめている
	// → メモリを介さないので効率的
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusBadGateway,
			Data:   nil,
			Err:    err,
		})
		return
	}

	article := reqArticle

	if err := h.articleService.InsertArticle(article); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusInternalServerError,
			Data:   nil,
			Err:    err,
		})
		return
	}

	// こちらもストリームを介すことで無駄なメモリを利用することなく目的を達せれる
	json.NewEncoder(w).Encode(domain.ResponseData{
		Status: http.StatusOK,
		Data:   article,
		Err:    nil,
	})
}

// curl http://localhost:8080/article/nice -X POST -d '{"title":"Handlers test.", "article_id":9, "contents":"are you seeing anyone?", "user_name":"take-001"}'
func (h *Handlers) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusBadGateway,
			Data:   nil,
			Err:    err,
		})
		return
	}

	article := reqArticle

	if err := h.articleService.UpdateArticleNiceNum(article.ID); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusInternalServerError,
			Data:   nil,
			Err:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(domain.ResponseData{
		Status: http.StatusOK,
		Data:   article,
		Err:    nil,
	})
}

// curl http://localhost:8080/article/comment -X POST -d '{"comment_id":1, "article_id":3,"message":"here you go"}'
func (h *Handlers) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err,
		})
		return
	}

	comment := reqComment

	if err := h.commentService.InsertComment(comment); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusInternalServerError,
			Data:   nil,
			Err:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(domain.ResponseData{
		Status: http.StatusOK,
		Data:   comment,
		Err:    nil,
	})
}

// curl http://localhost:8080/article/list/0 -X GET
func (h *Handlers) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	articles, err := h.articleService.GetArticleList()
	if err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusInternalServerError,
			Data:   nil,
			Err:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(domain.ResponseData{
		Status: http.StatusOK,
		Data:   articles,
		Err:    nil,
	})
}

// curl http://localhost:8080/article/9 -X GET
func (h *Handlers) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err,
		})
		return
	}

	article, err := h.articleService.GetArticleDetail(articleId)
	if err != nil {
		json.NewEncoder(w).Encode(domain.ResponseData{
			Status: http.StatusInternalServerError,
			Data:   nil,
			Err:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(domain.ResponseData{
		Status: http.StatusOK,
		Data:   article,
		Err:    nil,
	})
}
