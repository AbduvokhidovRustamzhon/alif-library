package main

import (
	"context"
	"flag"
	"github.com/AbduvokhidovRustamzhon/library/cmd/crud/app"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/services/books"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/services/comments"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/services/files"
	"github.com/AbduvokhidovRustamzhon/mux2/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var (
	hostF   = flag.String("host", "", "Server host")
	portF   = flag.String("port", "", "Server port")
	dsnF    = flag.String("dsn", "", "Postgres DSN")
)
var (
	EHOST   = "HOST"
	EPORT   = "PORT"
	EDSN    = "DATABASE_URL"
)

func main() {
	flag.Parse()

	host, ok := FlagOrEnv(*hostF, EHOST)
	if !ok {
		log.Panic("can't get host")
	}
	port, ok := FlagOrEnv(*portF, EPORT)
	if !ok {
		log.Panic("can't get port")
	}
	dsn, ok := FlagOrEnv(*dsnF, EDSN)
	if !ok {
		log.Panic("can't get dsn")
	}
	addr := net.JoinHostPort(host, port)
	start(addr, dsn)
}

func start(addr string, dsn string) {
	router := mux.NewExactMux()
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	templatesPath := filepath.Join("web", "templates")
	assetsPath := filepath.Join("web", "assets")
	mediaPath := filepath.Join("web", "media")

	booksSvc := books.NewBooksSvc(pool)
	commentsSvc := comments.NewCommentsSvc(pool)
	filesSvc := files.NewFilesSvc(mediaPath)
	server := app.NewServer(
		router,
		pool,
		booksSvc,
		commentsSvc,
		filesSvc,
		templatesPath,
		assetsPath,
		mediaPath,
	)
	server.InitRoutes()

	panic(http.ListenAndServe(addr, server))
}
func FlagOrEnv(flag string, envKey string) (string, bool) {
	if flag != "" {
		return flag, true
	}
	return os.LookupEnv(envKey)
}