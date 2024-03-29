package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbConn *pgxpool.Pool

const createLinkBooksToAuthorsQuery = "create table if not exists linkbookstoauthors(books_id int references books (id) on update cascade, authors_id int references authors (id) on update cascade,primary key (books_id,authors_id))"
const createAuthorsQuery = "create table if not exists authors(id serial primary key, full_name varchar(255) unique not null)"
const createBooksQuery = "create table if not exists books(id serial primary key, title varchar(255) unique not null, release_year int not null, genre varchar(255))"
const createUsersQuery = "create table if not exists users(id serial primary key, email varchar(255) not null, password varchar(255) not null, created_at date default CURRENT_DATE)"
const createMembersQuery = "create table if not exists members(id serial primary key, full_name varchar(255) not null, phone_number char(10) unique not null, joined_at date default CURRENT_DATE, expires_at date default CURRENT_DATE+365, active bool not null default true)"
const createInstancesQuery = "create table if not exists instances(id serial primary key, book_id int references books (id) on update cascade, available bool not null default true)"
const createLoanQuery = "create table if not exists loans(id serial primary key, instance_id int references instances (id) on update cascade, member_id int references members (id) on update cascade, issued_at date default CURRENT_DATE, returned_at date , pending bool not null default true)"

func ConnectPostgres(url string) error {
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return err
	}
	dbConn = conn
	return nil
}

func CreateTables() error {
	tables := []string{
		createBooksQuery,
		createAuthorsQuery,
		createMembersQuery,
		createUsersQuery,
		createInstancesQuery,
		createLinkBooksToAuthorsQuery,
		createLoanQuery,
	}

	for _, query := range tables {
		_, err := dbConn.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}

func DropTables() error {
	tables := []string{
		"books",
		"authors",
		"members",
		"users",
		"instances",
		"linkbookstoauthors",
		"loans",
	}

	for _, query := range tables {
		query := fmt.Sprintf("drop table if exists %s cascade", query)
		_, err := dbConn.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}
