package app

import (
	"errors"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/services/books"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/services/comments"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/services/files"
	"github.com/AbduvokhidovRustamzhon/mux2/pkg/mux"
	_ "github.com/AbduvokhidovRustamzhon/mux2/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type server struct {
	pool          *pgxpool.Pool
	router        *mux.ExactMux
	booksSvc      *books.BooksSvc
	commentsSvc   *comments.CommentsSvc
	filesSvc      *files.FilesSvc
	templatesPath string
	assetsPath    string
	mediaPath     string
	storagePath   string
}

func NewServer(router *mux.ExactMux, pool *pgxpool.Pool, burgersSvc *books.BooksSvc, CommentSvc *comments.CommentsSvc, filesSvc *files.FilesSvc, templatesPath string, assetsPath string, mediaPath string) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	if burgersSvc == nil {
		panic(errors.New("booksSvc can't be nil"))
	}
	if filesSvc == nil {
		panic(errors.New("booksSvc can't be nil"))
	}
	if templatesPath == "" {
		panic(errors.New("templatesPath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}
	if mediaPath == "" {
		panic(errors.New("mediaPath can't be empty"))
	}
	if CommentSvc == nil {
		panic(errors.New("CommentSvc can't be nil"))
	}
	return &server{
		router:        router,
		pool:          pool,
		booksSvc:      burgersSvc,
		filesSvc:      filesSvc,
		templatesPath: templatesPath,
		assetsPath:    assetsPath,
		mediaPath:     mediaPath,
		commentsSvc:   CommentSvc,
	}
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}
