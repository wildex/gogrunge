package main

import (
	"crypto/rand"
	"fmt"
)

type userSession struct {
	Name  string
	Admin bool
}

type Session struct {
	data map[string]*userSession
}

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func NewSession() *Session {
	s := new(Session)
	s.data = make(map[string]*userSession)
	return s
}

/* Adds NAME to map and returns session id */
func (s *Session) Add(name string) string {
	id := GenerateId()

	data := &userSession{Name: name, Admin: false}
	s.data[id] = data

	return id
}

/* Return name related to session id */
func (s *Session) Get(id string) string {

	data := s.data[id]

	if data == nil {
		return ""
	}

	return data.Name
}

func (s *Session) AddRoot(name string) string {
	id := GenerateId()

	data := &userSession{Name: name, Admin: true}
	s.data[id] = data

	return id
}

func (s *Session) IsRoot(id string) bool {
	data := s.data[id]

	if data == nil {
		return false
	}

	return data.Admin
}
