package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *API) handleMember(r chi.Router) {
	r.Post("/", makeHttpFunc(s.handleCreateMember))
	r.Get("/{id}", makeHttpFunc(s.handleGetMember))
}

func (s *API) handleGetMember(w http.ResponseWriter, r *http.Request) error {
	idOfMember := chi.URLParam(r, "id")
	i, err := strconv.Atoi(idOfMember)
	if err != nil {
		return err
	}
	member, err := s.db.getMemberById(i)
	writeJson(w, r, member, http.StatusOK)
	return nil
}

func (s *API) handleCreateMember(w http.ResponseWriter, r *http.Request) error {
	var m Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return err
	}
	err := s.db.createMemeber(&m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	writeJson(w, r, m, http.StatusOK)
	return nil
}

func (s *dbPgx) createMembersTable() error {
	_, err := s.conn.Query(context.Background(), "create table if not exists members(id serial primary key, full_name varchar(255) not null, phone_number char(10) not null unique, joined_at date default CURRENT_DATE, expires_at date default CURRENT_DATE+365, current_status varchar(10) default 'active')")
	if err != nil {
		return err
	}
	return nil
}

func (s *dbPgx) getAllMembers() ([]Member, error) {
	members := []Member{}
	rows, err := s.conn.Query(context.Background(), "select id,full_name,phone_number,current_status from Members")
	if err != nil {
		return nil, err
	}

	var Name string
	var Phone string
	var id int
	var status string

	for rows.Next() {
		err := rows.Scan(&id, &Name, &Phone, &status)
		if err != nil {
			return nil, err
		}
		members = append(members, Member{id, Name, Phone, status})
	}
	return members, nil
}

func (s *dbPgx) getMemberById(id int) (*Member, error) {
	var a Member

	err := s.conn.QueryRow(context.Background(), "select id,full_name,phone_number,current_status from Members where id=$1", id).Scan(&a.Id, &a.Name, &a.Phone, &a.Status)
	if err != nil {
		return nil, err
	}

	return &a, err
}

func (s *dbPgx) createMemeber(a *Member) error {

	_, err := s.conn.Query(context.Background(), "insert into members(full_name,phone_number) values ($1,$2)", a.Name, a.Phone)
	if err != nil {
		return err
	}

	return nil
}
