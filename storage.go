package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccounts(int) ([]*Account, error)
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

type PostgresConfig struct {
	DbUser     string
	DbName     string
	DbPassword string
}

func NewPostgresStore(pc PostgresConfig) (*PostgresStore, error) {
	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable", pc.DbUser, pc.DbName, pc.DbPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}
func (s *PostgresStore) GetAccounts(limit int) ([]*Account, error) {
	query := `SELECT * FROM accounts LIMIT $1;`
	rows, err := s.db.Query(query, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []*Account

	for rows.Next() {
		account := new(Account)

		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number,
			&account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil

}
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO accounts (first_name , last_name , number ,balance , created_at) VALUES ($1,$2,$3,$4,$5)`
	_, err := s.db.Query(query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {

	return nil
}
func (s *PostgresStore) UpdateAccount(acc *Account) error {

	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	return nil, nil
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL,
			first_name varchar(50),
			last_name varchar(50),
			number SERIAL,
			balance NUMERIC(9,2),
			created_at TIMESTAMP,
			PRIMARY KEY(id)
		);
	`

	_, err := s.db.Exec(query)

	return err

}

func (s *PostgresStore) DropAccountTable() error {
	query := `
		DROP TABLE IF EXISTS accounts;
	`
	_, err := s.db.Exec(query)

	return err
}
