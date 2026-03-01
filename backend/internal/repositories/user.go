package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	DB *pgxpool.Pool
}

func GetUserByEmail(email string) {

}
