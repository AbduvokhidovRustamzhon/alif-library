package books

import (
	"context"
	"errors"
	"github.com/AbduvokhidovRustamzhon/library/pkg/crud/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type BooksSvc struct {
	pool *pgxpool.Pool
}

func NewBooksSvc(pool *pgxpool.Pool) *BooksSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	return &BooksSvc{pool: pool}
}

func (service *BooksSvc) Books(ctx context.Context) (list []models.Book, err error) {
	list = make([]models.Book, 0)
	conn, err := service.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, name, pages, file_name FROM books WHERE removed = FALSE")
	if err != nil {
		return nil, err // TODO: wrap to specific error
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Book{}
		err := rows.Scan(&item.Id, &item.Name, &item.Pages, &item.FileName)
		if err != nil {
			return nil, err // TODO: wrap to specific error
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *BooksSvc) Save(ctx context.Context, model models.Book) (err error) {
	conn, err := service.pool.Acquire(ctx)
	if err != nil {
		log.Print(err)
		return errors.New("can't execute pool: ")
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "INSERT INTO books(name, pages, file_name) VALUES ($1, $2, $3);", model.Name, model.Pages, model.FileName)
	if err != nil {
		log.Print(err)
		return errors.New("can't save burger: ")
	}
	return nil
}

func (service *BooksSvc) RemoveById(ctx context.Context, id int) (err error) {
	conn, err := service.pool.Acquire(ctx)
	if err != nil {
		return errors.New("can't execute pool: ")
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "UPDATE books SET removed = true where id = $1;", id)
	if err != nil {
		return errors.New("can't remove burger: ")
	}
	return nil
}

